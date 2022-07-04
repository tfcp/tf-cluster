package test

import (
	"tf-cluster/internal/model"
	"tf-cluster/library/gredis"
	"tf-cluster/library/log"
	"github.com/gogf/gf/frame/g"
)

func SetupServer() {
	g.Config().SetPath("../../../config")
	log.Setup()
	if err := gredis.Setup(); err != nil {
		log.Logger.Fatalf("redis init error:%v", err)
	}
	if err := model.Setup(); err != nil {
		log.Logger.Fatalf("db init error:%v", err)
		return
	}
}
