package controllers

import (
	"net/http"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_CAPTCHAID_LENGTH = 4
)

type HTTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginCaptcha struct {
	CaptchaID string `json:"captcha_id"`
}

//GetCaptchaID	获取验证码ID
//@Summary 获取验证码ID
//@Success 200 LoginCaptcha
//@Router GET /api/v1/login/captchaid
func GetCaptchaID(c *gin.Context) {
	captchaID := captcha.NewLen(DEFAULT_CAPTCHAID_LENGTH)

	loginCaptchaID := &LoginCaptcha{CaptchaID: captchaID}
	c.JSON(http.StatusOK, loginCaptchaID)
}

//GetCaptchaImage 获取验证码图片
//@Summary 获取验证码图片
//@Success 200
func GetCaptchaImage(c *gin.Context) {
	captchaID := c.Query("id")
	if captchaID == "" {
		c.JSON(http.StatusBadRequest, &HTTPResponse{Code: 0, Message: "无效的请求参数"})
	}

	captchaConf := config.GetGlobalConfig().CaptchaConf()
	err := captcha.WriteImage(c.Writer, captchaID, captchaConf.Width, captchaConf.Height)
}
