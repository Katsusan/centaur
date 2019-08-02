package middleware

import (
	"time"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware will return middleware which handles cross-origin requests.
func CORSMiddleware() gin.HandlerFunc {
	CORScfg := config.GetGlobalConfig().CORS()
	return cors.New(cors.Config{
		AllowOrigins:     CORScfg.AllowOrigins,
		AllowMethods:     CORScfg.AllowMethods,
		AllowHeaders:     CORScfg.AllowHeaders,
		AllowCredentials: CORScfg.AllowCredentials,
		MaxAge:           time.Second * time.Duration(CORScfg.MaxAge),
	})
}
