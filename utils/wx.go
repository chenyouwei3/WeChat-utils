package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"wechat-utils/model"

	"sort"
	"strings"
)

func GetAccessToken(appid, secret string) (string, error) { //获取accessToken
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret
	obj, err := http.Get(url)
	if err != nil {
		return "连接错误", nil
	}
	defer obj.Body.Close()
	body, err := io.ReadAll(obj.Body)
	if err != nil {
		return "读取响应失败", err
	}
	var token model.AccessTokenResponse
	err = json.Unmarshal(body, &token) //非流式传输
	if err != nil {
		return "", err
	}
	return token.AccessToken, err
}

func GetWebAccessToken(appid, secret, code string) (string, string, error) { //获取web端accessToken
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appid + "&secret=" + secret + "&code=" + code + "&grant_type=authorization_code"
	obj, err := http.Get(url)
	if err != nil {
		return "连接错误", "", nil
	}
	defer obj.Body.Close()
	//defer func() {
	//	if cerr := obj.Body.Close(); cerr != nil {
	//		err = fmt.Errorf("关闭响应失败: %s", cerr.Error())
	//	}
	//}()
	body, err := io.ReadAll(obj.Body)
	if err != nil {
		return "读取响应失败", "", err
	}
	var token model.WebAccessTokenResponse
	err = json.Unmarshal(body, &token) //非流式传输
	if err != nil {
		return "", "", err
	}
	return token.AccessToken, token.OpenID, err
}

func GetUserInformation(accessToken, openid string) (model.WechatUserInformation, error) { // 拉取用户信息
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openid + "&lang=zh_CN"
	obj, err := http.Get(url)
	if err != nil {
		return model.WechatUserInformation{}, err
	}
	defer obj.Body.Close()
	body, err := io.ReadAll(obj.Body)
	if err != nil {
		return model.WechatUserInformation{}, err
	}
	var user model.WechatUserInformation
	err = json.Unmarshal(body, &user) //非流式传输
	if err != nil {
		return model.WechatUserInformation{}, err
	}
	return user, err
}

func CheckSignature(signature, timestamp, nonce, token string) bool { //微信服务器证明辅助函数
	arr := []string{timestamp, nonce, token}
	sort.Strings(arr) // 字典序排序
	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// 进行Sha1编码
func Sha1(str string) string { //微信服务器证明辅助函数
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func GetUnionID(accessToken, openid string) (string, error) { //获取unionId
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openid + "&lang=zh_CN"
	obj, err := http.Get(url)
	if err != nil {
		return "连接错误", nil
	}
	defer obj.Body.Close()
	body, err := io.ReadAll(obj.Body)
	if err != nil {
		return "读取响应失败", err
	}
	var unionId model.UnionIdResponse
	err = json.Unmarshal(body, &unionId) //非流式传输
	if err != nil {
		return "", err
	}
	return unionId.UnionID, nil
}
