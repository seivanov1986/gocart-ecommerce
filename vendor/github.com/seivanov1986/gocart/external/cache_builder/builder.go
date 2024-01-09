package cache_builder

import (
	"bytes"
	"context"
	"fmt"
	htmlTemplate "html/template"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/external/asset_manager"
	"github.com/seivanov1986/gocart/external/observer"
	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository"
	"github.com/seivanov1986/gocart/internal/repository/page"
	"github.com/seivanov1986/gocart/internal/repository/product"
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
		rows, err := b.hub.SefUrl().ListLimitId(ctx, offsetID)
		if err != nil {
			return err
		}

		if len(rows) == 0 {
			break
		}

		for _, row := range rows {
			content, err := b.makeObject(ctx, client.SefUrlItem{
				ID:       row.ID,
				Url:      row.Url,
				Path:     row.Path,
				Name:     row.Name,
				Type:     row.Type,
				IdObject: row.IdObject,
				Template: row.Template,
			})
			if err != nil {
				return err
			}

			fileName := helpers.GetFileNameByUrl(row.Url)
			err = helpers.SaveFile("/tmp/cache/"+fileName, bytes.NewReader(content))
			if err != nil {
				return err
			}
		}

		offsetID = rows[len(rows)-1].ID
	}

	return nil
}

func (b *builder) RenderObject(ctx context.Context, url string) ([]byte, error) {
	row, err := b.hub.SefUrl().Read(ctx, url)
	if err != nil {
		return nil, err
	}

	return b.makeObject(ctx, client.SefUrlItem{
		ID:       row.ID,
		Url:      row.Url,
		Path:     row.Path,
		Name:     row.Name,
		Type:     row.Type,
		IdObject: row.ObjectID,
		Template: row.Template,
	})
}

func (b *builder) makeObject(ctx context.Context, row client.SefUrlItem) ([]byte, error) {
	var content []byte
	var err error

	switch row.Type {
	case 1:
		content, err = b.renderPage(ctx, row)
	case 2:
		content, err = b.renderCategory(ctx, row)
	case 3:
		content, err = b.renderProduct(ctx, row)
	default:
		return nil, fmt.Errorf("sefUrl type is not defined")
	}
	if err != nil {
		return nil, err
	}

	return content, err
}

func (b *builder) getPageData(ctx context.Context, idObject int64) map[string]interface{} {
	result := map[string]interface{}{}

	pageRow, err := b.hub.Page().Read(ctx, page.PageReadInput{ID: idObject})
	if err != nil {
		return result
	}

	result["Year"] = time.Now().Year()

	result["Name"] = pageRow.Name
	content := ""
	if pageRow.Content != nil {
		content = *pageRow.Content
	}
	result["Content"] = htmlTemplate.HTML(content)

	return result
}

func (b *builder) renderPage(ctx context.Context, row client.SefUrlItem) ([]byte, error) {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	layoutName := "page"
	if row.Template != nil && *row.Template != "" {
		layoutName = *row.Template
	}

	layoutFile := []string{
		serviceBasePath + "/schemes/templates/layouts/common.html",
		serviceBasePath + "/schemes/templates/layouts/" + layoutName + ".html",
	}

	templateCommonFiles, _ := filepath.Glob(
		serviceBasePath + "/schemes/templates/blocks/common/*.html",
	)
	blockFiles, _ := filepath.Glob(
		serviceBasePath + "/schemes/templates/blocks/" + layoutName + "/*.html",
	)
	templateFiles := append(layoutFile, templateCommonFiles...)
	templateFiles = append(templateFiles, blockFiles...)

	tmpl, err := template.New("common").ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	assetManager := asset_manager.New()
	b.widgetManager.SetAssets(assetManager)

	err = tmpl.ExecuteTemplate(buf, "common", b.getPageData(ctx, row.IdObject))
	if err != nil {
		return nil, err
	}

	content := buf.String()

	b.widgetManager.Render(ctx, "header", row)

	reg, _ := regexp.Compile(`{#outertemplate%([A-Za-z_0-9]+)#}`)
	for _, match := range reg.FindAllStringSubmatch(content, -1) {
		match_1 := match[1]
		if len(match) > 2 {
			match_1 = match[2]
		}

		res, _ := b.widgetManager.Render(ctx, match_1, row)
		if res != nil {
			content = strings.Replace(content, match[0], *res, -1)
		}
	}

	content = strings.Replace(content, "{#systemtemplate%bottomjs#}", assetManager.GetJsTemplate(), -1)
	content = strings.Replace(content, "{#systemtemplate%topcss#}", assetManager.GetCssTemplate(), -1)
	content = strings.Replace(content, "{#systemtemplate%toppreload#}", assetManager.GetPreloadTemplate(), -1)

	return []byte(content), nil
}

func (b *builder) getCategoryData(ctx context.Context, idObject int64) map[string]interface{} {
	result := map[string]interface{}{}

	categoryRow, err := b.hub.Category().Read(ctx, idObject)
	if err != nil {
		return result
	}

	result["Year"] = time.Now().Year()

	result["Name"] = categoryRow.Name
	content := ""
	if categoryRow.Content != nil {
		content = *categoryRow.Content
	}
	result["Content"] = htmlTemplate.HTML(content)

	return result
}

func (b *builder) renderCategory(ctx context.Context, row client.SefUrlItem) ([]byte, error) {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	layoutName := "category"
	if row.Template != nil && *row.Template != "" {
		layoutName = *row.Template
	}

	layoutFile := []string{
		serviceBasePath + "/schemes/templates/layouts/common.html",
		serviceBasePath + "/schemes/templates/layouts/" + layoutName + ".html",
	}

	templateCommonFiles, _ := filepath.Glob(
		serviceBasePath + "/schemes/templates/blocks/category/*.html",
	)
	blockFiles, _ := filepath.Glob(
		serviceBasePath + "/schemes/templates/blocks/" + layoutName + "/*.html",
	)
	templateFiles := append(layoutFile, templateCommonFiles...)
	templateFiles = append(templateFiles, blockFiles...)

	tmpl, err := template.New("common").ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	assetManager := asset_manager.New()
	b.widgetManager.SetAssets(assetManager)

	err = tmpl.ExecuteTemplate(buf, "common", b.getCategoryData(ctx, row.IdObject))
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

		res, _ := b.widgetManager.Render(ctx, match_1, row)
		if res != nil {
			content = strings.Replace(content, match[0], *res, -1)
		}
	}

	return []byte(content), nil
}

func (b *builder) getProductData(ctx context.Context, idObject int64) map[string]interface{} {
	result := map[string]interface{}{}

	productRow, err := b.hub.Product().Read(ctx, product.ProductReadInput{ID: idObject})
	if err != nil {
		return result
	}

	result["Year"] = time.Now().Year()

	result["Name"] = productRow.Name
	content := ""
	if productRow.Content != nil {
		content = *productRow.Content
	}
	result["Content"] = htmlTemplate.HTML(content)

	return result
}

func (b *builder) renderProduct(ctx context.Context, row client.SefUrlItem) ([]byte, error) {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	layoutName := "product"
	if row.Template != nil && *row.Template != "" {
		layoutName = *row.Template
	}

	layoutFile := []string{
		serviceBasePath + "/schemes/templates/layouts/common.html",
		serviceBasePath + "/schemes/templates/layouts/" + layoutName + ".html",
	}

	templateCommonFiles, _ := filepath.Glob(
		serviceBasePath + "/schemes/templates/blocks/product/*.html",
	)
	blockFiles, _ := filepath.Glob(
		serviceBasePath + "/schemes/templates/blocks/" + layoutName + "/*.html",
	)
	templateFiles := append(layoutFile, templateCommonFiles...)
	templateFiles = append(templateFiles, blockFiles...)

	tmpl, err := template.New("common").ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	assetManager := asset_manager.New()
	b.widgetManager.SetAssets(assetManager)

	err = tmpl.ExecuteTemplate(buf, "common", b.getProductData(ctx, row.IdObject))
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

		res, _ := b.widgetManager.Render(ctx, match_1, row)
		if res != nil {
			content = strings.Replace(content, match[0], *res, -1)
		}
	}

	return []byte(content), nil
}
