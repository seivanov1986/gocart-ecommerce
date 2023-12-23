package sefurl

import (
	"github.com/seivanov1986/gocart/internal/service/sefurl"
)

type handle struct {
	service sefurl.Service
}

func New(service sefurl.Service) *handle {
	return &handle{service: service}
}
