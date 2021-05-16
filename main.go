package main

import (
	"fmt"

	"github.com/lucasfcosta/wikilinks/api"
	"github.com/lucasfcosta/wikilinks/crawling"
)

func main() {
	// api.GetAllLinks("brazil", api.NewConfigFromEnv())
	path := crawling.FindPath("brazil", "Nafi Thiam", api.NewConfigFromEnv(), nil)

	fmt.Printf("%v", path)
}
