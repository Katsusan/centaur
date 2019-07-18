package models

//LoginParam 定义了登陆参数
type LoginParam struct {
	UserName    string
	Password    string
	CaptchaID   string
	CaptchaCode string
}

//LoginInfo 定义了登陆信息
type LoginInfo struct {
	UserName  string
	RealName  string
	RoleNames []string
}

//UpdatePwdParam 定义了密码更新请求参数
type UpdatePwdParam struct {
	OldPassword string
	NewPassword string
}

//LogginCaptcha 定义了登陆验证码结构
type LogginCaptcha struct {
	CaptchaID string
}

//LoginToken 定义了登陆令牌信息
type LoginToken struct {
	AccessToken string
	TokenType   string
	ExpireAt    int64
}
