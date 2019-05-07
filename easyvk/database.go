package easyvk

import (
	"encoding/json"
	"fmt"
)

// These methods provide access to VK's database
// of educational institutions. Access to data
// is free and authorization is not required,
// but the number of requests from one IP address
// can be limited. If you need to execute many
// requests, we recommend that you call these
// methods from the client side using JSONP.
// https://vk.com/dev/database
type Database struct {
	vk *VK
}

// DatabaseGetCitiesResponse describes
// https://vk.com/dev/database.getCities
type DatabaseGetCitiesResponse struct {
	Count int
	Items []CityObject
}

// DatabaseGetCitiesParams provides structure for
// parameters for get method.
// https://vk.com/dev/database.getCities
type DatabaseGetCitiesParams struct {
	CountryId uint
	RegionId  uint
	Query     string
	NeedAll   bool
	Offset    uint
	Count     uint
}

// Returns a list of cities.
// https://vk.com/dev/database.getCities
func (d *Database) GetCities(par DatabaseGetCitiesParams) (DatabaseGetCitiesResponse, error) {
	params := make(map[string]string)
	params["country_id"] = fmt.Sprint(par.CountryId)

	if par.RegionId != 0 {
		params["region_id"] = fmt.Sprint(par.RegionId)
	}

	if par.Query != "" {
		params["q"] = par.Query
	}

	params["need_all"] = boolConverter(par.NeedAll)
	params["offset"] = fmt.Sprint(par.Offset)

	if par.Count != 0 {
		params["count"] = fmt.Sprint(par.Count)
	}

	resp, err := d.vk.Request("database.getCities", params)
	if err != nil {
		return DatabaseGetCitiesResponse{}, err
	}

	var cities DatabaseGetCitiesResponse
	err = json.Unmarshal(resp, &cities)
	if err != nil {
		return DatabaseGetCitiesResponse{}, err
	}
	return cities, nil
}
