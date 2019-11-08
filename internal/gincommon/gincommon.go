package gincommon

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	UserIDString = "user_id"
)

func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDString, userID)
}

func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDString)
}

func GetToken(c *gin.Context) string {
	var jwttoken string
	authHeader := c.GetHeader("Autherization")
	jwtPrefix := "Bearer "
	if authHeader != "" && strings.HasPrefix(authHeader, jwtPrefix) {
		jwttoken = authHeader[len(jwtPrefix):]
	}
	return jwttoken
}
