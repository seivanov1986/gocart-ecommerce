package attribute_to_product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository/attribute_to_product"
	attribute_to_product2 "github.com/seivanov1986/gocart/internal/service/attribute_to_product"
)

type AttributeToProductListRpcIn struct {
	ProductID int64 `json:"id_product"`
	Page      int64 `json:"page"`
}

type AttributeToProductListOut struct {
	List  []attribute_to_product.AttributeToProductListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	listInput, err := validateAttributeToProductList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *listInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, AttributeToProductListOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateAttributeToProductList(bodyBytes []byte) (*attribute_to_product2.AttributeToProductListIn, error) {
	listInt := attribute_to_product2.AttributeToProductListIn{}
	userCreateRpcIn := AttributeToProductListRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = userCreateRpcIn.Page
	listInt.ProductID = userCreateRpcIn.ProductID

	return &listInt, nil
}
