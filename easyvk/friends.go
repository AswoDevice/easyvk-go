package easyvk

import (
	"encoding/json"
	"fmt"
)

// Methods for working with friends.
// https://vk.com/dev/friends
type Friends struct {
	vk *VK
}

// PhotosGetResponse describes
// https://vk.com/dev/photos.get
type FriendsGetResponse struct {
	Count int
	Items []uint
}

// UsersGetParams provides structure for
// parameters for get method.
// Returns only ids.
// https://vk.com/dev/friends.get
type FriendsGetParams struct {
	UserID   uint
	Order    string
	Count    uint
	Offset   uint
}

// Returns a list of user IDs or detailed
// information about a user's friends.
// https://vk.com/dev/friends.get
func (f *Friends) Get(par FriendsGetParams) (FriendsGetResponse, error) {
	params := make(map[string]string)
	params["user_id"] = fmt.Sprint(par.UserID)
	if par.Order != "" {
		params["order"] = par.Order
	}
	params["count"] = fmt.Sprint(par.Count)
	params["offset"] = fmt.Sprint(par.Offset)

	resp, err := f.vk.Request("friends.get", params)
	if err != nil {
		return FriendsGetResponse{}, err
	}

	var friends FriendsGetResponse
	err = json.Unmarshal(resp, &friends)
	if err != nil {
		return FriendsGetResponse{}, err
	}
	return friends, nil
}
