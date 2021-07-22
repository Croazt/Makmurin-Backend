package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Croazt/Makmurin/src/routes"
	"github.com/Croazt/Makmurin/src/utils"
	"github.com/Croazt/Makmurin/src/utils/dbConfig"
	"github.com/gin-gonic/gin"
)

type inisiationVar struct {
	l *log.Logger
}

func main() {
	dbConfig.ConnectDatabase()
	init := &inisiationVar{log.New(os.Stdout, "project-pemweb", log.LstdFlags)}

	r := gin.New()
	r.MaxMultipartMemory = 8 << 20
	r.Use(gin.Recovery(), gin.Logger())

	port := os.Getenv("PORT")

	r.Static("/image", "./assets/img/")
	serveMux := routes.Router(r, init.l)

	if port == "" {
		port = "9090"
	}
	s := &http.Server{
		Addr:         ":" + port,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	utils.GraceFullShutdown(init.l, s)
}
