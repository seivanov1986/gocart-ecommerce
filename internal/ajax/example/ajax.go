package example

import (
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/external/cache"
)

type loggerPlugin struct {
}

func New() *loggerPlugin {
	return &loggerPlugin{}
}

func (l *loggerPlugin) Execute(header http.Header, body io.ReadCloser) (*string, error) {
	cache.Cache.AddEvent()
	result := "example"
	return &result, nil
}
