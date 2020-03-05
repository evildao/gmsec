package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserInfoTbl 用户信息
type UserInfoTbl struct {
	gorm.Model
	Nickname string
	Headurl  string
}

// Oauth2AccessToken token认证
type Oauth2AccessToken struct {
	ID          int    `gorm:"primary_key"`
	AccessToken string `gorm:"unique"`
	TokenType   string
	AppKey      string    // key
	Username    string    // 用户名
	Expires     time.Time // 过期时间
}

// Oauth2ClientTbl client key 信息
type Oauth2ClientTbl struct {
	ID              int `gorm:"primary_key"`
	AppKey          string
	AppSecret       string
	ExpireTime      time.Time // 超时时间
	StrictSign      int       // 是否强制验签:0：用户自定义，1：强制
	StrictVerify    int       // 是否强制验证码:0：用户自定义，1：强制
	TokenExpireTime int       // token过期时间
	Aaa             string
}

// Oauth2RefreshToken 刷新token
type Oauth2RefreshToken struct {
	ID              int    `gorm:"primary_key"`
	RefreshToken    string `gorm:"unique"`
	TokenType       string
	AppKey          string
	Username        string
	Expires         time.Time
	TokenExpireTime int
}

// SignClientTbl client secret info
type SignClientTbl struct {
	ID              int `gorm:"primary_key"`
	AppKey          string
	AppSecret       string
	ExpireTime      time.Time // 超时时间
	StrictSign      int       // 是否强制验签:0：用户自定义，1：强制
	StrictVerify    int       // 是否强制验证码:0：用户自定义，1：强制
	TokenExpireTime int       // token过期时间
}

// UserAccountTbl 用户信息
type UserAccountTbl struct {
	ID          int    `gorm:"primary_key"`
	Account     string `gorm:"unique"`
	Password    string
	AccountType int         // 帐号类型:0手机号，1邮件
	AppKey      string      `gorm:"unique_index:UNIQ_5696AD037D3656A4"` // authbucket_oauth2_client表的id
	UserInfoID  int         `gorm:"unique_index:UNIQ_5696AD037D3656A4;index"`
	UserInfoTbl UserInfoTbl `gorm:"association_foreignkey:user_info_id;foreignkey:id"` // 用户信息
	RegTime     time.Time
	RegIP       string
	BundleID    string
	Describ     string
}