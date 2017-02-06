package handlers

import "net/http"

type HelloHandler struct {
}

func (handler HelloHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	RenderJSON(rw, http.StatusOK, map[string]string{"msg": "hello"})
}
