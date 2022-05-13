package validation

import (
	"unicode"

	validator "github.com/go-playground/validator/v10"
)

// 参考: https://github.com/go-playground/validator/blob/c68441b7f4748b48ad9a0c9a79d346019730e207/baked_in.go#L65
var (
	Default = validator.New()
)

func init() {
	Default.RegisterValidation("hiragana", isHiragana)
}

// isHiragana
func isHiragana(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return false
	}
	for _, r := range fl.Field().String() {
		if !unicode.In(r, unicode.Hiragana) {
			return false
		}
	}
	return true
}
