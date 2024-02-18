package models

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 设置自带的 store 存在服务器内存中
//var store = base64Captcha.DefaultMemStore

// 设置自定义的 store 存在redis中
var store base64Captcha.Store = RedisStore{}

// GetCaptcha 获取验证码
func GetCaptcha() (string, string, error) {
	// 生成验证码
	var driver base64Captcha.Driver
	//driver = base64Captcha.NewDriverString(80, 240, 0, 0, 4, "1234567890", nil)
	//设置验证码的配置
	driverString := base64Captcha.DriverChinese{
		Height:          40,           //设置验证码图片的高度
		Width:           100,          //设置验证码图片的宽度
		NoiseCount:      0,            //设置干扰线的数量
		ShowLineOptions: 2 | 4,        //设置线条的类型
		Length:          4,            //设置验证码的长度
		Source:          "1234567890", //设置验证码的字符源
		BgColor: &color.RGBA{ //设置验证码图片的背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"}, //设置验证码的字体
	}

	//生成验证码
	driver = driverString.ConvertFonts()

	//生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)
	//验证码id base64图片字符串 验证码字符串 error
	id, b64s, _, err := c.Generate()
	return id, b64s, err
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	//验证验证码，参数1是验证码的id，参数2是用户输入的验证码
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
