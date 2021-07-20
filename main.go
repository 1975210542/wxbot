package main

import (
	"fmt"
	"github.com/1975210542/wxbot/global"
	"github.com/1975210542/wxbot/setting"
	"github.com/1975210542/wxbot/wx"
)

func init() {
	//加载配置文件
	setting, err := setting.NewSetting()
	if err != nil {
		return
	}
	if err = setting.ReadSection("Wx", &global.WxSetting); err != nil {
		return
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			data := wx.NewWechatApplication("@all")
			rest := data.SendTextMessage(err)
			fmt.Println(rest)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("chufa:", 10/i)
	}
}
