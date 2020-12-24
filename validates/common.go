package validates

import (
	"gin_test/pkg/logging"

	validator "github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni           *ut.UniversalTranslator
	validate      *validator.Validate
	validateTrans ut.Translator
)

func init() {
	zh2 := zh.New()
	uni = ut.New(zh2, zh2)
	local_language := "zh"
	validateTrans, _ = uni.GetTranslator(local_language)

	validate = validator.New()
	// 选择语言
	var err error
	switch local_language {
	case "zh":
		err = zh_translations.RegisterDefaultTranslations(validate, validateTrans)
	default:
		err = zh_translations.RegisterDefaultTranslations(validate, validateTrans)
	}

	if err != nil {
		logging.LogError(err)
	}
}
