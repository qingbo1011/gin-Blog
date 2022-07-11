package test

import (
	"gin-Blog/service"
	"log"
	"testing"
)

func TestCleanAllTag(t *testing.T) {
	err := service.CleanAllTag()
	if err != nil {
		log.Fatalln(err)
	}
}
