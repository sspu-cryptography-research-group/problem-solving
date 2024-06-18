package utils

import (
	"log"

	"github.com/spf13/viper"
)

type MysqlConf struct {
	User     string
	Password string
	Server   string
	Port     string
	Db       string
}

func GetConfig() (string, string, MysqlConf) {
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() //读取配置
	if err != nil {
		log.Fatalln(err)
	}
	//获取配置信息
	repo := viper.GetString("repo")
	token := viper.GetString("github_token")
	dbconf := MysqlConf{
		User:     viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		Server:   viper.GetString("mysql.server"),
		Port:     viper.GetString("mysql.port"),
		Db:       viper.GetString("mysql.db"),
	}
	return repo, token, dbconf
}
