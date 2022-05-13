package router

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/haradayoshitsugucz/purple-server/api/di"
	"github.com/haradayoshitsugucz/purple-server/api/env"
	"github.com/haradayoshitsugucz/purple-server/config"
	"github.com/haradayoshitsugucz/purple-server/constant"
	"github.com/haradayoshitsugucz/purple-server/logger"
)

var (
	ts       *httptest.Server
	testDate = flag.String("date", "2021-11-25 10:00:00", "set date. yyyy-mm-dd HH:MM:SS")
)

// TestMain
func TestMain(m *testing.M) {

	// テスト全体に関する前処理
	setUp()
	// テスト全体に関する後処理
	defer cleanUp()
	// テスト実行
	//code := m.Run()
	m.Run()
	//os.Exit(code)
}

func setUp() {

	log.Println("before all...")

	conf := env.NewConfig(config.Test)
	logger.InitLogger(conf, "")

	flag.Parse()

	utc := time.FixedZone("UTC", 0)
	time.Local = utc

	date, err := time.ParseInLocation(constant.DatetimeFormat, *testDate, time.Local)
	if err != nil {
		log.Panic(fmt.Sprintf("DATE FORMAT ERROR. err:%+v\n", err.Error()))
	}

	handler := di.InitHandler(conf, func() time.Time { return date }, cli)
	ctrl := handler.Ctrl
	middle := handler.Middle

	// router
	mux := chi.NewRouter()
	mux.Use(middle.TimeMiddle.Now())
	mux.Get("/status", ctrl.StatusController.Ping())
	mux.Get("/products/{product_id}", ctrl.ProductController.GetProduct(conf))
	mux.Delete("/products/{product_id}", ctrl.ProductController.DeleteProduct(conf))
	mux.Put("/products/{product_id}", ctrl.ProductController.EditProduct(conf))
	mux.Post("/products", ctrl.ProductController.AddProduct(conf))
	mux.Get("/products", ctrl.ProductController.GetProductsByName(conf))

	// server
	ts = httptest.NewServer(mux)

}

func setData(fileNames ...string) {

	// 全テーブルtruncate
	fileNames = append([]string{"truncate.sql"}, fileNames...)

	for _, fileName := range fileNames {

		// DBテストデータで初期化
		cmd := exec.Command("../../resource/test/bin/exec_sql_file.sh", fileName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Panic(fmt.Sprintf("Command Exec Error. err:%+v\n", err.Error()))
		}
	}
}

func cleanUp() {
	log.Println("after all...")
	ts.Close()
}

type want struct {
	status int
	body   string // レスポンス JSON BODY
	check  func(t *testing.T, resp *http.Response, respBody string)
}

type testCase struct {
	name   string
	method string
	url    string
	header map[string]string
	body   string // POSTするJSON BODY
	want   want
}

func runTestCase(t *testing.T, testCases []testCase) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody io.Reader
			if len(tt.body) > 0 {
				reqBody = bytes.NewBufferString(tt.body)
			}

			req, err := http.NewRequest(strings.ToUpper(tt.method), tt.url, reqBody)
			if err != nil {
				t.Error(err)
				return
			}

			for key, value := range tt.header {
				req.Header.Add(key, value)
			}

			client := new(http.Client)
			resp, err := client.Do(req)
			if err != nil {
				t.Error(err)
				return
			}

			//responseのステータスコードのテスト
			if resp.StatusCode != tt.want.status {
				t.Errorf("statusCode test Invalid. \nwant=%+v\nactual=%+v", tt.want.status, resp.StatusCode)
			}

			respBody := ""

			// body
			if len(tt.want.body) > 0 {

				b, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Error(err)
				}
				respBody = string(b)

				var respMap map[string]interface{}
				if err := json.Unmarshal(b, &respMap); err != nil {
					t.Error(err)
				}

				var wantBodyMap map[string]interface{}
				if err := json.Unmarshal([]byte(tt.want.body), &wantBodyMap); err != nil {
					t.Error(err)
				}

				if !reflect.DeepEqual(wantBodyMap, respMap) {
					t.Errorf("body test Invalid. \nwant=%+v\nactual=%+v", tt.want.body, respBody)
				}
			}

			// テストケース毎の固有のテスト
			if tt.want.check != nil {
				tt.want.check(t, resp, respBody)
			}
		})
	}
}
