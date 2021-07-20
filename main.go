package main

import (
	"fmt"
	"github.com/1975210542/wxbot/wx"
	"runtime"
)

//func init() {
//	//加载配置文件
//	setting, err := setting.NewSetting()
//	if err != nil {
//		return
//	}
//	if err = setting.ReadSection("Wx", &global.WxSetting); err != nil {
//		return
//	}
//}

func main() {
	defer func() {
		if err := recover(); err != nil {

			errStr := ""
			switch x := err.(type) {
			case runtime.Error:
				// 这是运行时错误类型异常
				errStr = x.Error()
			case error:
				// 普通错误类型异常
			default:
				// 其他类型异常
			}
			data := wx.NewWechatApplication("@all")
			data.WithCallersFrames().SendTextMessage(errStr)

		}
	}()

	var res []int
	res = append(res, 10)
	fmt.Println("res:", res[10])

	for i := 0; i < 20; i++ {
		fmt.Println("打印i：", i)
	}
}
