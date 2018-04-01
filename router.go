package main

import (
	. "BlogArticle/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.GET("/article", GetListAPI)

	router.GET("/article/:id", GetContentApi)

	router.PUT("/article/:id", UpdateArticApi)

	router.DELETE("/article/:id", DelArticApi)

	return router
}
