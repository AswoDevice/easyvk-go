package easyvk

import (
	"encoding/json"
	"fmt"
)

// Methods for working with friends.
// https://vk.com/dev/storage
type Storage struct {
	vk *VK
}

// Returns a value of variable with the
// name set by key parameter.
// https://vk.com/dev/storage.get
func (f *Storage) Get(userId uint32, key string) (string, error) {
	params := make(map[string]string)
	params["user_id"] = fmt.Sprint(userId)
	params["key"] = key

	resp, err := f.vk.Request("storage.get", params)
	if err != nil {
		return "", err
	}

	var value string
	err = json.Unmarshal(resp, &value)
	if err != nil {
		return "", err
	}
	return value, nil
}


// Returns a value of variable with the
// name set by key parameter.
// https://vk.com/dev/storage.set
func (f *Storage) Set(userId uint32, key, value string) error {
	params := make(map[string]string)
	params["user_id"] = fmt.Sprint(userId)
	params["key"] = key
	params["value"] = value

	_, err := f.vk.Request("storage.set", params)
	return err
}
