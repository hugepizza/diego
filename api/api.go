package api

import (
	"net"
	"net/http"

	"github.com/ckeyer/commons/version"
	"github.com/ckeyer/logrus"
	"github.com/gorilla/mux"
)

// Serve start http server.
func Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	m := mux.NewRouter()
	m.KeepContext = true

	apiRoute(m.PathPrefix("/api/").Subrouter())
	uiRoute(m.PathPrefix("/ui/").Subrouter())

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

func GetVersion(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(version.GetVersion()))
	logrus.Infof("get Version")
}
