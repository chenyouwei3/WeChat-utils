package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"wechat-utils/global"
	"wechat-utils/model"
	"wechat-utils/service"
	"wechat-utils/utils"
)

func WxReceiveMess(c *gin.Context) {
	var data model.WxReceiveMsg
	err := c.ShouldBindXML(&data)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}
	switch data.Event {
	case "subscribe": //订阅
		//c.JSON(http.StatusOK, service.CreateWxUser(data))
	case "unsubscribe": //取消订阅
		//c.JSON(http.StatusOK, service.FollowWxUser(data))
	case "CLICK": //点击事件
		switch data.EventKey {
		case "MeiKou":
			service.WXMsgReply(c, data.ToUserName, data.FromUserName)
		}
	case "VIEW": //url跳转事件
		switch data.EventKey {
		case "https://zouzh.cn/":

		}
	}
}
func WxCheckSignature(c *gin.Context) { //验证服务器
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	ok := utils.CheckSignature(signature, timestamp, nonce, global.WxToken)
	if !ok {
		log.Println("微信公众号接入校验失败!")
		return
	}
	log.Println("微信公众号接入校验成功!")
	_, _ = c.Writer.WriteString(echostr) // 写入响应体返回
}

func GetMenu(c *gin.Context) { //查询菜单
	appToken, err := utils.GetAccessToken(global.AppId, global.AppSecret)
	if err != nil {
		log.Println("获取accessToken失败", err.Error())
	}
	url := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=" + appToken
	obj, err := http.Get(url)
	if err != nil {
		log.Println("url打开失败", err.Error())
	}
	defer obj.Body.Close()
	body, err := io.ReadAll(obj.Body)
	if err != nil {
		log.Println("获取菜单失败", err.Error())
		return
	}
	c.JSON(http.StatusOK, string(body))
}
