package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// getLinkPage gets a particular page for a title
func getLinkPage(title string, targetAPI *Config, plcontinue *string) (*LinkResponse, error) {
	u, _ := url.Parse(targetAPI.APIRoot)
	u.Scheme = targetAPI.Protocol

	q := u.Query()
	q.Set("action", "query")
	q.Set("titles", title)
	q.Set("prop", "links")
	q.Set("pllimit", "max")
	q.Set("format", "json")

	if plcontinue != nil {
		q.Set("plcontinue", *plcontinue)
	}

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

func fetchLoop(title string, targetAPI *Config, plcontinue *string) ([]string, error) {
	page, err := getLinkPage(title, targetAPI, plcontinue)
	if err != nil {
		return nil, err
	}

	if page.BatchComplete != nil {
		return []string{}, nil
	}

	var result []string
	for _, v := range page.Query.PageLinks {
		titles := make([]string, len(v.Links))
		for i, l := range v.Links {
			titles[i] = l.Title
		}

		result = append(result, titles...)
	}

	nextResponse, err := fetchLoop(title, targetAPI, &page.Continue.Plcontinue)
	if err != nil {
		return nil, err
	}

	return append(result, nextResponse...), nil
}

// GetAllLinks returns a list of all the links in a page
func GetAllLinks(title string, targetAPI *Config) ([]string, error) {
	return fetchLoop(title, targetAPI, nil)
}
