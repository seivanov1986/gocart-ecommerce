package image

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in ImageCreateInput) error
	Delete(ctx context.Context, idImage int64) error
	List(ctx context.Context, in ImageListInput) (*ImageListOut, error)
	Read(ctx context.Context, id int64) (*ImageReadRow, error)
}
