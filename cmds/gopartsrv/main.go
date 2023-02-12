package main

import (
	log "github.com/sirupsen/logrus"
	"gopartsrv/public/service"
	_ "gopartsrv/utils/logs"
)



func main() {
	log.WithFields(log.Fields{"info": "这是golang日志框架--logrus"}).Info("描述信息为golang日志框架logrus的学习")
	service.Serviceinit() //启动服务
}
