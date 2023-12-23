package sefurl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	sefurl2 "github.com/seivanov1986/gocart/internal/repository/sefurl"
	"github.com/seivanov1986/gocart/internal/service/sefurl"
)

type SefUrlListRpcIn struct {
	Page int64 `json:"page"`
}

type SefUrlListRpcOut struct {
	List  []sefurl2.SefUrlListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	listInput, err := validateSefUrlList(bodyBytes)
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

	helpers.HttpResponse(w, http.StatusOK, SefUrlListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateSefUrlList(bodyBytes []byte) (*sefurl.SefUrlListIn, error) {
	listInt := sefurl.SefUrlListIn{}
	sefUrlListRpcIn := SefUrlListRpcIn{}

	err := json.Unmarshal(bodyBytes, &sefUrlListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = sefUrlListRpcIn.Page

	return &listInt, nil
}
