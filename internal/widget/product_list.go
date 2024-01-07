package widget

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"strings"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/external/observer"
	"github.com/seivanov1986/sql_client"
)

type productListWidget struct {
	assetManager client.AssetManager
	sqlClient    sql_client.DataBase
}

func NewProductList(sqlClient sql_client.DataBase) *productListWidget {
	return &productListWidget{
		sqlClient: sqlClient,
	}
}

func (l *productListWidget) Execute(ctx context.Context, url client.SefUrlItem) (*string, error) {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	layoutFile := []string{
		serviceBasePath + "/schemes/templates/widgets/product_list.html",
	}

	tmpl, err := template.New("common").ParseFiles(layoutFile...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	buf := &bytes.Buffer{}

	productIDs := []int64{}

	err = l.sqlClient.SelectContext(
		ctx,
		&productIDs,
		`SELECT id_product FROM product_to_category WHERE id_category=?`,
		url.IdObject,
	)
	if err != nil {
		return nil, err
	}

	valuesArr := []string{}
	vv := []interface{}{}
	for _, v := range productIDs {
		valuesArr = append(valuesArr, "?")
		vv = append(vv, v)
	}
	valuesStr := strings.Join(valuesArr, ",")

	sql := fmt.Sprintf("SELECT id, name, price FROM product WHERE id IN (%s)", valuesStr)

	var list []struct {
		Id    string `db:"id" json:"id"`
		Name  string `db:"name" json:"name"`
		Price *int64 `db:"price" json:"price"`
	}

	err = l.sqlClient.SelectContext(
		ctx,
		&list,
		sql,
		vv...,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = tmpl.ExecuteTemplate(buf, "product_list", map[string]interface{}{
		"Items": list,
	})
	if err != nil {
		return nil, err
	}

	result := buf.String()
	return &result, nil
}

func (l *productListWidget) SetAssets(assetManager client.AssetManager) {
	l.assetManager = assetManager
}
