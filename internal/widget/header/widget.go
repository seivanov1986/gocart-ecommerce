package header_widget

import (
	"context"

	"github.com/seivanov1986/gocart/client"
)

type widget struct {
	assetManager client.AssetManager
}

func New() *widget {
	return &widget{}
}

func (l *widget) Execute(ctx context.Context, url client.SefUrlItem) (*string, error) {
	if l.assetManager == nil {
		return nil, nil
	}

	l.assetManager.AddCssList([]client.AssetItemDependency{
		{Path: "/static/css/master.min.css"},
	})

	return nil, nil
}

func (l *widget) SetAssets(assetManager client.AssetManager) {
	l.assetManager = assetManager
}
