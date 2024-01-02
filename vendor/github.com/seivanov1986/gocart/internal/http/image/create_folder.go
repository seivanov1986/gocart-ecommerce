package image

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/service/image"
)

type ImageCreateFolderRpcIn struct {
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
}

func (a *handle) CreateFolder(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	list, err := validateImageCreateFolder(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = a.service.CreateFolder(r.Context(), *list)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validateImageCreateFolder(bodyBytes []byte) (*image.ImageCreateFolderIn, error) {
	listInt := image.ImageCreateFolderIn{}
	categoryListRpcIn := ImageCreateFolderRpcIn{}

	err := json.Unmarshal(bodyBytes, &categoryListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Name = categoryListRpcIn.Name
	listInt.ParentID = categoryListRpcIn.ParentID

	return &listInt, nil
}
