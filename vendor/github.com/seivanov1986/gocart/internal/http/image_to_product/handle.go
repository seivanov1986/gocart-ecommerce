package image_to_product

import (
	"github.com/seivanov1986/gocart/internal/service/image_to_product"
)

type handle struct {
	service image_to_product.Service
}

func New(service image_to_product.Service) *handle {
	return &handle{service: service}
}
