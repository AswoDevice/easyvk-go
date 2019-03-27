package easyvk

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// An Upload describes a set of methods
// that helps with update to VK servers.
type Upload struct{}

// A UploadResponse describes an info
// about uploaded photo.
type UploadPhotoResponse struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// PhotoWall upload file (on filePath) to given url.
// Return info about uploaded photo.
func (u *Upload) PhotoWall(url, filePath string) (UploadPhotoResponse, error) {
	return u.Photo(url, filePath)
}

// Photo upload file (on filePath) to given url.
// Return info about uploaded photo.
func (u *Upload) Photo(url, filePath string) (UploadPhotoResponse, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("photo", filePath)
	if err != nil {
		return UploadPhotoResponse{}, err
	}

	fh, err := os.Open(filePath)
	if err != nil {
		return UploadPhotoResponse{}, err
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return UploadPhotoResponse{}, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return UploadPhotoResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UploadPhotoResponse{}, err
	}

	var uploaded UploadPhotoResponse
	err = json.Unmarshal(body, &uploaded)
	if err != nil {
		return UploadPhotoResponse{}, err
	}

	return uploaded, nil
}
