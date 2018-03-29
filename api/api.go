package api

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ckeyer/diego/api/view"
	"github.com/ckeyer/diego/storage"
	"github.com/ckeyer/logrus"
	"github.com/gin-gonic/gin"
)

var (
	stogr storage.Storeger
)

const (
	API_PREFIX = "/api"
	UI_PREFIX  = "/release"
)

// Serve start http server.
func Serve(addr string, str storage.Storeger) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	stogr = str

	gs := gin.New()
	gs.NoRoute(view.UI())
	gs.Use(MDRecovery(), MDLogger())
	gs.Use(MDCors())

	apiRoute(gs.Group(API_PREFIX))

	dlRoute(gs.Group(API_PREFIX))

	err = http.Serve(lis, gs)
	if err != nil {
		return err
	}

	return nil
}

// apiRoute api router.
func apiRoute(gr *gin.RouterGroup) {
	// gr.GET("/path/*path", todo)

	gr.GET("/users", ListUsers())
	gr.POST("/users", CreateUser())
	gr.GET("/users/:name", GetUserProfile())
	gr.GET("/users/:name/check", CheckUserName())

	gr.GET("/orgs", ListOrgs())
	gr.POST("/orgs", CreateOrg())
	gr.GET("/orgs/:name", GetOrgProfile())
	gr.GET("/orgs/:name/check", CheckUserName())

	gr.GET("/users/:name/projects", ListProjects())
	gr.GET("/users/:name/projects/:project", GetProjectProfile())
	gr.GET("/orgs/:name/projects", ListProjects())
	gr.GET("/orgs/:name/projects/:project", GetProjectProfile())
}

// dlRoute api router.
func dlRoute(gr *gin.RouterGroup) {
	// gr.GET("/:ns/:ns_name", todo)
}

func todo(ctx *gin.Context) {
	logrus.Infof("path: %s", ctx.Param("path"))
	InternalServerErr(ctx, "todo")
}

func decodeBody(ctx *gin.Context, v interface{}) error {
	return json.NewDecoder(ctx.Request.Body).Decode(v)
}
