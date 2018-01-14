package api

import (
	"github.com/ckeyer/diego/api/view"
	"github.com/gin-gonic/gin"
)

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		return view.UI()
	}
}
