package page

import (
	"github.com/seivanov1986/gocart/external/cache"
	"github.com/seivanov1986/gocart/internal/service/page"
)

type handle struct {
	service     page.Service
	cacheObject cache.CacheObject
}

func New(service page.Service, cacheObject cache.CacheObject) *handle {
	return &handle{service: service, cacheObject: cacheObject}
}
