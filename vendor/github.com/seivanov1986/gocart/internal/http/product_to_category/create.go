package product_to_category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/service/product_to_category"
)

type ProductToCategoryCreateRpcIn struct {
	ProductID    int64  `json:"id_product"`
	CategoryID   int64  `json:"id_category"`
	MainCategory *int64 `json:"main_category"`
}

type ProductToCategoryCreateError struct {
}

func (u *handle) Create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	listInput, err := validateProductToCategoryCreate(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = u.service.Create(r.Context(), *listInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, ProductToCategoryCreateError{})
		return
	}

	u.cacheObject.AddEvent()

	helpers.HttpResponse(w, http.StatusOK, ProductToCategoryCreateError{})
}

func validateProductToCategoryCreate(bodyBytes []byte) (*product_to_category.ProductToCategoryCreateInput, error) {
	listInt := product_to_category.ProductToCategoryCreateInput{}
	userCreateRpcIn := ProductToCategoryCreateRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ProductID = userCreateRpcIn.ProductID
	listInt.CategoryID = userCreateRpcIn.CategoryID
	listInt.MainCategory = userCreateRpcIn.MainCategory

	return &listInt, nil
}
