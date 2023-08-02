package item

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/KarolosLykos/hackertea/internal/constants"
)

func TestItem_Time(t *testing.T) {
	// Test successful conversion
	item := Item{
		Timestamp: int(time.Date(2021, 5, 25, 0, 0, 0, 0, time.Local).Unix()),
	} // May 25, 2021 12:00:00 AM UTC
	expected := time.Date(2021, 5, 25, 0, 0, 0, 0, time.Local)
	actual := item.Time()
	assert.Equal(t, expected, actual)
}

func TestItem_Title(t *testing.T) {
	item := Item{Titl: "Test Title", URL: "https://example.com"}
	expected := "Test Title (https://example.com)"
	assert.Equal(t, expected, item.Title())

	item.Visited = true
	expected = visitedStyle().Render("Test Title (https://example.com)")
	assert.Equal(t, expected, item.Title())
}

func TestItem_Description(t *testing.T) {
	item := Item{
		Score:       42,
		By:          "testuser",
		Timestamp:   int(time.Date(2023, 4, 30, 0, 0, 0, 0, time.Local).Unix()),
		Descendants: 3,
	}
	elapsed := constants.CurrentTime.Sub(item.Time()).Round(time.Second).String()
	expected := fmt.Sprintf("42 points by testuser %s ago 3 comments", elapsed)
	assert.Equal(t, expected, item.Description())

	item.Visited = true
	expected = visitedStyle().Render(fmt.Sprintf("42 points by testuser %s ago 3 comments", elapsed))
	assert.Equal(t, expected, item.Description())
}

func TestItem_FilterValue(t *testing.T) {
	item := Item{Titl: "Test Title"}
	assert.Equal(t, "Test Title", item.FilterValue())
}
