package category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/service/category"
	categoryRepository "github.com/seivanov1986/gocart/internal/repository/category"
)

type CategoryListRpcIn struct {
	Page     int64
	ParentID int64
}

type CategoryListRpcOut struct {
	List  []categoryRepository.CategoryListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	list, err := validateCategoryList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *list)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, CategoryListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateCategoryList(bodyBytes []byte) (*category.CategoryListIn, error) {
	listInt := category.CategoryListIn{}
	categoryListRpcIn := CategoryListRpcIn{}

	err := json.Unmarshal(bodyBytes, &categoryListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = categoryListRpcIn.Page
	listInt.ParentID = categoryListRpcIn.ParentID

	return &listInt, nil
}
