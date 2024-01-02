package image

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image"
)

type ImageCreateFolderIn struct {
	Name     string
	ParentID int64
}

func (u *service) CreateFolder(ctx context.Context, in ImageCreateFolderIn) error {
	return u.hub.Image().Create(ctx, image.ImageCreateInput{
		Name:     in.Name,
		ParentID: in.ParentID,
		Path:     "/",
		FType:    1,
	})
}
