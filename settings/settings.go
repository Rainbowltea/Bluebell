package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//使用viper加载配置文件
func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件路径
	viper.SetConfigType("yaml")   //制定配置文件类型
	viper.AddConfigPath(".")      //制定查找配置文件的路径
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {               // 读取配置信息失败
		fmt.Println("re", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
	})
	return
}
