// Package validates ...
/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-05 10:28:04
 * @LastEditTime: 2019-09-10 15:06:11
 * @LastEditors: Dawn
 */
package validates

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	//"github.com/wangcong0918/sunrise/log"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

// CheckParameter make validates ...
func CheckParameter(s interface{}) (msg bool, info string) {
	var messageInfo string
	var message bool // = false
	zh_ch := zh.New()
	var validate = validator.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			//fmt.Println(err.Translate(trans))
			messageInfo += err.Translate(trans)
		}
		return message, messageInfo
	}
	message = true
	return message, ""
}
