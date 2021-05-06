package main

import "github.com/lucasfcosta/wikilinks/api"

func main() {
	api.GetAllLinks("brazil", api.NewConfigFromEnv())
}
