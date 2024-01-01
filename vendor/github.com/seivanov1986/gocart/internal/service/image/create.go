package image

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image"
)

type ImageCreateIn struct {
	Name      string
	ParentID  int64
	Path      string
	FType     int64
	CreatedAT int64
}

func (u *service) Create(ctx context.Context, in ImageCreateIn) error {
	return u.hub.Image().Create(ctx, image.ImageCreateInput{
		Name:      in.Name,
		ParentID:  in.ParentID,
		Path:      in.Path,
		FType:     in.FType,
		CreatedAT: in.CreatedAT,
	})
}
