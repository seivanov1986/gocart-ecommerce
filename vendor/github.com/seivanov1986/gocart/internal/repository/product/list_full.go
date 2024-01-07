package product

import (
	"context"
)

type ProductListFullInput struct {
	Page int64
}

type ProductListFullOut struct {
	List  []ProductListFullRow
	Total int64
}

type ProductListFullRow struct {
	Id            int64   `db:"id"`
	Name          string  `db:"name"`
	Content       *string `db:"content"`
	IdMeta        *int64  `db:"id_meta"`
	Sort          int64   `db:"sort"`
	Price         *string `db:"price"`
	IdImage       *int64  `db:"id_image"`
	Path          string  `db:"path"`
	SfName        string  `db:"sfname"`
	IPath         string  `db:"ipath"`
	IName         string  `db:"iname"`
	Disabled      int64   `db:"disabled"`
	CreatedAt     int64   `db:"created_at"`
	UpdatedAt     int64   `db:"updated_at"`
	PtcIdCategory int64   `db:"ptc_id_category"`
}

func (i *repository) ListFull(ctx context.Context, in ProductListFullInput) (*ProductListFullOut, error) {
	imageRows := []ProductListFullRow{}
	err := i.db.SelectContext(
		ctx,
		&imageRows,
		`SELECT 
			p.id,  
			p.name, 
			p.content, 
			p.id_meta,
			p.sort, 
			p.price, 
			p.id_image, 
			sf.id as sfid, 
			sf.url, 
			sf.path, 
			sf.name as sfname, 
			sf.type, 
			sf.id_object, 
			sf.template,
			i.id as i_id,
			i.name as i_name,
			i.path as i_path,
			i.ftype as i_ftype,
			ptc.id_category as ptc_id_category
		FROM product p
		LEFT JOIN sefurl sf 
			ON sf.id_object = p.id 
			AND sf.type = 3 
		LEFT JOIN image i
			ON i.id = p.id_image
		LEFT JOIN product_to_category ptc
			ON ptc.id_product = p.id AND ptc.main_category = 1
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

	return &ProductListFullOut{
		List:  imageRows,
		Total: total,
	}, nil
}
