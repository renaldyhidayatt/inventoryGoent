package upload

import (
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
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
)

type locationUpload struct {
	location string
	image    string
}

func NewLocationUpload(location string, image string) *locationUpload {
	return &locationUpload{location: location, image: image}
}

func (l *locationUpload) FileUpload(w http.ResponseWriter, r *http.Request) string {
	// Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "internal server error")
		return ""
	}

	// Prepare the folder location
	location := fmt.Sprintf("images/%s", l.location)
	folderLocation := filepath.Join(dir, location)

	// Create the folder if it doesn't exist
	if _, err := os.Stat(folderLocation); os.IsNotExist(err) {
		os.MkdirAll(folderLocation, 0700)
	}

	// Retrieve the uploaded file
	file, handler, err := r.FormFile("image")
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "internal server error")
		return ""
	}
	defer file.Close()

	// Check if the content type is supported
	contentType := handler.Header.Get("Content-Type")
	if !l.CheckType(contentType) {
		response.ResponseError(w, http.StatusInternalServerError, "content type not supported")
		return ""
	}

	// Decode the image based on content type
	var img image.Image
	switch contentType {
	case "image/jpg", "image/jpeg", "image/png", "image/gif":
		img, _, err = image.Decode(file)
	default:
		response.ResponseError(w, http.StatusInternalServerError, "unsupported image format")
		return ""
	}

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "internal server error")
		return ""
	}

	// Resize the image
	size := resize.Resize(300, 300, img, resize.Lanczos3)

	// Generate a random filename
	randomString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]rune, 10)
	for i := range b {
		b[i] = rune(randomString[rand.Intn(len(randomString))])
	}
	filename := fmt.Sprintf("%s%s", string(b), filepath.Ext(handler.Filename))

	// Prepare the file location
	fileLocation := filepath.Join(dir, l.image)

	// Remove the existing file if it exists
	if _, err := os.Stat(fileLocation); err == nil {
		if err := os.Remove(fileLocation); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, "internal server error")
			return ""
		}
	}

	// Create the new file
	out, err := os.Create(filepath.Join(folderLocation, filename))
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "internal server error")
		return ""
	}
	defer out.Close()

	// Encode and save the resized image
	switch contentType {
	case "image/jpg", "image/jpeg":
		err = jpeg.Encode(out, size, nil)
	case "image/png":
		err = png.Encode(out, size)
	case "image/gif":
		err = gif.Encode(out, size, nil)
	default:
		response.ResponseError(w, http.StatusInternalServerError, "unsupported image format")
		return ""
	}

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "internal server error")
		return ""
	}

	if _, err := io.Copy(out, file); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "internal server error")
		return ""
	}

	return fmt.Sprintf("/images/posts/%s", filename)
}

func (l *locationUpload) CheckType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/jpg" || contentType == "image/gif"
}
