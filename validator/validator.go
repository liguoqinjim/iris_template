package validator

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/logger"
	"reflect"

	//"fmt"
	//"github.com/go-playground/locales/zh"
	//ut "github.com/go-playground/universal-translator"
	//"github.com/go-playground/validator/v10"
	//"reflect"
	//
	////en_translations "github.com/go-playground/validator/v10/translations/en"
	//zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni           *ut.UniversalTranslator
	Validate      *validator.Validate
	ValidateTrans ut.Translator
)

func init() {
	initValidator()
}

func initValidator() {
	zh2 := zh.New()
	uni = ut.New(zh2, zh2)

	ValidateTrans, _ = uni.GetTranslator("zh")

	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})

	zh_translations.RegisterDefaultTranslations(Validate, ValidateTrans)

	//自定义翻译
	//Validate.RegisterTranslation("required", ValidateTrans, func(ut ut.Translator) error {
	//	return ut.Add("required", "{0} 不能为空!", true) // see universal-translator for details
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("required", fe.Field())
	//
	//	return t
	//})
	//
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})
}

func ValidateStruct(s interface{}) error {
	if err := Validate.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		logger.Errorf(errs[0].Translate(ValidateTrans))
		return consts.ErrorValidate
	} else {
		return nil
	}
}

func ValidateValue(v interface{}, tag string) error {
	if err := Validate.Var(v, tag); err != nil {
		errs := err.(validator.ValidationErrors)
		logger.Errorf(errs[0].Translate(ValidateTrans))
		return consts.ErrorValidate
	} else {
		return nil
	}
}
