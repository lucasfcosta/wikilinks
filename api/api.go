package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetAllLinks gets all links for a Wikipedia page
func GetAllLinks(title string, targetAPI *Config) (*LinkResponse, error) {
	u, _ := url.Parse(targetAPI.APIRoot)
	u.Scheme = targetAPI.Protocol

	q := u.Query()
	q.Set("action", "query")
	q.Set("titles", title)
	q.Set("prop", "links")
	q.Set("pllimit", "max")
	q.Set("format", "json")

	u.RawQuery = q.Encode()

	res, reqErr := http.Get(u.String())

	if reqErr != nil {
		fmt.Println("Request failed!")
		return nil, reqErr
	}

	defer res.Body.Close()

	body, readBodyErr := ioutil.ReadAll(res.Body)
	if readBodyErr != nil {
		fmt.Println("Can't read response body!")
		return nil, readBodyErr
	}

	data := LinkResponse{}
	jsonParseErr := json.Unmarshal(body, &data)
	if jsonParseErr != nil {
		fmt.Println("Invalid json!")
		return nil, readBodyErr
	}

	return &data, nil
}
