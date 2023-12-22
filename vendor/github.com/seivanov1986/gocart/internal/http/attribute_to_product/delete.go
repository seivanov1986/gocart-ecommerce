package attribute_to_product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

type AttributeToProductDeleteRpcIn struct {
}

type AttributeToProductDeleteError struct {
}

func (u *handle) Delete(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	_, err = validateAttributeToProductDelete(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validateAttributeToProductDelete(bodyBytes []byte) ([]int64, error) {
	listInt := []int64{}
	userCreateRpcIn := AttributeToProductDeleteRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return listInt, nil
}
