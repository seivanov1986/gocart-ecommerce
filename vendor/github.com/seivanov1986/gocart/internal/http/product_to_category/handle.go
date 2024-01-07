package product_to_category

import (
	"github.com/seivanov1986/gocart/external/cache"
	"github.com/seivanov1986/gocart/internal/service/product_to_category"
)

type handle struct {
	service     product_to_category.Service
	cacheObject cache.CacheObject
}

func New(service product_to_category.Service, cacheObject cache.CacheObject) *handle {
	return &handle{service: service, cacheObject: cacheObject}
}
