package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd() // 获取当前工作目录
	fmt.Println("Current working directory:", workDir)
	viper.SetConfigName("config")            // 设置配置文件名
	viper.SetConfigType("yml")               // 设置配置文件类型
	viper.AddConfigPath(workDir + "/config") // 设置配置文件路径
	err := viper.ReadInConfig()              // 读取配置文件路径
	if err != nil {
		return
	}

}
