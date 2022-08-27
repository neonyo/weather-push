package main

import (
	"fmt"
	"github.com/spf13/viper"
	"love/core"
)

func main() {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	wx := core.NewWeChat(config)
	wx.Send()
	fmt.Println("ok")
}
