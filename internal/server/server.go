package server

import (
	"context"
	"net/http"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/gin-gonic/gin"
)

func Start(c context.Context, conf *config.Config) {
	if conf.HttpServerMode() != "" {
		gin.SetMode(conf.HttpServerMode())
	} else if !conf.Debug() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.Static("/", conf.StaticPath())

	router.Use(func(ctx *gin.Context) {
		ctx.Header("Content-Tyoe", "application/json")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	})

	v := router.Group("/api/")
}
