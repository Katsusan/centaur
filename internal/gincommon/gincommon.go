package gincommon

import (
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
