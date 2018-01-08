package api

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ckeyer/commons/version"
	"github.com/ckeyer/logrus"
	"github.com/ckeyer/mux"
)

const (
	API_PREFIX = "/api"
	UI_PREFIX  = "/ui"
)

// Serve start http server.
func Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	m := mux.NewRouter()
	m.KeepContext = true

	m.AddMiddlewareFunc(MDCors)
	m.AddMiddlewareFunc(MDLogger)
	apiRoute(m.PathPrefix(API_PREFIX).Subrouter())
	uiRoute(m.PathPrefix(UI_PREFIX).Subrouter())

	err = http.Serve(lis, m)
	if err != nil {
		return err
	}

	return nil
}

// apiRoute api router.
func apiRoute(m *mux.Router) {
	m.NotFoundHandler = NewPage404()
	m.Path("/version").HandlerFunc(GetVersion)
}

// apiRoute api router.
func uiRoute(m *mux.Router) {
	m.NotFoundHandler = NewPage404()
}

// GetVersion return api version.
func GetVersion(rw http.ResponseWriter, req *http.Request) {
	logrus.Infof("get Version")
	JSON(rw, map[string]string{"version": version.GetVersion()})
}

func JSON(rw http.ResponseWriter, data interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		logrus.Errorf("json format failed, %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
