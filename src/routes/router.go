package routes

import (
	"log"
	"net/http"

	"github.com/Croazt/Makmurin/src/middlewares"
	"github.com/Croazt/Makmurin/src/utils"
	"github.com/gin-gonic/gin"
)

func Router(serveMux *gin.Engine, l *log.Logger) http.Handler {

	serveMux.Use(middlewares.CORSMiddleware())

	serveMux.GET("/", func(c *gin.Context) {
		utils.SuccessHandler(c, nil, "API IS RUNNING WELL")
	})

	return serveMux
}
