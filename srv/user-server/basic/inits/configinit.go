package inits

import (
	"cms/srv/user-server/basic/config"
	"fmt"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigFile("C:\\Users\\ZhuanZ\\Desktop\\zuoye11\\cms\\srv\\config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config.GlobalConf)
	if err != nil {
		return
	}
	fmt.Println("配置文件读取成功")
}
