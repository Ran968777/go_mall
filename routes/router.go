package routes

import (
	"go-mall/api"
	"go-mall/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置所有路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用数据库中间件
	r.Use(middleware.DB())

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 轮播图相关路由
	r.GET("/indexImgs", api.ListIndexImg)

	r.GET("/shop/notice/topNoticeList", api.ListNotice)

	r.GET("/prod/tag/prodTagList", api.ListTag)

	return r
}
