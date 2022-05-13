package controller

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	StatusController  StatusController
	ProductController ProductController
}

// bindJson パラメータBODY(JSON)をStructにバインドします
func bindJson(r *http.Request, form interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&form)
	return err
}
