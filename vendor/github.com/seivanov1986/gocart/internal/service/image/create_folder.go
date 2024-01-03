package image

import (
	"context"
	"fmt"
	"os"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository/image"
)

type ImageCreateFolderIn struct {
	Name     string
	ParentID int64
}

func (u *service) CreateFolder(ctx context.Context, in ImageCreateFolderIn) error {
	path := "/" + in.Name + "/"
	if in.ParentID > 0 {
		row, err := u.hub.Image().Read(ctx, in.ParentID)
		if err != nil {
			return err
		}

		path = row.Path + in.Name + "/"
	}

	filePath := "/tmp/project/images" + path

	if helpers.IsExists(filePath) {
		return fmt.Errorf("file exists")
	}

	err := os.MkdirAll(filePath, 0777)
	if err != nil {
		return err
	}

	return u.hub.Image().Create(ctx, image.ImageCreateInput{
		Name:     in.Name,
		ParentID: in.ParentID,
		Path:     path,
		FType:    1,
	})
}
