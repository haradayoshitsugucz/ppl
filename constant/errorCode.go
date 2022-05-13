package constant

/*
presentation層で返すエラーコードを定義します
生成ルール: HTTP_STATUS + 連番(3桁)
(注意) エラー箇所特定のため、同じエラーコードを2箇所で使わないこと
*/
const (

	// StatusBadRequest
	Err400001 = "400001"

	// StatusUnauthorized
	Err401001 = "401001"

	// StatusForbidden
	Err403001 = "403001"

	// StatusNotFound
	Err404001 = "404001"

	// StatusNotAcceptable
	Err406001 = "406001"

	//StatusConflict
	Err409001 = "409001"

	//StatusInternalServerError
	Err500001 = "500001"
)
