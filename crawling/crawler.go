package crawling

import (
	"fmt"
	"strings"
	"sync"

	"github.com/lucasfcosta/wikilinks/api"
)

func FindPath(sourceTitle string, targetTitle string, targetAPI *api.Config) []string {
	visitedLinks := make(map[string]bool)

	requestSemaphore := make(chan int, targetAPI.MaxParallelRequests)
	searchChannel := make(chan []string)
	resultChannel := make(chan []string)

	visitedLinksMutex := &sync.Mutex{}

	go func() []string {
		for {
			select {
			case path := <-searchChannel:
				{
					currentTitle := path[len(path)-1]

					if strings.EqualFold(currentTitle, targetTitle) {
						resultChannel <- path
						break
					}

					visitedLinksMutex.Lock()
					isVisited := visitedLinks[currentTitle]
					visitedLinksMutex.Unlock()

					if isVisited {
						break
					}

					go func() {
						requestSemaphore <- 1
						fmt.Printf("Searching for %s - %v\n", currentTitle, path)
						nextTitles, _ := api.GetAllLinks(currentTitle, targetAPI)
						<-requestSemaphore
						visitedLinksMutex.Lock()
						visitedLinks[currentTitle] = true
						visitedLinksMutex.Unlock()

						for _, nextTitle := range nextTitles {
							newPath := append(path, nextTitle)
							searchChannel <- newPath
						}
					}()
				}

			case result := <-resultChannel:
				{
					defer close(requestSemaphore)
					defer close(searchChannel)
					defer close(resultChannel)
					return result
				}
			}
		}
	}()

	searchChannel <- []string{sourceTitle}
	result := <-resultChannel
	return result
}
