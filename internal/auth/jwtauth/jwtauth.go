package jwtauth

import (
	"time"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/dgrijalva/jwt-go"
)

var jwtauth *JWTAuth

func init() {
	rediscfg := config.GetGlobalConfig().RedisConf()
	jwtcfg := config.GetGlobalConfig().JWTConf()
	s := NewStore(&RedisConfig{
		addr:      rediscfg.Addr,
		password:  rediscfg.Password,
		keyprefix: jwtcfg.RedisPrefix,
	})

	var opts []Option
	opts = append(opts, SetExpireTime(jwtcfg.Expired))
	opts = append(opts, SetSigningKey(jwtcfg.SigningKey))

}

type JWTAuth struct {
	opts  *jwtoptions
	store Storer
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

func NewJWTAuth(s Storer, opts ...Option) *JWTAuth {
	o := DefaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	return &JWTAuth{
		opts:  &o,
		store: s,
	}
}

//GenerateToken 为生成令牌
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

func (auth *JWTAuth) ParseToken(token string) {
	token, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, nil)
}
