package view

import (
	"net/http"
	"strings"

	"github.com/ckeyer/logrus"
	"github.com/gin-gonic/gin"
)

var (
	Index, _ = Asset("index.html")
)

func UI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := strings.TrimPrefix(ctx.Request.URL.Path, "/")

		ctx.Writer.WriteHeader(http.StatusOK)
		ctx.Writer.Header().Add("Content-Encoding", "gzip")

		switch {
		case strings.HasSuffix(path, ".html"):
			ctx.Writer.Header().Set("Content-Type", "text/html")
		case strings.HasSuffix(path, ".js"):
			ctx.Writer.Header().Set("Content-Type", "application/x-javascript")
		case strings.HasSuffix(path, ".css"):
			ctx.Writer.Header().Set("Content-Type", "text/css")
		case strings.HasSuffix(path, ".svg"):
			ctx.Writer.Header().Set("Content-Type", "text/xml")
		case strings.HasSuffix(path, ".jpg"),
			strings.HasSuffix(path, ".jepg"):
			ctx.Writer.Header().Set("Content-Type", "image/jpeg")
		default:
			ctx.Writer.Header().Set("Content-Type", "text/plain")
		}

		body, err := Asset(path)
		if err != nil {
			logrus.Debugf("get path %s failed, use index.", path)
			ctx.Writer.Header().Set("Content-Type", "text/html")
			ctx.Writer.Write(Index)
			return
		}
		logrus.Debugf("get path %s and return.", path)
		ctx.Writer.Write(body)
	}
}
