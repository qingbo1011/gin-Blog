package conf

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	RunMode string

	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JwtSecret string
	PageSize  int

	MysqlUser     string
	MysqlPassword string
	MysqlHost     string
	MysqlPort     string
	MysqlName     string
)

func Init(path string) {
	if path == "" { // 如果从main函数处调用，可以无需填写具体的path
		path = "./conf/config.ini"
	}
	file, err := ini.Load(path)
	if err != nil {
		log.Fatalln("Fail to parse 'conf/app.ini': ", err)
	}

	LoadBase(file)
	LoadServer(file)
	LoadApp(file)
	LoadMysql(file)
}

func LoadBase(file *ini.File) {
	RunMode = file.Section("").Key("RUN_MODE").MustString("debug") // MustString,默认值为debug
}

func LoadServer(file *ini.File) {
	section, err := file.GetSection("server")
	if err != nil {
		log.Fatalln(err)
	}
	HttpPort = section.Key("HTTP_PORT").MustString("8080")
	ReadTimeout = time.Duration(section.Key("HTTP_PORT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(section.Key("HTTP_PORT").MustInt(60)) * time.Second
}

func LoadApp(file *ini.File) {
	section, err := file.GetSection("app")
	if err != nil {
		log.Fatalln(err)
	}
	JwtSecret = section.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = section.Key("PAGE_SIZE").MustInt(10)
}

func LoadMysql(file *ini.File) {
	section, err := file.GetSection("mysql")
	if err != nil {
		log.Fatalln(err)
	}
	MysqlUser = section.Key("USER").String()
	MysqlPassword = section.Key("PASSWORD").String()
	MysqlHost = section.Key("HOST").String()
	MysqlPort = section.Key("PORT").String()
	MysqlName = section.Key("NAME").String()
}
