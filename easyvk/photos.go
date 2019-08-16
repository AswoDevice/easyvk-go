package easyvk

import (
	"encoding/json"
	"fmt"
)

// A Photos describes a set of methods
// to work with photos.
// https://vk.com/dev/photos
type Photos struct {
	vk *VK
}

// PhotosGetWallUploadServerResponse describes the server address
// for photo upload onto a user's wall.
// https://vk.com/dev/photos.getWallUploadServer
type PhotosGetWallUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int `json:"album_id"`
	UserID    int `json:"user_id"`
}

// GetWallUploadServer returns the server address for photo upload onto a user's wall.
// https://vk.com/dev/photos.getWallUploadServer
func (p *Photos) GetWallUploadServer(groupID uint) (PhotosGetWallUploadServerResponse, error) {
	params := map[string]string{"group_id": fmt.Sprint(groupID) }
	resp, err := p.vk.Request("photos.getWallUploadServer", params)
	if err != nil {
		return PhotosGetWallUploadServerResponse{}, err
	}
	var server PhotosGetWallUploadServerResponse
	err = json.Unmarshal(resp, &server)
	if err != nil {
		return PhotosGetWallUploadServerResponse{}, err
	}
	return server, nil
}

// PhotosGetMessagesUploadServerResponse describes the server address
// for photo upload onto a message.
// https://vk.com/dev/photos.getMessagesUploadServer
type PhotosGetMessagesUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int `json:"album_id"`
	UserID    int `json:"user_id"`
}

// GetMessagesUploadServer returns the server address for photo upload onto a message.
// https://vk.com/dev/photos.getMessagesUploadServer
func (p *Photos) GetMessagesUploadServer(peerId uint) (PhotosGetMessagesUploadServerResponse, error) {
	params := map[string]string{"peer_id": fmt.Sprint(peerId) }
	resp, err := p.vk.Request("photos.getMessagesUploadServer", params)
	if err != nil {
		return PhotosGetMessagesUploadServerResponse{}, err
	}
	var server PhotosGetMessagesUploadServerResponse
	err = json.Unmarshal(resp, &server)
	if err != nil {
		return PhotosGetMessagesUploadServerResponse{}, err
	}
	return server, nil
}

// SaveMessagesPhoto saves a photo.
// For upload look at file upload.go.
// https://vk.com/dev/photos.saveMessagesPhoto
func (p *Photos) SaveMessagesPhoto(photo string, server int, hash string) ([]PhotoObject, error) {
	params := map[string]string{
		"photo":  photo,
		"hash":   hash,
		"server": fmt.Sprint(server),
	}
	resp, err := p.vk.Request("photos.saveMessagesPhoto", params)
	if err != nil {
		return nil, err
	}

	var info []PhotoObject
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// PhotosSaveWallPhotoParams provides structure for
// parameters for saveWallPhoto method.
// https://vk.com/dev/photos.saveWallPhoto
type PhotosSaveWallPhotoParams struct {
	UserID  uint
	GroupID uint
	Photo   string
	Hash    string
	Caption string
	Server  int
	Lat     float64
	Long    float64
}

// SaveWallPhoto saves a photo to a user's or community's wall after being uploaded.
// For upload look at file upload.go.
// https://vk.com/dev/photos.saveWallPhoto
func (p *Photos) SaveWallPhoto(par PhotosSaveWallPhotoParams) ([]PhotoObject, error) {
	params := map[string]string{
		"user_id":   fmt.Sprint(par.UserID),
		"group_id":  fmt.Sprint(par.GroupID),
		"photo":     par.Photo,
		"hash":      par.Hash,
		"caption":   par.Caption,
		"server":    fmt.Sprint(par.Server),
		"latitude":  fmt.Sprint(par.Lat),
		"longitude": fmt.Sprint(par.Long),
	}
	resp, err := p.vk.Request("photos.saveWallPhoto", params)
	if err != nil {
		return nil, err
	}

	var info []PhotoObject
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// PhotosGetResponse describes
// https://vk.com/dev/photos.get
type PhotosGetResponse struct {
	Count int
	Items []PhotoObject
}

// PhotosGetParams provides structure for
// parameters for Get method.
// https://vk.com/dev/photos.get
type PhotosGetParams struct {
	OwnerID  uint
	AlbumID  string
	PhotoIDs string
	Rev      bool
	Offset   int
	Count    int
}

// Returns a list of a user's or community's photos.
// https://vk.com/dev/photos.get
func (p *Photos) Get(par PhotosGetParams) (PhotosGetResponse, error) {
	params := make(map[string]string)

	if par.OwnerID != 0 {
		params["owner_id"] = fmt.Sprint(par.OwnerID)
	}

	if par.AlbumID != "" {
		params["album_id"] = par.AlbumID
	}

	params["photo_ids"] = par.PhotoIDs
	params["rev"] = boolConverter(par.Rev)

	if par.Offset != 0 {
		params["offset"] = fmt.Sprint(par.Offset)
	}

	if par.Count != 0 {
		params["count"] = fmt.Sprint(par.Count)
	}

	resp, err := p.vk.Request("photos.get", params)
	if err != nil {
		return PhotosGetResponse{}, err
	}

	var photos PhotosGetResponse
	err = json.Unmarshal(resp, &photos)
	if err != nil {
		return PhotosGetResponse{}, err
	}
	return photos, nil
}
