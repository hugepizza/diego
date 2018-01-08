package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ckeyer/logrus"
)

// MDCors middleware for CORS.
func MDCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type,Limit,Offset,Origin,Accept,X-Signature")
		rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Max-Age", fmt.Sprint(24*time.Hour/time.Second))

		if req.Method == "OPTIONS" {
			rw.WriteHeader(http.StatusNoContent)
		}

		next.ServeHTTP(rw, req)
	})
}

// MDLogger middleware for http logger.
func MDLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()

		next.ServeHTTP(rw, req)

		logent := logrus.WithFields(logrus.Fields{
			"Method": req.Method,
			"URL":    req.URL.Path,
			"Remote": req.RemoteAddr,
			// "Status": ctx.Writer.Status(),
			"Period": fmt.Sprintf("%.6f", time.Now().Sub(start).Seconds()),
		})
		logrus.Infof("%#v", rw)

		for _, prefix := range []string{API_PREFIX, UI_PREFIX} {
			if strings.HasPrefix(req.URL.Path, prefix) {
				logent.Info("bye jack.")
				return
			}
		}
	})
}
