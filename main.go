package main

import (
	"fmt"

	"github.com/lucasfcosta/wikilinks/api"
)

func main() {
	fmt.Println(api.GetAllLinks("brazil", api.NewConfigFromEnv()))
}
