package main

import (
	"GroupChatBot/server"
	"github.com/eatmoreapple/openwechat"
	"log"
	"net/http"
	"time"
)

func main() {
	//hot click 两个模式
	server.VulBot.Login(server.VulBot.Bot, "click")
	//http服务
	go func() {
		http.Handle("/vul", &server.NewBot{})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	//启动撤回监听池
	go server.Buckets.StartExpirationCheck(time.Minute * 1)
	//正式处理
	server.VulBot.Bot.MessageHandler = func(msg *openwechat.Message) {
		server.VulBot.HandleReplyGroupMessage(msg)
	}
	select {}
}
