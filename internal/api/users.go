package api

import (
	"net/http"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/Katsusan/centaur/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserByID(router *gin.RouterGroup, conf *config.Config) {
	router.GET("/:userid", func(c *gin.Context) {
		userid := c.Param("userid")
		var user models.UserDetail
		if err := conf.DB().Where("id = ", userid).First(&user).Error; err != nil {
			log.Errorf("failed to get user by id[%s], error=%v\n", userid, err)
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, user)
		}

	})
}
