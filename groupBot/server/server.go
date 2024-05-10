package server

import (
	"github.com/eatmoreapple/openwechat"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var VulBot = NewBot{
	Bot: openwechat.DefaultBot(openwechat.Desktop),
}

type NewBot struct {
	Bot *openwechat.Bot
}

func (n NewBot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		data, err := decodeAndMarshalData(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		groups, err := getAllGroup()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var cacheS SentCache
		for _, group := range groups {
			MSG, err := group.SendText(string(data))
			cacheS.time = time.Now()
			cacheS.group = append(cacheS.group, MSG)
			n.FailErrer(err)
		}
		Buckets.cache[string(data)] = cacheS
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		data, err := decodeAndMarshalData(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for _, msg := range Buckets.cache {
			for _, m := range msg.group {
				m.Revoke()
			}
		}
		Buckets.mu.Lock()
		defer Buckets.mu.Unlock()
		delete(Buckets.cache, string(data))
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (n NewBot) HandleReplyGroupMessage(msg *openwechat.Message) {
	err := msg.AsRead()
	n.FailErrer(err)
	if strings.Contains(msg.Content, "@银沙") {
		sender, err := msg.Sender()         //群聊
		sender1, err := msg.SenderInGroup() //提问人
		n.FailErrer(err)
		switch {
		case msg.Content == "@银沙":
			_, err := msg.ReplyText("请您输入对应的数字来确定您的问题(1-5)")
			n.FailErrer(err)
		case strings.Contains(msg.Content, "问题1"):
			_, err := msg.ReplyText("您好,这是问题1的答案")
			n.FailErrer(err)
		case strings.Contains(msg.Content, "问题2"):
			_, err := msg.ReplyText("您好,这是问题2的答案")
			n.FailErrer(err)
		case strings.Contains(msg.Content, "问题3"):
			_, err := msg.ReplyText("您好,这是问题3的答案")
			n.FailErrer(err)
		case strings.Contains(msg.Content, "问题4"):
			_, err := msg.ReplyText("您好,这是问题4的答案")
			n.FailErrer(err)
		case strings.Contains(msg.Content, "问题5"):
			_, err := msg.ReplyText("您好,这是问题5的答案")
			n.FailErrer(err)
		default:
			pushVul(sender.NickName, sender1.NickName, msg.Content)
			_, err := msg.ReplyText("您提问的问题不在我的能力之内,稍后会有人工客服联系您")
			n.FailErrer(err)
		}
	}
}

func (n *NewBot) Login(bot *openwechat.Bot, method string) {
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl // 注册登陆二维码回调
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer func(reloadStorage io.ReadWriteCloser) {
		err := reloadStorage.Close()
		n.FailErrer(err)
		return
	}(reloadStorage)
	switch method {
	case "hot": //热登录
		if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
			log.Println("HotLogin err:", err)
			return
		}
	case "click": //微信点击登录
		if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
			log.Println("ClickLogin err:", err.Error())
			return
		}
	default:
		log.Fatalln("No Login")
	}
}

func (n NewBot) FailErrer(err error) {
	log.Println(err)
}
