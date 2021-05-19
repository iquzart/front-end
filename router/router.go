package router

import (
	"front-end/controller"

	"github.com/gin-gonic/gin"
)

//func CreateTmplRender() multitemplate.Renderer {

//	baseTmpl := "public/templates/base.tmpl.html"
//	homeTmpl := "public/templates/home.tmpl.html"
//	aboutTmpl := "public/templates/about.tmpl.html"
//	nameTmpl := "public/templates/name.tmpl.html"
//	s404Tmpl := "public/templates/404.tmpl.html"

//	r := multitemplate.NewRenderer()
//	r.AddFromFiles("home", baseTmpl, homeTmpl)
//	r.AddFromFiles("404", baseTmpl, s404Tmpl)
//	r.AddFromFiles("about", baseTmpl, aboutTmpl)
//	r.AddFromFiles("name", baseTmpl, nameTmpl)
//	return r
//}

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Frontend
	r.Static("/static", "./public/static")
	//r.HTMLRender = CreateTmplRender()
	r.LoadHTMLGlob("public/templates/**/*")

	// Routes
	r.GET("/", controller.Home)
	//	r.GET("/about", controller.About)
	//	r.GET("/api", controller.Api)
	//	r.GET("/user/:name", controller.UrlParam)
	r.GET("/health", controller.Health)
	r.NoRoute(controller.NoFound)

	return r
}
