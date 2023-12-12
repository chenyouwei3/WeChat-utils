package initialize

import (
	"bytes"
	"log"
	"net/http"
	"wechat-utils/global"
	"wechat-utils/utils"
)

func WxInit() {
	AccessToken, err := utils.GetAccessToken(global.AppId, global.AppSecret)
	if err != nil {
		log.Fatalf("获取accessToken失败：%s", err.Error())
		return
	}
	url := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + AccessToken
	res, err := http.Post(url, "application/json", bytes.NewBuffer(global.Menu))
	if err != nil {
		log.Fatalf("发送菜单失误：%s", err.Error())
		return
	}
	defer res.Body.Close()

}
