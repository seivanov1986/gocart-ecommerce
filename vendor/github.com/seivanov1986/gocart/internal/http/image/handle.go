package image

import (
	"github.com/seivanov1986/gocart/internal/service/image"
)

type handle struct {
	service image.Service
}

func New(service image.Service) *handle {
	return &handle{service: service}
}
