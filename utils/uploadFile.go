package utils

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
	"github.com/renaldyhidayatt/inventorygoent/dto/response"
)

type locationUpload struct {
	location string
	image    string
}

func NewLocationUpload(location string, image string) *locationUpload {
	return &locationUpload{location: location, image: image}
}

func (l *locationUpload) FileUpload(w http.ResponseWriter, r *http.Request) string {
	dir, err := os.Getwd()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}
	location := fmt.Sprintf("images/%s", l.location)

	folderLocation := filepath.Join(dir, location)

	if _, err := os.Stat(folderLocation); os.IsNotExist(err) {
		os.MkdirAll(folderLocation, 0700)
	}

	file, handler, err := r.FormFile("image")

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

	}
	defer file.Close()

	contentType := handler.Header.Get("Content-Type")

	if !l.CheckType(contentType) {
		response.ResponseError(w, http.StatusInternalServerError, errors.New("format file tidak didukung"))

	}

	var img image.Image

	switch contentType {
	case "image/jpg":
		img, err = jpeg.Decode(file)
	case "image/jpeg":
		img, err = jpeg.Decode(file)
	case "image/png":
		img, err = png.Decode(file)
	case "image/gif":
		img, err = gif.Decode(file)
	}

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

	}
	size := resize.Resize(300, 300, img, resize.Lanczos3)

	// Retrieve file information
	randomString := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)

	for i := range b {
		b[i] = randomString[rand.Intn(len(randomString))]
	}

	filename := fmt.Sprintf("%s%s", string(b), filepath.Ext(handler.Filename))

	fileLocation := ""

	if l.image != "" {
		fileLocation = filepath.Join(dir, l.image)
	}

	exits, err := os.Stat(fileLocation)

	if exits != nil {
		e := os.Remove(fileLocation)
		if e != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)

		}
	}

	out, err := os.Create(folderLocation + `/` + filename)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

	}
	defer out.Close()

	switch contentType {
	case "image/jpg":
		err = jpeg.Encode(out, size, nil)
	case "image/jpeg":
		err = jpeg.Encode(out, size, nil)
	case "image/png":
		err = png.Encode(out, size)
	case "image/gif":
		err = gif.Encode(out, size, nil)
	}

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	if _, err := io.Copy(out, file); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

	}

	return fmt.Sprintf("/images/posts/%s", filename)
}

func (l *locationUpload) CheckType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/jpg" || contentType == "image/gif"
}
