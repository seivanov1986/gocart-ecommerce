package common

import (
	"github.com/seivanov1986/gocart/external/cache_builder"
	"github.com/seivanov1986/gocart/internal/service/common"
)

type handle struct {
	service      common.Service
	cacheBuilder cache_builder.Service
}

func New(service common.Service, cacheBuilder cache_builder.Service) *handle {
	return &handle{
		service:      service,
		cacheBuilder: cacheBuilder,
	}
}
