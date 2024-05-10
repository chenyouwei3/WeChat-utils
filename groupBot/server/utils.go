package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"log"
	"net/http"
	"time"
)

func decodeAndMarshalData(r *http.Request) ([]byte, error) {
	var vulData interface{}
	if err := json.NewDecoder(r.Body).Decode(&vulData); err != nil {
		return nil, err
	}
	data, err := json.Marshal(vulData)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getAllGroup() ([]*openwechat.Group, error) {
	selfUser, err := VulBot.Bot.GetCurrentUser()
	if err != nil {
		log.Println("获取当前用户失败:", err)
		return nil, err
	}
	groups, err := selfUser.Groups()
	if err != nil {
		log.Println("获取群组失败:", err)
		return nil, err
	}
	return groups, nil
}

func pushVul(group, sender, question string) {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=4c1cd243-6e64-49b1-be04-edd2ed6370f5"
	// 构造消息体
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	sTring := "[" + group + "]" + "-" + "[" + timeNow + "]" + "-" + "[" + question + "]" + "-" + "[" + sender + "]"
	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": sTring,
		},
	}
	// 将消息体转换为 JSON 格式
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}
	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("HTTP request error:", err)
		return
	}
	defer resp.Body.Close()

}
