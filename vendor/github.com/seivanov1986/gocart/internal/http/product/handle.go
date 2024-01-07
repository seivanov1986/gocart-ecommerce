package product

import (
	"github.com/seivanov1986/gocart/external/cache"
	"github.com/seivanov1986/gocart/internal/service/product"
)

type handle struct {
	service     product.Service
	cacheObject cache.CacheObject
}

func New(service product.Service, cacheObject cache.CacheObject) *handle {
	return &handle{service: service, cacheObject: cacheObject}
}
