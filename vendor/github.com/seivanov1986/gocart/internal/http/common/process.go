package common

import (
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

func (u *handle) Process(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	bytes, err := u.service.Process(r.Context(), path)
	if err != nil {
		helpers.HttpResponse(w, http.StatusNotFound)
		return
	}

	// TODO if not found - try to render
	u.cacheBuilder.RenderObject(r.Context(), path)

	_, err = w.Write(bytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}
}
