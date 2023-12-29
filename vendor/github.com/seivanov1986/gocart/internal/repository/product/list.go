package product

import (
	"context"
)

type ProductListInput struct {
	Page int64
}

type ProductListOut struct {
	List  []ProductListRow
	Total int64
}

type ProductListRow struct {
	Id        int64   `db:"id" json:"id"`
	Name      string  `db:"name" json:"name"`
	Content   *string `db:"content" json:"content"`
	IdMeta    *int64  `db:"id_meta" json:"id_meta"`
	Sort      int64   `db:"sort" json:"sort"`
	Price     *string `db:"price" json:"price"`
	IdImage   *int64  `db:"id_image" json:"id_image"`
	Disabled  int64   `db:"disabled" json:"disabled"`
	CreatedAt int64   `db:"created_at" json:"created_at"`
	UpdatedAt int64   `db:"updated_at" json:"updated_at"`
}

func (i *repository) List(ctx context.Context, in ProductListInput) (*ProductListOut, error) {
	imageRows := []ProductListRow{}
	err := i.db.SelectContext(
		ctx,
		&imageRows,
		`SELECT    
				p.id, p.name, p.content, p.id_meta, 
				p.sort, p.price, p.id_image, p.disabled,
				p.created_at, p.updated_at
			FROM product p
          LIMIT ?, ?`,
		in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM product`)
	if err != nil {
		return nil, err
	}

	return &ProductListOut{
		List:  imageRows,
		Total: total,
	}, nil
}
