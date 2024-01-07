package example

import (
	"context"

	"github.com/seivanov1986/gocart/client"
)

type loggerPlugin struct {
	assetManager client.AssetManager
}

func New() *loggerPlugin {
	return &loggerPlugin{}
}

func (l *loggerPlugin) Execute(ctx context.Context, url client.SefUrlItem) (*string, error) {
	result := "input Logger is inactive"
	return &result, nil
}

func (l *loggerPlugin) SetAssets(assetManager client.AssetManager) {
	l.assetManager = assetManager
}
