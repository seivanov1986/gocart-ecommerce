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
	// GET path from parent

	err := u.hub.Image().Create(ctx, image.ImageCreateInput{
		Name:      in.Name,
		ParentID:  in.ParentID,
		Path:      "/",
		FType:     in.FType,
		CreatedAT: in.CreatedAT,
	})
	if err != nil {
		return err
	}

	return u.makeThumb(ctx, in.Path, in.Name)
}

func (u *service) makeThumb(ctx context.Context, path, name string) error {
	/*
		img, err := imaging.Open(filePath, imaging.AutoOrientation(true))
		if err != nil {
			return fmt.Errorf("error open file: %v, error %w", filePath, err)
		}

		format, _ := imaging.FormatFromFilename(filePath)

		centercroping := imaging.Resize(img, size.Width, size.Height, imaging.CatmullRom)
		buf := new(bytes.Buffer)

		if size.Watermark {
			mark := watermark(serviceBasePath, centercroping, textWatermark)
			imaging.Encode(buf, mark, format)
		} else {
			imaging.Encode(buf, centercroping, format)
		}

		err = u.objectStorage.File().Save("/images"+imagePath+fileName, buf)
	*/

	return nil
}
