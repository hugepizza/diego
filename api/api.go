package api

import (
	"net"
	"net/http"

	"github.com/ckeyer/diego/api/view"
	"github.com/ckeyer/logrus"
	"github.com/gin-gonic/gin"
)

const (
	API_PREFIX = "/api"
	UI_PREFIX  = "/release"
)

// Serve start http server.
func Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	gs := gin.New()
	gs.Use(MDRecovery(), MDLogger())
	gs.Use(MDCors())

	gs.NoRoute(view.UI())

	apiRoute(gs.Group(API_PREFIX))

	err = http.Serve(lis, gs)
	if err != nil {
		return err
	}

	return nil
}

// apiRoute api router.
func apiRoute(gr *gin.RouterGroup) {
	gr.GET("/path/*path", testH)
}

func testH(ctx *gin.Context) {
	logrus.Infof("path: %s", ctx.Param("path"))
}
