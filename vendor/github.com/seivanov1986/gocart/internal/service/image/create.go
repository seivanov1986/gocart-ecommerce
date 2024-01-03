package image

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
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
	Total      int64
	Offset     int64
	Body       io.ReadCloser
}

var (
	ErrNotDone = fmt.Errorf("error not done")
)

func (u *service) Create(ctx context.Context, in ImageCreateIn) error {
	filePath := in.OriginPath + in.UID

	if in.Offset == 0 {
		if helpers.IsExists(filePath) {
			return fmt.Errorf("file exists")
		}

		// TODO start go rutine monitor for delete phantom
	}

	err := os.MkdirAll(in.OriginPath, 0777)
	if err != nil {
		return fmt.Errorf("mkdir", in.OriginPath, err.Error())
	}

	if !helpers.IsExists(filePath) && in.Offset > 0 {
		return fmt.Errorf("file exists")
	}

	p, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer p.Close()

	_, err = io.Copy(p, in.Body)
	if err != nil {
		return err
	}

	path := "/"
	if in.ParentID > 0 {
		row, err := u.hub.Image().Read(ctx, in.ParentID)
		if err != nil {
			return err
		}
		path = row.Path
	}

	fmt.Println("/tmp/project/images" + path + in.Name)

	if helpers.IsExists("/tmp/project/images" + path + in.Name) {
		return fmt.Errorf("file exists")
	}

	if in.Offset < in.Total {
		return ErrNotDone
	}

	err = u.hub.Image().Create(ctx, image.ImageCreateInput{
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

	format, err := imaging.FormatFromFilename(name)
	if err != nil {
		return err
	}

	err = imaging.Encode(buf, centerCropping, format)
	if err != nil {
		return err
	}

	fileNameExt := filepath.Ext(name)
	fileName := strings.TrimSuffix(name, fileNameExt)

	return helpers.SaveFile("/tmp/project/images"+path+fileName+"_180x180"+fileNameExt, buf)
}
