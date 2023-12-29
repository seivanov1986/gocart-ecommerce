package product_to_category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	product_to_category2 "github.com/seivanov1986/gocart/internal/repository/product_to_category"
	"github.com/seivanov1986/gocart/internal/service/product_to_category"
)

type ProductToCategoryListRpcIn struct {
	ProductID int64 `json:"id_product"`
	Page      int64 `json:"page"`
}

type ProductToCategoryListOut struct {
	List  []product_to_category2.ProductToCategoryListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	listInput, err := validateProductToCategoryList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *listInput)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, ProductToCategoryListOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateProductToCategoryList(bodyBytes []byte) (*product_to_category.ProductToCategoryListInput, error) {
	listInt := product_to_category.ProductToCategoryListInput{}
	userListRpcIn := ProductToCategoryListRpcIn{}

	err := json.Unmarshal(bodyBytes, &userListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = userListRpcIn.Page
	listInt.ProductID = userListRpcIn.ProductID

	return &listInt, nil
}
