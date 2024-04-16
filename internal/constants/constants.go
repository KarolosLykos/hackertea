package constants

import (
	"strings"
	"time"
)

var CurrentTime = time.Now()

type ItemType string

func (i ItemType) Title() string {
	return strings.ToTitle(string(i))
}

const (
	BaseURL      = "https://hacker-news.firebaseio.com/v0"
	NewSuffix    = "newstories.json"
	TopSuffix    = "topstories.json"
	BestSuffix   = "beststories.json"
	AskSuffix    = "askstories.json"
	SingleSuffix = "item/%s.json"

	TabTop  = "Top"
	TabNew  = "New"
	TabBest = "Best"
	TabAsk  = "Ask"
	Linux   = "linux"
	Windows = "windows"
	Darwin  = "darwin"
)

var Items = struct {
	NewItems   ItemType
	TopItems   ItemType
	BestItems  ItemType
	AskItems   ItemType
	SingleItem ItemType
}{
	NewItems:   "new",
	TopItems:   "top",
	BestItems:  "best",
	AskItems:   "ask",
	SingleItem: "item",
}

const (
	RoundedBorder = "rounded"
	ThickBorder   = "thick"
	DoubleBorder  = "double"
)
