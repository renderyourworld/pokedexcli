package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	cacheData, found := c.cache.Get(url)
	if found {
		fmt.Println("Cache hit for", url)

		locationResp := Location{}
		err := json.Unmarshal(cacheData, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	fmt.Println("Cache miss for", url)

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	// Add to cache
	c.cache.Add(url, data)

	return locationResp, nil
}
