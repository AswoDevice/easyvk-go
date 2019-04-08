package easyvk

import (
	"encoding/json"
)

// Methods for working with groups.
// https://vk.com/dev/groups
type Groups struct {
	vk *VK
}

// GroupsGetResponse describes
// https://vk.com/dev/groups.getById
type GroupsGetResponse struct {
	Count int
	Items []uint
}

// GroupsGetParams provides structure for
// parameters for get method.
// https://vk.com/dev/groups.getById
type GroupsGetParams struct {
	GroupIDs string
	GroupID  string
	Fields   string
}

// Returns information about communities
// by their IDs.
// https://vk.com/dev/groups.getById
func (g *Groups) GetByID(par GroupsGetParams) ([]GroupObject, error) {
	params := make(map[string]string)
	if par.GroupIDs != "" {
		params["group_ids"] = par.GroupIDs
	}
	if par.GroupID != "" {
		params["group_id"] = par.GroupID
	}
	if par.Fields != "" {
		params["fields"] = par.Fields
	}

	resp, err := g.vk.Request("groups.getById", params)
	var array = make([]GroupObject, 0)
	if err != nil {
		return array, err
	}

	err = json.Unmarshal(resp, &array)
	if err != nil {
		return array, err
	}
	return array, nil
}
