package yandex

import (
	"github.com/seivanov1986/gocart/internal/service/yandex_feed"
)

type handle struct {
	service yandex_feed.Service
}

func New(service yandex_feed.Service) *handle {
	return &handle{
		service: service,
	}
}
