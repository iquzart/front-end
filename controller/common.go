package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	ctx.String(
		http.StatusOK,
		"Working!")
}

func NoFound(ctx *gin.Context) {
	//	ctx.HTML(
	//		404,
	//		"404",
	//		gin.H{
	//			"title": "404",
	//		},
	//	)
	ctx.String(
		http.StatusNotFound,
		"Page not found!")

}
