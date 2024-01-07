package yandex

import (
	"net/http"
)

type Handle interface {
	Feed(w http.ResponseWriter, r *http.Request)
}
