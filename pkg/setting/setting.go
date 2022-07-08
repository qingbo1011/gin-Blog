package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	RunMode string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Fatalln("Fail to parse 'conf/app.ini': ", err)
	}

	LoadBase(file)
	LoadServer(file)
	LoadApp(file)

}

func LoadBase(file *ini.File) {
	RunMode = file.Section("").Key("RUN_MODE").MustString("debug") // MustString,默认值为debug
}

func LoadServer(file *ini.File) {
	section, err := file.GetSection("server")
	if err != nil {
		log.Fatalln(err)
	}
	HttpPort = section.Key("HTTP_PORT").MustInt(8080)
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
