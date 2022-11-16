package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitCfg() (err error) {
	viper.SetConfigFile("config.yaml")
	// 下面这种是用于远程的配置文件
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置信息失败...")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("读取配置文件成功...")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了...")
	})
	return
}
