package api

import "net/http"

type Page404 struct {
}

// NewPage404
func NewPage404() *Page404 {
	return &Page404{}
}

// ServeHTTP ...
func (p *Page404) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("404 hahaha"))
}
