package router

import (
	v1 "ginblog/api/v1"
	"ginblog/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块接口
		router.POST("/user/add", v1.AddUser)
		router.GET("/users", v1.GetUsers)
		router.PUT("/user/:id", v1.EditUser)
		router.DELETE("/user/:id", v1.DeleteUser)
		// 分类模块接口

		// 文章模块接口
	}

	_ = r.Run(config.HttpPort)
}
