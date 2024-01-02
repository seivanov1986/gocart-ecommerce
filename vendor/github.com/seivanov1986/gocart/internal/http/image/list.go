package image

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository/image"
	image2 "github.com/seivanov1986/gocart/internal/service/image"
)

type ImageListRpcIn struct {
	Page     int64 `json:"page"`
	ParentID int64 `json:"parent_id"`
}

type ImageListRpcOut struct {
	List  []image.ImageListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	list, err := validateImageList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *list)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, ImageListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateImageList(bodyBytes []byte) (*image2.CategoryListIn, error) {
	listInt := image2.CategoryListIn{}
	categoryListRpcIn := ImageListRpcIn{}

	err := json.Unmarshal(bodyBytes, &categoryListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = categoryListRpcIn.Page
	listInt.ParentID = categoryListRpcIn.ParentID

	return &listInt, nil
}
