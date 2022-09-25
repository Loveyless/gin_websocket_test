package validator

//gin > 1.4.0

//将验证器错误翻译成中文

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	//注册翻译器 New返回一个新的UniversalTranslator实例集，其中包含回退语言环境和它应该支持的语言环境
	zh := zh.New()
	uni = ut.New(zh, zh)

	//为给定的地区返回指定的翻译程序
	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate = binding.Validator.Engine().(*validator.Validate)

	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

//翻译错误信息
func Translate(err error) map[string][]string {

	var result = make(map[string][]string)

	errors, ok := err.(validator.ValidationErrors)

	if ok {
		for _, err := range errors {
			result[err.Field()] = append(result[err.Field()], err.Translate(trans))
		}
		return result
	} else {
		//目前发现只有类型错误 才会无法断言 err.(validator.ValidationErrors)
		result := map[string][]string{
			"TypeError": {"参数类型错误"},
		}
		return result
	}

}
