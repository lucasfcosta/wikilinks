package api

// LinkResponse represents the response for a link query on the Wikipedia API.
type LinkResponse struct {
	BatchComplete *string `json:"batchcomplete"`
	Continue      struct {
		Plcontinue string `json:"plcontinue"`
		Continue   string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Normalized []struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"normalized"`
		PageLinks map[string]PageResult `json:"pages"`
	} `json:"query"`
	Limits struct {
		Links int `json:"links"`
	} `json:"limits"`
}

// PageResult represents the result for each page queried
type PageResult struct {
	Pageid int    `json:"pageid"`
	Ns     int    `json:"ns"`
	Title  string `json:"title"`
	Links  []struct {
		Ns    int    `json:"ns"`
		Title string `json:"title"`
	} `json:"links"`
}
