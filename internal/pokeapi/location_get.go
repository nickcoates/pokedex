package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(name string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + name

	if data, ok := c.cache.Get(url); ok {
		locationsResp := RespLocationArea{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespLocationArea{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return RespLocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	locationsResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
