package yandex

import (
	"net/http"

	"github.com/seivanov1986/gocart/external/observer"
	"github.com/seivanov1986/gocart/helpers"
)

func (u *handle) Feed(w http.ResponseWriter, r *http.Request) {
	basePath := observer.GetServiceBasePath(r.Context())
	listOut := u.service.Generate(r.Context(), basePath, "", "", "")

	helpers.HttpResponse(w, http.StatusOK, listOut)
}
