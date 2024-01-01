package upload_handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	b64 "encoding/base64"

	"github.com/seivanov1986/gocart/helpers"
)

type handler struct {
}

type Handler interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

func New() *handler {
	return &handler{}
}

type UploadOut struct {
	Success bool  `json:"success"`
	Done    *bool `json:"done"`
}

func (a *handler) Upload(w http.ResponseWriter, r *http.Request) {
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

		/*
			TODO:: service -> save to DB
			service -> make thumbs
		*/
	}

	helpers.HttpResponse(w, http.StatusOK, uploadOut)
}
