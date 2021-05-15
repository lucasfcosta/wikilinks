package crawling

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lucasfcosta/wikilinks/api"
)

func FindPath(sourceTitle string, targetTitle string, targetAPI *api.Config) ([]string, error) {
	visitedLinks := make(map[string]bool)

	queue := [][]string{[]string{sourceTitle}}
	currentDepth := 0

	for len(queue) > 0 {
		path := queue[0]
		currentTitle := path[len(path)-1]
		queue = queue[1:]

		if len(path) > currentDepth {
			fmt.Printf("Moving to depth %d\n", currentDepth+1)
			currentDepth = currentDepth + 1
		}

		if strings.EqualFold(currentTitle, targetTitle) {
			return path, nil
		}

		if visitedLinks[currentTitle] {
			continue
		}

		fmt.Printf("Searching for %s - %v\n", currentTitle, path)
		nextTitles, _ := api.GetAllLinks(currentTitle, targetAPI)
		visitedLinks[currentTitle] = true

		for _, nextTitle := range nextTitles {
			newPath := append(path, nextTitle)
			queue = append(queue, newPath)
		}
	}

	return nil, errors.New("Path not found.")
}
