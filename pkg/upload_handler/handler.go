package upload_handler

import (
	"io"
	"net/http"
)

type handler struct {
}

type Handler interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

func New() *handler {
	return &handler{}
}

func (a *handler) Upload(w http.ResponseWriter, r *http.Request) {
	io.ReadAll(r.Body)
	w.Write([]byte("{}"))
}
