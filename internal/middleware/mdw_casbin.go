package middleware

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	CasbinUserID = "CasbinUID"
)

func Casbin(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetString(CasbinUserID)
		path := c.Request.URL.Path
		method := c.Request.Method
		if isOK, err := e.EnforceSafe(user, path, method); err != nil {
			log.Errorf("Casbin check failed, err:%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "Error occured in CasbinMiddleware, " + err.Error(),
			})
			return
		} else if !isOK {
			log.Debugf("CasbinMiddleware: User don't have the permission, role=[%v], path=[%v], method=[%v]", user, path, method)
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "Sorry you don't have the permission",
			})
			c.Abort()
		}
		c.Next()
	}
}
