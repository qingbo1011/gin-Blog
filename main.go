package main

import (
	"gin-Blog/conf"
	"gin-Blog/db/mysql"
	"gin-Blog/route"
	"log"
)

func main() {
	r := route.NewRoute()
	err := r.Run(conf.HttpPort)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	conf.Init("./conf/config.ini")
	mysql.MysqlDBInit()
}
