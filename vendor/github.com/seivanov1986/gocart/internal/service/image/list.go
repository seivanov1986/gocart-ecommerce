package image

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image"
)

type CategoryListIn struct {
	Page     int64
	ParentID int64
}

func (u *service) List(ctx context.Context, in CategoryListIn) (*image.ImageListOut, error) {
	return u.hub.Image().List(ctx, image.ImageListInput{
		Page:     in.Page,
		ParentID: in.ParentID,
	})
}
