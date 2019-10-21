package jwtauth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtauth *JWTAuth

func init() {

}

type JWTAuth struct {
	opts  *jwtoptions
	store Storer
}

type jwtoptions struct {
	signmethod jwt.SigningMethod
	signingkey string
	expire     int
	tokentype  string
}

type Storer interface {
	//放入token并设置过期时间
	Set(token string, expire time.Duration) error
	//验证token是否存在
	Verify(token string) (bool, error)
	//关闭存储
	Close() error
}

type TokenInfo struct {
	Token     string
	TokenType string
	ExpiredAt int64
}

func (auth *JWTAuth) GenerateToken(userID string) (*TokenInfo, error) {
	now := time.Now()
	expired := now.Add(time.Second * time.Duration(auth.opts.expire)).Unix()

	claims := jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expired,
		Subject:   userID,
		NotBefore: now.Unix(),
	}
	token := jwt.NewWithClaims(auth.opts.signmethod, claims)
	tokenstring, err := token.SignedString(auth.opts.signingkey)
	if err != nil {
		return &TokenInfo{}, err
	}

	tokeinfo := &TokenInfo{
		Token:     tokenstring,
		TokenType: "",
		ExpiredAt: expired,
	}
	return tokeinfo, nil
}
