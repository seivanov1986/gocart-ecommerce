package image

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
	parentIDStr := r.Header.Get("X-Parent-ID")
	parentID, _ := strconv.ParseInt(parentIDStr, 10, 64)
	uid := r.Header.Get("X-Uid")
	totalStr := r.Header.Get("X-Total")
	total, _ := strconv.ParseInt(totalStr, 10, 64)
	offsetStr := r.Header.Get("X-Offset")
	offset, _ := strconv.ParseInt(offsetStr, 10, 64)
	b64Name := r.Header.Get("X-Name")
	sDec, _ := b64.StdEncoding.DecodeString(b64Name)
	name := string(sDec)

	fmt.Println(uid, total, offset, name)

	filePath := "/tmp/" + uid

	if offset == 0 {
		// TODO start go rutine monitor for delete phantome
	}

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) && offset > 0 {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	p, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}
	defer p.Close()

	_, err = io.Copy(p, r.Body)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	uploadOut := UploadOut{
		Success: true,
	}

	if offset >= total {
		var done = true
		uploadOut.Done = &done

		err := a.service.Create(r.Context(), image.ImageCreateIn{
			Name:      name,
			ParentID:  parentID,
			FType:     1,
			CreatedAT: time.Now().Unix(),
		})
		if err != nil {
			helpers.HttpResponse(w, http.StatusInternalServerError)
			return
		}
	}

	helpers.HttpResponse(w, http.StatusOK, uploadOut)
}
