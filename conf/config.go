package conf

import (
	"fmt"
	"golang/TodoList/model"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	// 打印当前工作目录（关键调试信息）
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败：", err)
		return
	}
	fmt.Println("当前工作目录：", cwd) // 运行后看这里的输出
	file, err := ini.Load("./TodoList/conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败,请检查文件路径", err)
		return
	}
	LoadServer(file)
	LoadMysql(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	model.Database(path)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
