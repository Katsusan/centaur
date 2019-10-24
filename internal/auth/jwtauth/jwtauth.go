package jwtauth

import (
	"errors"
	"time"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

var JWTentity *JWTAuth

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
	opts = append(opts, SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("InvalidSigningMethod")
		}
		return []byte(DefaultSigningKey), nil
	}))

	switch jwtcfg.SigningMethod {
	case "HS256":
		opts = append(opts, SetSigningMethod(jwt.SigningMethodHS256))
	case "HS384":
		opts = append(opts, SetSigningMethod(jwt.SigningMethodHS384))
	case "HS512":
		opts = append(opts, SetSigningMethod(jwt.SigningMethodHS512))
	default:
		log.Println("SigningMethod not set in jwt config, will be set default HS512")
		opts = append(opts, SetSigningMethod(jwt.SigningMethodHS512))
	}

	JWTentity = NewJWTAuth(s, opts...)
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
	Close()
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

//解析令牌
func (auth *JWTAuth) ParseToken(token string) (*jwt.StandardClaims, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, auth.opts.keyfunc)
	if err != nil {
		return &jwt.StandardClaims{}, err
	}

	return t.Claims.(*jwt.StandardClaims), nil
}

//销毁令牌
func (auth *JWTAuth) DestroyToken() error {

}
