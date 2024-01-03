package image_to_product

import (
	"context"
)

type ImageToProductListInput struct {
	Page      int64
	ProductID int64
}

type ImageToProductListOut struct {
	List  []ImageToProductListRow
	Total int64
}

type ImageToProductListRow struct {
	ID        int64  `db:"id"`
	ImageID   int64  `db:"id_image"`
	ImagePath string `db:"path_image"`
}

func (i *repository) List(ctx context.Context, in ImageToProductListInput) (*ImageToProductListOut, error) {
	imageRows := []ImageToProductListRow{}
	err := i.db.SelectContext(
		ctx,
		&imageRows,
		i.getQuery(),
		in.ProductID, in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM image_to_product WHERE id_product = ?`,
		in.ProductID)
	if err != nil {
		return nil, err
	}

	return &ImageToProductListOut{
		List:  imageRows,
		Total: total,
	}, nil
}

func (i *repository) getQuery() string {
	query := ""

	switch i.db.GetDB().DriverName() {
	case "sqlite3":
		query = `SELECT itc.id, itc.id_image, i.path || i.name as path_image
					FROM image_to_product itc
					LEFT JOIN image i ON i.id = itc.id_image
					WHERE id_product = ? LIMIT ?, ?`
	case "mysql":
		query = `SELECT itc.id, itc.id_image, CONCAT(i.path, i.name) as path_image
					FROM image_to_product itc
					LEFT JOIN image i ON i.id = itc.id_image
					WHERE id_product = ? LIMIT ?, ?`
	}

	return query
}
