package conf

import "time"

type TokenConf struct {
	TokenHeader     string
	TokenSecret     string
	TokenExpireTime time.Duration
}

func LoadToken() *TokenConf {
	//TokenHeader     string
	//TokenSecret     string
	//TokenExpireTime int
	TokenHeader := Conf.GetString("TOKEN_HEADER")
	if TokenHeader == "" {
		TokenHeader = "Authorization"
	}
	TokenSecret := Conf.GetString("TOKEN_SECRET")
	if TokenSecret == "" {
		TokenSecret = "abcdefghijklmnopqrstuvwxyz"
	}
	TokenExpireTime := Conf.GetInt64("TOKEN_EXPIRE_TIME")
	if TokenExpireTime == 0 {
		TokenExpireTime = 30
	}
	return &TokenConf{
		TokenHeader:     TokenHeader,
		TokenSecret:     TokenSecret,
		TokenExpireTime: time.Duration(TokenExpireTime) * time.Minute,
	}
}
