package router

import (
	"github.com/gin-gonic/gin"
	"wechat-utils/controller"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()
	engine.GET("/wx", controller.WxCheckSignature)   //微信签名
	engine.POST("/api/wx", controller.WxReceiveMess) //微信接收消息
	return engine
}
