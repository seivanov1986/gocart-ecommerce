package category

import (
	"github.com/seivanov1986/gocart/external/cache"
	"github.com/seivanov1986/gocart/internal/service/category"
)

type handle struct {
	service     category.Service
	cacheObject cache.CacheObject
}

func New(service category.Service, cacheObject cache.CacheObject) *handle {
	return &handle{service: service, cacheObject: cacheObject}
}
