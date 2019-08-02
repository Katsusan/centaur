package server

import (
	"context"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/Katsusan/centaur/internal/middleware"
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

	router.Use(middleware.TraceMiddleware())

	//是否启用跨域
	if conf.CORS().CORSEnable {
		router.Use(middleware.CORSMiddleware())
	}

	/*
		router.Use(func(ctx *gin.Context) {
			ctx.Header("Content-Type", "application/json")

			if ctx.Request.Method == "OPTIONS" {
				ctx.AbortWithStatus(http.StatusNoContent)
			}
		})
	*/

	g := router.Group("/api/")
}
