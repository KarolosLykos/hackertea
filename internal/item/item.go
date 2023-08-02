package item

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"

	"github.com/KarolosLykos/hackertea/internal/constants"
	"github.com/KarolosLykos/hackertea/internal/tui/theme"
)

type Item struct {
	ID          int    `json:"id"`
	Parent      int    `json:"parent"`
	Kids        []int  `json:"kids"`
	Descendants int    `json:"descendants"`
	Parts       []int  `json:"parts"`
	Score       int    `json:"score"`
	Timestamp   int    `json:"time"`
	By          string `json:"by"`
	Type        string `json:"type"`
	Titl        string `json:"title"`
	Text        string `json:"text"`
	URL         string `json:"url"`
	Dead        bool   `json:"dead"`
	Deleted     bool   `json:"deleted"`
	Visited     bool
}

func (i *Item) Time() time.Time {
	return time.Unix(int64(i.Timestamp), 0)
}

func (i *Item) Title() string {
	text := fmt.Sprintf("%s", i.Titl)
	if i.URL != "" {
		text = fmt.Sprintf("%s (%s)", i.Titl, i.URL)
	}

	if i.Visited {
		return visitedStyle().Render(text)
	}

	return text
}

func (i *Item) Description() string {
	desc := fmt.Sprintf(
		"%d points by %s %s ago %d comments",
		i.Score,
		i.By,
		constants.CurrentTime.Sub(i.Time()).Round(time.Second).String(),
		i.Descendants,
	)

	if i.Visited {
		return visitedStyle().Render(desc)
	}

	return desc
}

func (i *Item) FilterValue() string { return i.Titl }

func visitedStyle() lipgloss.Style {
	t, _ := theme.GetTheme()

	return t.Visited
}
