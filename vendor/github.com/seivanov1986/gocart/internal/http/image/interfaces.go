package image

import (
	"net/http"
)

type Handler interface {
	Upload(w http.ResponseWriter, r *http.Request)
	CreateFolder(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}
