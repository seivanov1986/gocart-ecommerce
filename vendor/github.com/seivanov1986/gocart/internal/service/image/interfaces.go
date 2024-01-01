package image

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image"
)

type Service interface {
	Create(ctx context.Context, in ImageCreateIn) error
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context, in CategoryListIn) (*image.ImageListOut, error)
}
