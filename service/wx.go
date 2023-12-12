package service

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
	"wechat-utils/model"
)

func WXMsgReply(c *gin.Context, fromUser, toUser string) { //微信收发消息
	var res model.WxResponseAP
	//err := global.WxUserColl.FindOne(context.TODO(), bson.M{"openId": toUser}).Decode(&res)
	//if err != nil {
	//	log.Println("查询失败:", err)
	//}
	account := strconv.FormatInt(res.Account, 10)
	text := "Account:" + string(account) + " Password:" + res.Password
	repTextMsg := model.WxResponse{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      text,
	}
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}
