package easyvk

import (
	"carousel/service/logger"
	"encoding/json"
	"fmt"
)

// A Users describes a set of methods
// to work with user.
// https://vk.com/dev/users
type Users struct {
	vk *VK
}

// UsersGetParams provides structure for
// parameters for get method.
// https://vk.com/dev/users.get
type UsersGetParams struct {
	UserIDs  string
	Fields   string
	NameCase string
}

// Returns detailed information on users.
// https://vk.com/dev/users.get
func (u *Users) Get(par UsersGetParams) ([]UserObject, error) {
	params := make(map[string]string)
	if par.UserIDs != "" {
		params["user_ids"] = par.UserIDs
	}
	params["fields"] = par.Fields
	params["name_case"] = par.NameCase

	resp, err := u.vk.Request("users.get", params)
	var array = make([]UserObject, 0)
	if err != nil {
		return array, err
	}
logger.Debug(string(resp))
	err = json.Unmarshal(resp, &array)
	if err != nil {
		return array, err
	}
	return array, nil
}

// PhotosGetResponse describes
// https://vk.com/dev/photos.get
type PhotosGetFollowersResponse struct {
	Count int
	Items []UserObject
}

// UsersGetParams provides structure for
// parameters for get method.
// https://vk.com/dev/users.getFollowers
type UsersGetFollowersParams struct {
	UserID   uint
	Offset   int
	Count    int
	Fields   string
	NameCase string
}

// Returns a list of IDs of followers of the user in
// question, sorted by date added, most recent first.
// https://vk.com/dev/users.getFollowers
func (u *Users) GetFollowers(par UsersGetFollowersParams) (PhotosGetFollowersResponse, error) {
	params := make(map[string]string)
	if par.UserID != 0 {
		params["user_id"] = fmt.Sprint(par.UserID)
	}
	params["offset"] = fmt.Sprint(par.Offset)
	params["count"] = fmt.Sprint(par.Count)
	params["fields"] = par.Fields
	params["name_case"] = par.NameCase

	resp, err := u.vk.Request("users.getFollowers", params)
	if err != nil {
		return PhotosGetFollowersResponse{}, err
	}

	var getFollowers PhotosGetFollowersResponse
	err = json.Unmarshal(resp, &getFollowers)
	if err != nil {
		return PhotosGetFollowersResponse{}, err
	}
	return getFollowers, nil
}
