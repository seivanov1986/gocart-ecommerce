package image_to_product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository/image_to_product"
	image_to_product2 "github.com/seivanov1986/gocart/internal/service/image_to_product"
)

type ImageToProductListRpcIn struct {
	Page      int64 `json:"page"`
	ProductID int64 `json:"id_product"`
}

type ImageToProductListRpcOut struct {
	List  []image_to_product.ImageToProductListRow
	Total int64
}

type ImageToProductListError struct {
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	PageListInput, err := validateImageToProductList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *PageListInput)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, ImageToProductListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateImageToProductList(bodyBytes []byte) (*image_to_product2.ImageToProductListIn, error) {
	listInt := image_to_product2.ImageToProductListIn{}
	userListRpcIn := ImageToProductListRpcIn{}

	err := json.Unmarshal(bodyBytes, &userListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ProductID = userListRpcIn.ProductID
	listInt.Page = userListRpcIn.Page

	return &listInt, nil
}
