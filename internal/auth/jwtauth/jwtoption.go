package jwtauth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type jwtoptions struct {
	signmethod jwt.SigningMethod
	signingkey string
	expire     int
	tokentype  string
	keyfunc    jwt.Keyfunc
}

type Option func(*jwtoptions)

const DefaultSigningKey = "centaur"

var DefaultOptions = jwtoptions{
	signmethod: jwt.SigningMethodHS512,
	tokentype:  "Bearer",
	signingkey: DefaultSigningKey,
	expire:     3600,
	keyfunc: func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("InvalidToken")
		}
		return []byte(DefaultSigningKey), nil
	},
}

//SetSigningMethod 设定签名方法
func SetSigningMethod(method jwt.SigningMethod) Option {
	return func(o *jwtoptions) {
		o.signmethod = method
	}
}

//SetSigningKey 设定签名secretkey
func SetSigningKey(key string) Option {
	return func(o *jwtoptions) {
		o.signingkey = key
	}
}

//SetKeyfunc 设定验证secretkey的回调函数
func SetKeyfunc(f jwt.Keyfunc) Option {
	return func(o *jwtoptions) {
		o.keyfunc = f
	}
}

//SetExpireTime 设定token过期时间(秒)，默认3600s
func SetExpireTime(expire int) Option {
	return func(o *jwtoptions) {
		o.expire = expire
	}
}
