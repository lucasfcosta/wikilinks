package main

import (
	"fmt"

	"github.com/lucasfcosta/wikilinks/api"
	"github.com/lucasfcosta/wikilinks/crawling"
)

func main() {
	// api.GetAllLinks("brazil", api.NewConfigFromEnv())
	path, err := crawling.FindPath("brazil", "Nafi Thiam", api.NewConfigFromEnv())

	if err != nil {
		fmt.Printf("Couldn't find path")
	}

	fmt.Printf("%v", path)
}
