package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache first
	cacheData, found := c.cache.Get(url)
	if found {
		fmt.Println("Cache hit for", url)

		locationAreasResp := LocationAreaResp{}
		err := json.Unmarshal(cacheData, &locationAreasResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreasResp, nil
	}

	fmt.Println("Cache miss for", url)

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreasResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// Add to cache
	c.cache.Add(url, data)

	return locationAreasResp, nil
}
