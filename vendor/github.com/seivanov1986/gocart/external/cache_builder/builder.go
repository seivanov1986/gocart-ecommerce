package cache_builder

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/external/asset_manager"
	"github.com/seivanov1986/gocart/external/observer"
	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository"
	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type builder struct {
	hub           repository.Hub
	widgetManager client.WidgetManager
	assetManager  client.AssetManager
}

func NewBuilder(hub repository.Hub, widgetManager client.WidgetManager) *builder {
	return &builder{hub: hub, widgetManager: widgetManager}
}

func (b *builder) RegisterWidget(name string, widget client.Widget) {
	b.widgetManager.Register(name, widget)
}

func (b *builder) SetAssets(assetManager client.AssetManager) {
	b.assetManager = assetManager
}

func (b *builder) Pages(ctx context.Context) ([]client.UrlListRow, error) {
	return []client.UrlListRow{}, nil
}

func (b *builder) Handler(ctx context.Context, pages []client.UrlListRow) error {
	var offsetID int64 = 0

	for {
		fmt.Println(offsetID)
		rows, err := b.hub.SefUrl().ListLimitId(ctx, offsetID)
		if err != nil {
			return err
		}

		if len(rows) == 0 {
			break
		}

		for _, row := range rows {
			err := b.makeObject(ctx, row)
			if err != nil {
				return err
			}
		}

		offsetID = rows[len(rows)-1].ID
	}

	return nil
}

func (b *builder) makeObject(ctx context.Context, row sefurl.SefUrlListLimitIdRow) (err error) {
	var content []byte

	switch row.Type {
	case 1:
		content, err = b.renderPage(ctx, row)
	case 2:
		content, err = b.renderCategory(ctx, row)
	case 3:
		content, err = b.renderProduct(ctx, row)
	default:
		return fmt.Errorf("sefUrl type is not defined")
	}
	if err != nil {
		return err
	}

	fileName := helpers.GetFileNameByUrl(row.Url)
	return helpers.SaveFile("/tmp/cache/"+fileName, bytes.NewReader(content))
}

func (b *builder) renderPage(ctx context.Context, row sefurl.SefUrlListLimitIdRow) ([]byte, error) {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	templateFiles := []string{
		serviceBasePath + "/schemes/templates/layouts/common.html",
	}

	tmpl, err := template.New("common").ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	assetManager := asset_manager.New()
	b.widgetManager.SetAssets(assetManager)

	err = tmpl.ExecuteTemplate(buf, "common", nil)
	if err != nil {
		return nil, err
	}

	content := buf.String()

	reg, _ := regexp.Compile(`{#outertemplate%([A-Za-z_0-9]+)#}`)
	for _, match := range reg.FindAllStringSubmatch(content, -1) {
		match_1 := match[1]
		if len(match) > 2 {
			match_1 = match[2]
		}

		res, _ := b.widgetManager.Render(ctx, match_1)
		if res != nil {
			content = strings.Replace(content, match[0], *res, -1)
		}
	}

	return []byte(content), nil
}

func (b *builder) renderCategory(ctx context.Context, row sefurl.SefUrlListLimitIdRow) ([]byte, error) {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	templateFiles := []string{
		serviceBasePath + "/schemes/templates/layouts/common.html",
	}

	tmpl, err := template.New("common").ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	assetManager := asset_manager.New()
	b.widgetManager.SetAssets(assetManager)

	err = tmpl.ExecuteTemplate(buf, "common", nil)
	if err != nil {
		return nil, err
	}

	content := buf.String()

	reg, _ := regexp.Compile(`{#outertemplate%([A-Za-z_0-9]+)#}`)
	for _, match := range reg.FindAllStringSubmatch(content, -1) {
		match_1 := match[1]
		if len(match) > 2 {
			match_1 = match[2]
		}

		res, _ := b.widgetManager.Render(ctx, match_1)
		if res != nil {
			content = strings.Replace(content, match[0], *res, -1)
		}
	}

	return []byte(content), nil
}

func (b *builder) renderProduct(ctx context.Context, row sefurl.SefUrlListLimitIdRow) ([]byte, error) {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	templateFiles := []string{
		serviceBasePath + "/schemes/templates/layouts/common.html",
	}

	tmpl, err := template.New("common").ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	assetManager := asset_manager.New()
	b.widgetManager.SetAssets(assetManager)

	err = tmpl.ExecuteTemplate(buf, "common", nil)
	if err != nil {
		return nil, err
	}

	content := buf.String()

	reg, _ := regexp.Compile(`{#outertemplate%([A-Za-z_0-9]+)#}`)
	for _, match := range reg.FindAllStringSubmatch(content, -1) {
		match_1 := match[1]
		if len(match) > 2 {
			match_1 = match[2]
		}

		res, _ := b.widgetManager.Render(ctx, match_1)
		if res != nil {
			content = strings.Replace(content, match[0], *res, -1)
		}
	}

	return []byte(content), nil
}
