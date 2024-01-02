package image

import (
	"bytes"
	"context"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository/image"
)

type ImageCreateIn struct {
	Name       string
	ParentID   int64
	Path       string
	FType      int64
	CreatedAT  int64
	OriginPath string
	UID        string
}

func (u *service) Create(ctx context.Context, in ImageCreateIn) error {
	path := "/"
	if in.ParentID > 0 {
		row, err := u.hub.Image().Read(ctx, in.ParentID)
		if err != nil {
			return err
		}
		path = row.Path
	}

	err := u.hub.Image().Create(ctx, image.ImageCreateInput{
		Name:       in.Name,
		ParentID:   in.ParentID,
		Path:       path,
		FType:      in.FType,
		CreatedAT:  in.CreatedAT,
		OriginPath: in.OriginPath,
		UID:        in.UID,
	})
	if err != nil {
		return err
	}

	return u.makeThumb(in.OriginPath, in.UID, in.Name, path, Size{180, 180})
}

func (u *service) makeThumb(originPath, uid, name, path string, size Size) error {
	filePath := originPath + uid

	img, err := imaging.Open(filePath, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	centerCropping := imaging.Resize(img, size.Width, size.Height, imaging.CatmullRom)
	buf := new(bytes.Buffer)

	err = imaging.Encode(buf, centerCropping, 0)
	if err != nil {
		return err
	}

	fileNameExt := filepath.Ext(name)
	fileName := strings.TrimSuffix(name, fileNameExt)

	return helpers.SaveFile("/tmp/project/images"+path+fileName+"_180x180.jpeg", buf)
}
