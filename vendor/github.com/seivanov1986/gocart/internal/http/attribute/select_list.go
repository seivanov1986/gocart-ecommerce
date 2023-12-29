package attribute

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	attribute2 "github.com/seivanov1986/gocart/internal/service/attribute"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
)

type AttributeSelectListRpcIn struct {
	Query string `json:"query"`
	Page  int64  `json:"page"`
}

type AttributeSelectListRpcOut struct {
	List []attribute.AttributeSelectListRow
}

func (u *handle) SelectList(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	AttributeSelectListInput, err := validateAttributeSelectList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.SelectList(r.Context(), *AttributeSelectListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, AttributeSelectListRpcOut{
		List: listOut,
	})
}

func validateAttributeSelectList(bodyBytes []byte) (*attribute2.AttributeSelectListIn, error) {
	listInt := attribute2.AttributeSelectListIn{}
	attributeSelectListRpcIn := AttributeSelectListRpcIn{}

	err := json.Unmarshal(bodyBytes, &attributeSelectListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Query = attributeSelectListRpcIn.Query
	listInt.Page = attributeSelectListRpcIn.Page

	return &listInt, nil
}
