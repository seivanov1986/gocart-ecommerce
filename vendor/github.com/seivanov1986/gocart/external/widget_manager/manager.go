package widget_manager

import (
	"context"
	"fmt"

	"github.com/seivanov1986/gocart/client"
)

type widgetManager struct {
	widgets      map[string]client.Widget
	assetManager client.AssetManager
}

func New() *widgetManager {
	return &widgetManager{
		widgets: map[string]client.Widget{},
	}
}

func (w *widgetManager) SetAssets(assetManager client.AssetManager) {
	w.assetManager = assetManager
}

func (w *widgetManager) Register(name string, widget client.Widget) {
	w.widgets[name] = widget
}

func (w *widgetManager) Render(ctx context.Context, name string, url client.SefUrlItem) (*string, error) {
	widget, ok := w.widgets[name]
	if !ok {
		return nil, fmt.Errorf("widget not found")
	}

	if w.assetManager != nil {
		widget.SetAssets(w.assetManager)
	}

	return widget.Execute(ctx, url)
}
