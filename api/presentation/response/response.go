package response

import (
	"encoding/json"
	"net/http"

	"github.com/haradayoshitsugucz/purple-server/logger"
	"go.uber.org/zap"
)

type errorResponse struct {
	Status  int         `json:"status"` // HTTPステータス
	Code    string      `json:"code"`   // エラーコード　(constantパッケージの定数を使用)
	Message interface{} `json:"message"`
}

func newErrorResponse(status int, code string, message interface{}) *errorResponse {
	return &errorResponse{Status: status, Code: code, Message: message}
}

// Success 200 だけ返したい場合に使用
func Success(w http.ResponseWriter) int {
	w.WriteHeader(http.StatusOK)
	return http.StatusOK
}

// SuccessWith 200とJSONを返す場合に使用
func SuccessWith(w http.ResponseWriter, out interface{}) int {
	tBytes, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tBytes)

	return http.StatusOK
}

// ErrorWith HTTPエラーとJSONを返す場合に使用
func ErrorWith(w http.ResponseWriter, httpStatus int, errCode string, errMsg interface{}) int {

	// logger
	logging(httpStatus, errCode, errMsg)

	errorRes := newErrorResponse(httpStatus, errCode, errMsg)
	tBytes, err := json.MarshalIndent(errorRes, "", "  ")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(tBytes)

	return httpStatus
}

// logging
func logging(httpStatus int, errCode string, errMsg interface{}) {
	switch {
	case 400 <= httpStatus && httpStatus <= 499:
		logger.GetLogger().Warn("", zap.String("errCode", errCode), zap.Any("errMsg", errMsg))
	case 500 <= httpStatus && httpStatus <= 599:
		logger.GetLogger().Error("", zap.String("errCode", errCode), zap.Any("errMsg", errMsg))
	default:
		logger.GetLogger().Info("", zap.String("errCode", errCode), zap.Any("errMsg", errMsg))
	}
}
