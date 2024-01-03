package image_to_product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	image_to_product2 "github.com/seivanov1986/gocart/internal/service/image_to_product"
)

type ImageToProductCreateRpcIn struct {
	ProductID int64 `json:"id_product"`
	ImageID   int64 `json:"id_image"`
}

type ImageToProductCreateRpcOut struct {
}

type ImageToProductCreateError struct {
	Error string
}

func (u *handle) Create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validateImageToProductCreate(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = u.service.Create(r.Context(), *CreateListInput)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusInternalServerError, ImageToProductCreateError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK, ImageToProductCreateRpcOut{})
}

func validateImageToProductCreate(bodyBytes []byte) (*image_to_product2.ImageToProductCreateInput, error) {
	listInt := image_to_product2.ImageToProductCreateInput{}
	userCreateRpcIn := ImageToProductCreateRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ProductID = userCreateRpcIn.ProductID
	listInt.ImageID = userCreateRpcIn.ImageID

	return &listInt, nil
}
