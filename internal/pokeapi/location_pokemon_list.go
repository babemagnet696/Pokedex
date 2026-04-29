package pokeapi

import  (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

// List all Pokemon in location -
func (c *Client) ListPokemon(location string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + location
	var body []byte

	if data, ok := c.cache.Get(url); ok {
		body = data
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocationArea{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocationArea{}, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespLocationArea{}, err
		}
		c.cache.Add(url, body)
	}

	locationResp := RespLocationArea{}
	err := json.Unmarshal(body, &locationResp)
	if err != nil {
		return RespLocationArea{}, fmt.Errorf("Invalid location entered")
	}

	return locationResp, nil
}