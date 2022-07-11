package test

import (
	"gin-Blog/conf"
	"gin-Blog/db/mysql"
	"gin-Blog/service"
	"log"
	"testing"
)

func TestCleanAllTag(t *testing.T) {
	conf.Init("D:\\go_study\\code\\pro\\gin-Blog\\conf\\config.ini")
	mysql.MysqlDBInit()
	err := service.CleanAllTag()
	if err != nil {
		log.Fatalln(err)
	}
}
