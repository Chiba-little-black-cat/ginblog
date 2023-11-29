package router

import (
	v1 "ginblog/api/v1"
	"ginblog/config"
	"ginblog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	auth := r.Group("api/v1")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.PUT("/user/:id", v1.EditUser)
		auth.DELETE("/user/:id", v1.DeleteUser)

		auth.POST("/category/add", v1.AddCategory)
		auth.PUT("/category/:id", v1.EditCategory)
		auth.DELETE("/category/:id", v1.DeleteCategory)

		auth.POST("/article/add", v1.AddArticle)
		auth.PUT("/article/:id", v1.EditArticle)
		auth.DELETE("/article/:id", v1.DeleteArticle)
	}

	router := r.Group("api/v1")
	{
		router.GET("/users", v1.GetUsers)
		router.POST("/user/add", v1.AddUser)

		router.GET("/categories", v1.GetCategories)

		router.POST("/login", v1.Login)

		router.GET("/article/:id", v1.GetArticle)
		router.GET("/articles", v1.GetArticles)
		router.GET("/articles/:id", v1.GetArticlesByCategoryId)

		router.POST("/upload", v1.UpLoad)
	}

	_ = r.Run(config.HttpPort)
}
