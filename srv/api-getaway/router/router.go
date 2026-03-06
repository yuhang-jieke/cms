package router

import (
	"cms/srv/handler/server"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/content/add", server.ContentAdd)
	r.GET("/get/by", server.GetById)
	r.GET("/search", server.Search)
	return r
}
