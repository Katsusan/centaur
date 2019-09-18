package controllers

import (
	"net/http"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	DEFAULT_CAPTCHAID_LENGTH = 4
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type LoginCaptcha struct {
	CaptchaID string `json:"captcha_id"`
}

type LoginParam struct {
	UserName    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

type LoginToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
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
//@Success 200 (PNG)Image 图形验证码
//@Failure 400 ErrorResponse {Code:"CaptchaIDNotProvided", Message:"未提供参数captcha_id"}
//@Failure 400 ErrorResponse {Code:"CaptchaIDNotFound", Message:"无效的captchaID"}
//@Failure 500 ErrorResponse {Code:"CaptchaGenerationFail", Message:"生成验证码图片时发生错误"}
func GetCaptchaImage(c *gin.Context) {
	captchaID := c.Query("captcha_id")
	if captchaID == "" {
		c.JSON(http.StatusBadRequest, &ErrorResponse{Code: "CaptchaIDNotProvided", Message: "未提供参数captcha_id"})
	}

	captchaConf := config.GetGlobalConfig().CaptchaConf()
	if captchaConf.Width == 0 {
		captchaConf.Width = captcha.StdWidth
	}
	if captchaConf.Height == 0 {
		captchaConf.Height = captcha.StdHeight
	}

	err := captcha.WriteImage(c.Writer, captchaID, captchaConf.Width, captchaConf.Height)
	if err != nil {
		if err == captcha.ErrNotFound {
			c.JSON(http.StatusBadRequest, &ErrorResponse{Code: "CaptchaIDNotFound",
				Message: "无效的captchaID",
				Detail:  "please use /api/v1/login/captchaid to fetch captchaID first"})
		}
		log.Println("Error occured when generating captcha image, error=", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorResponse{Code: "CaptchaGenerationFail", Message: "生成验证码图片时发生错误"})
	}
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Content-Type", "image/png")
}

//Login 用户登录
//@Summary 验证用户提交
//@Success 200 LoginToken {Code: "LoginOK", LoginToken: {access_token:"", token_type:"", expires_at:""}}
//@Failure 400 ErrorResponse {Code:"ParameterParsingFail", Message:"需提交Username,Password,CaptchaID,CaptchaCode" }
//@Failure 200 ErrorResponse {Code: "CaptchaCodeNotCorrect", Message: "请输入正确的验证码"}
//@Failure 200 ErrorResponse {Code: "UserNameOrPasswordNotCorrect", Message:"请输入正确的用户名或密码"}
//@Failure 500 ErrorResponse {Code:"", Message:""}
// @Router POST /api/v1/login
func Login(c *gin.Context) {
	var param LoginParam

	//确认提交参数
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse{Code: "ParameterParsingFail", Message: "需提交Username,Password,CaptchaID,CaptchaCode"})
	}

	//检查验证码
	if !captcha.VerifyString(param.CaptchaID, param.CaptchaCode) {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorResponse{Code: "CaptchaCodeNotCorrect", Message: "请输入正确的验证码"})
	}

	//确认是否是超级管理员账户
	spadmin := GetRootUser()
	if param.UserName == spadmin.Username {
		if err := bcrypt.CompareHashAndPassword([]byte(spadmin.Password), []byte(param.Password)); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, &ErrorResponse{Code: "UserNameOrPasswordNotCorrect", Message: "请输入正确的用户名或密码"})
		}
	}

}
