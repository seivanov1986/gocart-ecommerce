package category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	categoryRepository "github.com/seivanov1986/gocart/internal/repository/category"
	"github.com/seivanov1986/gocart/internal/service/category"
)

type CategorySelectListRpcIn struct {
	Query    string
	ParentID int64
}

type CategorySelectListRpcOut struct {
	List []categoryRepository.CategorySelectListRow
}

func (u *handle) SelectList(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	list, err := validateCategorySelectList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.SelectList(r.Context(), *list)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, CategorySelectListRpcOut{
		List: listOut,
	})
}

func validateCategorySelectList(bodyBytes []byte) (*category.CategorySelectListIn, error) {
	listInt := category.CategorySelectListIn{}
	categoryListRpcIn := CategorySelectListRpcIn{}

	err := json.Unmarshal(bodyBytes, &categoryListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Query = categoryListRpcIn.Query
	listInt.ParentID = categoryListRpcIn.ParentID

	return &listInt, nil
}
