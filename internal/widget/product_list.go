package widget

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/external/observer"
)

type productListWidget struct {
	assetManager client.AssetManager
}

func NewProductList() *productListWidget {
	return &productListWidget{}
}

func (l *productListWidget) Execute(ctx context.Context) (*string, error) {
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

	err = tmpl.ExecuteTemplate(buf, "product_list", map[string]interface{}{
		"Items": []int64{1, 2, 3, 4, 5},
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := buf.String()
	return &result, nil
}

func (l *productListWidget) SetAssets(assetManager client.AssetManager) {
	l.assetManager = assetManager
}
