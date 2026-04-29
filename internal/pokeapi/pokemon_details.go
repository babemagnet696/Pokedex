package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonDetails(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	var body []byte

	if data, ok := c.cache.Get(url); ok {
		body = data
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, body)
	}

	pokemonResp := Pokemon{}
	err := json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Invalid pokemon name")
	}

	return pokemonResp, nil
}