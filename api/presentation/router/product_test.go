package router

import (
	"fmt"
	"testing"
)

func Test_ProductControllerImpl_GetProduct(t *testing.T) {

	// initdb.d db
	setData("product.sql")

	tests := []testCase{
		{
			name:   "正常系_200",
			method: "GET",
			url:    fmt.Sprintf("%v/products/%v", ts.URL, "101"),
			header: map[string]string{"Content-Type": "application/json"},
			want: want{
				status: 200,
				body: `
					{
						"product_id": 101,
						"product_name": "商品名101",
						"brand_name": "ブランド名201"
					}
					`,
			},
		},
		{
			name:   "異常系_400_product_idが0の場合",
			method: "GET",
			url:    fmt.Sprintf("%v/products/%v", ts.URL, "0"),
			header: map[string]string{"Content-Type": "application/json"},
			want: want{
				status: 400,
				body: `
					{
						"status": 400,
						"code": "400001",
						"message": "Key: 'Product.ProductID' Error:Field validation for 'ProductID' failed on the 'min' tag"
					}
					`,
			},
		},
		{
			name:   "異常系_404_データが存在しない場合",
			method: "GET",
			url:    fmt.Sprintf("%v/products/%v", ts.URL, "99999"),
			header: map[string]string{"Content-Type": "application/json"},
			want: want{
				status: 404,
				body:   `{ "status": 404, "code": "404001", "message": "product is empty" }`,
			},
		},
	}
	// テストケース実行
	runTestCase(t, tests)
}

func Test_ProductControllerImpl_AddProduct(t *testing.T) {

	// initdb.d db
	setData("product.sql")

	tests := []testCase{
		{
			name:   "正常系_200",
			method: "POST",
			url:    fmt.Sprintf("%v/products", ts.URL),
			header: map[string]string{"Content-Type": "application/json"},
			body: `
					{
						"name": "商品名NEW1",
						"brand_id": 201
					}
					`,
			want: want{
				status: 200,
				body:   `{"product_id": 106}`,
			},
		},
		{
			name:   "異常系_400_nameが空の場合",
			method: "POST",
			url:    fmt.Sprintf("%v/products", ts.URL),
			header: map[string]string{"Content-Type": "application/json"},
			body: `
					{
						"name": "",
						"brand_id": 201
					}
					`,

			want: want{
				status: 400,
				body: `
						{
							"status": 400,
							"code": "400001",
							"message": "Key: 'AddProduct.Name' Error:Field validation for 'Name' failed on the 'min' tag"
						}
					`,
			},
		},
	}
	// テストケース実行
	runTestCase(t, tests)
}

func Test_ProductControllerImpl_EditProduct(t *testing.T) {

	// initdb.d db
	setData("product.sql")

	tests := []testCase{
		{
			name:   "正常系_200",
			method: "PUT",
			url:    fmt.Sprintf("%v/products/%v", ts.URL, 101),
			header: map[string]string{"Content-Type": "application/json"},
			body: `
					{
						"name": "商品名EDIT1",
						"brand_id": 202
					}
					`,
			want: want{
				status: 200,
			},
		},
		{
			name:   "異常系_400_nameが空の場合",
			method: "PUT",
			url:    fmt.Sprintf("%v/products/%v", ts.URL, 101),
			header: map[string]string{"Content-Type": "application/json"},
			body: `
					{
						"name": "",
						"brand_id": 202
					}
					`,

			want: want{
				status: 400,
				body: `
					{
						"status": 400,
						"code": "400001",
						"message": "Key: 'EditProduct.Name' Error:Field validation for 'Name' failed on the 'min' tag"
					}
					`,
			},
		},
	}
	// テストケース実行
	runTestCase(t, tests)
}

func Test_ProductControllerImpl_DeleteProduct(t *testing.T) {

	// initdb.d db
	setData("product.sql")

	tests := []testCase{
		{
			name:   "正常系_200",
			method: "DELETE",
			url:    fmt.Sprintf("%v/products/%v", ts.URL, 101),
			header: map[string]string{"Content-Type": "application/json"},
			want: want{
				status: 200,
			},
		},
		{
			name:   "異常系_400_product_idが0の場合",
			method: "DELETE",
			url:    fmt.Sprintf("%v/products/%v", ts.URL, 0),
			header: map[string]string{"Content-Type": "application/json"},
			want: want{
				status: 400,
				body: `
					{
						"status": 400,
						"code": "400001",
						"message": "Key: 'Product.ProductID' Error:Field validation for 'ProductID' failed on the 'min' tag"
					}
					`,
			},
		},
	}
	// テストケース実行
	runTestCase(t, tests)
}
