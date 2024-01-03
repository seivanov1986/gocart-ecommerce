package image

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	b64 "encoding/base64"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/service/image"
)

type UploadOut struct {
	Success bool  `json:"success"`
	Done    *bool `json:"done"`
}

func (a *handle) Upload(w http.ResponseWriter, r *http.Request) {
	out, err := validateImageUpload(r)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = a.service.Create(r.Context(), *out)
	if err != nil {
		if err == image.ErrNotDone {
			var done = true
			helpers.HttpResponse(w, http.StatusOK, UploadOut{
				Success: true,
				Done:    &done,
			})
			return
		}

		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, UploadOut{
		Success: true,
	})
}

func validateImageUpload(r *http.Request) (*image.ImageCreateIn, error) {
	out := image.ImageCreateIn{}

	createdAt := time.Now()

	parentIDStr := r.Header.Get("X-Parent-ID")
	parentID, err := strconv.ParseInt(parentIDStr, 10, 64)
	if err != nil {
		return nil, err
	}

	uid := r.Header.Get("X-Uid")
	totalStr := r.Header.Get("X-Total")
	total, err := strconv.ParseInt(totalStr, 10, 64)
	if err != nil {
		return nil, err
	}

	offsetStr := r.Header.Get("X-Offset")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return nil, err
	}

	b64Name := r.Header.Get("X-Name")
	sDec, err := b64.StdEncoding.DecodeString(b64Name)
	if err != nil {
		return nil, err
	}

	name := string(sDec)

	t := createdAt.Format(time.DateOnly)
	d := strings.Split(t, "-")

	path := "/tmp/" + d[0] + "/" + d[1] + "/" + d[2] + "/"

	out.Name = name
	out.ParentID = parentID
	out.FType = 0
	out.CreatedAT = createdAt.Unix()
	out.UID = uid
	out.OriginPath = path
	out.Total = total
	out.Offset = offset
	out.Body = r.Body

	return &out, nil
}
