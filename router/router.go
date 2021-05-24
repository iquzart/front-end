package router

import (
	"front-end/controller"
	"front-end/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Frontend
	r.Static("/static", "public/static")
	//r.HTMLRender = CreateTmplRender()
	r.LoadHTMLGlob("public/templates/**/*")

	authorized := r.Group("/")

	authorized.Use(middleware.Authentication())
	{
		// Secure Routes
		authorized.GET("/", controller.Home)

	}

	r.GET("/login", controller.Login)
	r.POST("/login", controller.Login)
	r.GET("/register", controller.Register)
	r.POST("/register", controller.Register)
	r.GET("/health", controller.Health)
	r.NoRoute(controller.NoFound)

	return r
}
