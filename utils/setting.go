package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPost string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置信息错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPost = file.Section("server").Key("HttpPost").MustString(":3000")
}

func LoadData(file *ini.File) {
	Db = file.Section("db").Key("Db").MustString("mysql")
	DbHost = file.Section("db").Key("DbHost").MustString("localhost")
	DbPort = file.Section("db").Key("DbPort").MustString("3306")
	DbUser = file.Section("db").Key("DbUser").MustString("root")
	DbPassWord = file.Section("db").Key("DbPassWord").MustString("123456")
	DbName = file.Section("db").Key("DbName").MustString("goweb")
}
