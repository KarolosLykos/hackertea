package utils

import (
	"context"
	"fmt"
	"os/exec"
	"sync"

	"github.com/charmbracelet/bubbles/list"

	"github.com/KarolosLykos/hackertea/internal/constants"
	"github.com/KarolosLykos/hackertea/internal/hn"
	"github.com/KarolosLykos/hackertea/internal/item"
)

// Max returns the maximum of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Open opens a URL in the default web browser for the user's platform.
// The supported platforms are Linux, Windows, and macOS.
func Open(url string, runtimeOS string) error {
	if url == "" {
		return nil
	}

	var err error

	switch runtimeOS {
	case constants.Linux:
		err = exec.Command("xdg-open", url).Start()
	case constants.Windows:
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case constants.Darwing:
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}

// FetchStories fetches the given stories asynchronously from the Hacker News API,
// using a pool of workers.
// It returns a slice of list.Items that can be used to display the stories in a list.
// The function takes the following parameters:
// - ctx: The context to use for the API requests.
// - client: The Hacker News client to use for the API requests.
// - ids: A 2D slice containing the IDs of the stories to fetch for each tab.
// - workers: The number of workers to use for fetching the stories.
// - tabID: The index of the tab containing the IDs of the stories to fetch.
// - start: The index of the first story to fetch.
// - end: The index of the last story to fetch.
func FetchStories(
	ctx context.Context,
	client hn.Service,
	ids [][]int,
	workers, tabID, start, end int,
) []list.Item {
	if tabID > len(ids)-1 {
		return make([]list.Item, 0)
	}

	if end-start > len(ids[tabID]) || end > len(ids[tabID]) {
		return make([]list.Item, 0)
	}

	workers = Min(workers, end-start)

	type workReq struct {
		id     int
		number int
	}
	type workResp struct {
		item   *item.Item
		number int
	}

	work := make(chan workReq)
	msg := make(chan workResp)

	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for s := range work {
				it, err := client.GetItem(ctx, s.id)
				if err != nil {
					msg <- workResp{
						item:   &item.Item{Titl: fmt.Sprintf("Could not get item (%s)", err.Error())},
						number: s.number,
					}
				} else {
					msg <- workResp{item: it, number: s.number}
				}
			}
		}()
	}

	go func() {
		for n, s := range ids[tabID][start:end] {
			work <- workReq{id: s, number: start + n}
		}

		close(work)

		wg.Wait()
		close(msg)
	}()

	items := make([]list.Item, end-start)

	for result := range msg {
		items[result.number-start] = result.item
	}

	return items
}
