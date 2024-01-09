package widget

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/external/observer"
	"github.com/seivanov1986/sql_client"
)

type postListWidget struct {
	assetManager client.AssetManager
	sqlClient    sql_client.DataBase
}

func NewPostList(sqlClient sql_client.DataBase) *postListWidget {
	return &postListWidget{
		sqlClient: sqlClient,
	}
}

func (l *postListWidget) Execute(ctx context.Context, url client.SefUrlItem) (*string, error) {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	layoutFile := []string{
		serviceBasePath + "/schemes/templates/widgets/post_list.html",
	}

	tmpl, err := template.New("common").ParseFiles(layoutFile...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	buf := &bytes.Buffer{}

	var list []struct {
		Id           string  `db:"id" json:"id"`
		Name         string  `db:"name" json:"name"`
		ShortContent *string `db:"short_content" json:"short_content"`
		Content      *string `db:"content" json:"content"`
		Url          string  `db:"url" json:"url"`
	}

	err = l.sqlClient.SelectContext(
		ctx,
		&list,
		`SELECT page.id, page.short_content, page.name, page.content, sefurl.url FROM page 
    			LEFT JOIN sefurl ON page.id = sefurl.id_object AND sefurl.type = 1 
                WHERE page.type = 1`,
		url.IdObject,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = tmpl.ExecuteTemplate(buf, "post_list", map[string]interface{}{
		"Items": list,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := buf.String()
	return &result, nil
}

func (l *postListWidget) SetAssets(assetManager client.AssetManager) {
	l.assetManager = assetManager
}
