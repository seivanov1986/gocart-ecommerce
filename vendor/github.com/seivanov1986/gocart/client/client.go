package client

import (
	"context"
	"io"
	"net/http"
	"time"
)

type WidgetManager interface {
	Render(ctx context.Context, name string, url SefUrlItem) (*string, error)
	Register(name string, widget Widget)
	SetAssets(assetManager AssetManager)
}

type Widget interface {
	Execute(ctx context.Context, url SefUrlItem) (*string, error)
	SetAssets(assetManager AssetManager)
}

type AjaxHandler interface {
	Execute(header http.Header, body io.ReadCloser) (*string, error)
}

type AjaxManager interface {
	RegisterPath(name string, widget AjaxHandler)
	Handler(w http.ResponseWriter, r *http.Request)
}

type UrlListRow struct {
	ID       int64
	Url      string
	Path     string
	Name     string
	Type     int64
	IdObject int64
}

type CacheBuilder interface {
	Pages(ctx context.Context) ([]UrlListRow, error)
	Handler(ctx context.Context, pages []UrlListRow) error
	RegisterWidget(name string, widget Widget)
}

type SessionManager interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Exists(keys ...string) (bool, error)
	Del(keys ...string) (bool, error)
}

type AssetOption struct {
	Sort    int64
	Type    string
	Preload bool
}

type AssetItemDependency struct {
	Path       string
	Type       string
	Preload    bool
	Dependency []string
}

type ResultList struct {
	Path    string
	Time    int64
	Type    string
	Preload bool
}

type AssetManager interface {
	AddJsList(pathList []AssetItemDependency)
	AddCssList(pathList []AssetItemDependency)
	AddPreloadList(pathList []AssetItemDependency)
	GetJsList() []ResultList
	GetJsTemplate() string
	GetCssList() []ResultList
	GetCssTemplate() string
	GetPreloadList() []ResultList
	GetPreloadTemplate() string
}

type SefUrlItem struct {
	ID       int64   `db:"id" json:"id"`
	Url      string  `db:"url" json:"url"`
	Path     string  `db:"path" json:"path"`
	Name     string  `db:"name" json:"name"`
	Type     int64   `db:"type" json:"type"`
	IdObject int64   `db:"id_object" json:"id_object"`
	Template *string `db:"template" json:"template"`
}
