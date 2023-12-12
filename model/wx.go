package model

import (
	"encoding/xml"
)

type WxReceiveMsg struct { //接收微信xml数据包
	Id           int64
	ToUserName   string //开发者
	FromUserName string //发送方
	CreateTime   int
	MsgType      string
	Event        string
	EventKey     string
}

type WxResponse struct { //向微信返回xml数据包
	ToUserName   string //接收
	FromUserName string //开发者
	CreateTime   int64
	MsgType      string
	Content      string
	XMLName      xml.Name `xml:"xml"`
}

type WxResponseAP struct {
	Account  int64
	Password string
}

type AccessTokenResponse struct { //获取AccessToken
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type WebAccessTokenResponse struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	OpenID         string `json:"openid"`
	Scope          string `json:"scope"`
	IsSnapshotuser int    `json:"is_snapshotuser"`
	UnionID        string `json:"unionid"`
}

type WechatUserInformation struct {
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionID    string   `json:"unionid"`
}

type UnionIdResponse struct { //获取UnionId
	Subscribe     int    `json:"subscribe"`
	OpenID        string `json:"openid"`
	Language      string `json:"language"`
	SubscribeTime int64  `json:"subscribe_time"`
	UnionID       string `json:"unionid"`
	Remark        string `json:"remark"`
	GroupID       int    `json:"groupid"`
	TagIDList     []int  `json:"tagid_list"`
	Scene         string `json:"subscribe_scene"`
	QRScene       int    `json:"qr_scene"`
	QRSceneString string `json:"qr_scene_str"`
}
