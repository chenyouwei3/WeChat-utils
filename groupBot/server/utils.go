package server

import (
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"log"
	"net/http"
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
