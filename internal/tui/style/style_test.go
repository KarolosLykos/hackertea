package style

import (
	"testing"

	"github.com/KarolosLykos/hackertea/internal/constants"
	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

func TestDocStyle(t *testing.T) {
	docStyle := DocStyle()
	assert.Equal(t, lipgloss.NewStyle().Padding(1, 2, 1, 2), docStyle)
}

func TestTabGapStyle(t *testing.T) {

	tt := []struct {
		name                string
		light, dark, border string
		expectedBorder      lipgloss.Border
	}{
		{
			name:           "default border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         "",
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "double border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.DoubleBorder,
			expectedBorder: lipgloss.DoubleBorder(),
		},
		{
			name:           "rounded border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.RoundedBorder,
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "thick border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.ThickBorder,
			expectedBorder: lipgloss.ThickBorder(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tabGapStyle := TabGapStyle(tc.light, tc.dark, tc.border)

			assert.Equal(t,
				lipgloss.AdaptiveColor{Light: tc.light, Dark: tc.dark},
				tabGapStyle.GetBorderRightForeground())

			assert.Equal(t, tc.expectedBorder.Top, tabGapStyle.GetBorderStyle().Top)
		})
	}
}

func TestActiveTabStyle(t *testing.T) {

	tt := []struct {
		name                string
		light, dark, border string
		expectedBorder      lipgloss.Border
	}{
		{
			name:           "default border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         "",
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "double border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.DoubleBorder,
			expectedBorder: lipgloss.DoubleBorder(),
		},
		{
			name:           "rounded border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.RoundedBorder,
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "thick border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.ThickBorder,
			expectedBorder: lipgloss.ThickBorder(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			activeTabStyle := ActiveTabStyle(tc.light, tc.dark, tc.border)

			assert.Equal(t,
				lipgloss.AdaptiveColor{Light: tc.light, Dark: tc.dark},
				activeTabStyle.GetBorderRightForeground())

			assert.Equal(t, tc.expectedBorder.Top, activeTabStyle.GetBorderStyle().Top)
		})
	}
}

func TestTabBorderWithBottom(t *testing.T) {

	tt := []struct {
		name                        string
		left, middle, right, border string
		expectedBorder              lipgloss.Border
	}{
		{
			name:           "default border",
			left:           "┴",
			middle:         "─",
			right:          "┐",
			border:         "",
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "double border",
			left:           "┴",
			middle:         "─",
			right:          "┐",
			border:         constants.DoubleBorder,
			expectedBorder: lipgloss.DoubleBorder(),
		},
		{
			name:           "rounded border",
			left:           "┴",
			middle:         "─",
			right:          "┐",
			border:         constants.RoundedBorder,
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "thick border",
			left:           "┴",
			middle:         "─",
			right:          "┐",
			border:         constants.ThickBorder,
			expectedBorder: lipgloss.ThickBorder(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tabStyle := tabBorderWithBottom(tc.left, tc.middle, tc.right, tc.border)

			assert.Equal(t, tc.left, tabStyle.BottomLeft)
			assert.Equal(t, tc.middle, tabStyle.Bottom)
			assert.Equal(t, tc.right, tabStyle.BottomRight)
			assert.Equal(t, tc.expectedBorder.Top, tabStyle.Top)
		})
	}
}

func TestWindowStyle(t *testing.T) {

	tt := []struct {
		name                string
		light, dark, border string
		expectedBorder      lipgloss.Border
	}{
		{
			name:           "default border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         "",
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "double border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.DoubleBorder,
			expectedBorder: lipgloss.DoubleBorder(),
		},
		{
			name:           "rounded border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.RoundedBorder,
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "thick border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.ThickBorder,
			expectedBorder: lipgloss.ThickBorder(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			windowStyle := WindowStyle(tc.light, tc.dark, tc.border)

			assert.Equal(t,
				lipgloss.AdaptiveColor{Light: tc.light, Dark: tc.dark},
				windowStyle.GetBorderRightForeground())

			assert.Equal(t, tc.expectedBorder.Top, windowStyle.GetBorderStyle().Top)
		})
	}
}

func TestTitleTabStyle(t *testing.T) {

	tt := []struct {
		name                string
		light, dark, border string
		expectedBorder      lipgloss.Border
	}{
		{
			name:           "default border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         "",
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "double border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.DoubleBorder,
			expectedBorder: lipgloss.DoubleBorder(),
		},
		{
			name:           "rounded border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.RoundedBorder,
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "thick border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.ThickBorder,
			expectedBorder: lipgloss.ThickBorder(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			titleTabStyle := TitleTabStyle(tc.light, tc.dark, tc.border)

			assert.Equal(t,
				lipgloss.AdaptiveColor{Light: tc.light, Dark: tc.dark},
				titleTabStyle.GetBorderRightForeground())

			assert.Equal(t, tc.expectedBorder.Top, titleTabStyle.GetBorderStyle().Top)
		})
	}
}

func TestInActiveTabStyle(t *testing.T) {

	tt := []struct {
		name                string
		light, dark, border string
		expectedBorder      lipgloss.Border
	}{
		{
			name:           "default border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         "",
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "double border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.DoubleBorder,
			expectedBorder: lipgloss.DoubleBorder(),
		},
		{
			name:           "rounded border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.RoundedBorder,
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "thick border",
			light:          "#FFFFFF",
			dark:           "#000000",
			border:         constants.ThickBorder,
			expectedBorder: lipgloss.ThickBorder(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			inActiveStyle := InActiveTabStyle(tc.light, tc.dark, tc.border)

			assert.Equal(t,
				lipgloss.AdaptiveColor{Light: tc.light, Dark: tc.dark},
				inActiveStyle.GetBorderRightForeground())

			assert.Equal(t, tc.expectedBorder.Top, inActiveStyle.GetBorderStyle().Top)
		})
	}
}

func TestItemNormalTitleStyle(t *testing.T) {
	expected := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Padding(0, 0, 0, 2)

	assert.Equal(t, expected, ItemNormalTitleStyle("#FFFFFF", "#000000"))
}

func TestItemNormalDescStyle(t *testing.T) {
	expected := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Padding(0, 0, 0, 2)

	assert.Equal(t, expected, ItemNormalDescStyle("#FFFFFF", "#000000"))
}

func TestItemSelectedTitleStyle(t *testing.T) {
	s := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Padding(0, 0, 0, 1)

	assert.Equal(t, s, ItemSelectedTitleStyle(
		"#FFFFFF", "#000000", "#FFFFFF", "#000000"))
}

func TestItemSelectedDescStyle(t *testing.T) {
	s := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Padding(0, 0, 0, 1)

	assert.Equal(t, s, ItemSelectedDescStyle(
		"#FFFFFF", "#000000", "#FFFFFF", "#000000"))
}

func TestItemDimmedTitleStyle(t *testing.T) {
	s := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Padding(0, 0, 0, 2)

	assert.Equal(t, s, ItemDimmedTitleStyle("#FFFFFF", "#000000"))
}

func TestItemDimmedDescStyle(t *testing.T) {
	s := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Padding(0, 0, 0, 2)

	assert.Equal(t, s, ItemDimmedDescStyle("#FFFFFF", "#000000"))
}

func TestVisitedStyle(t *testing.T) {
	s := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"})

	assert.Equal(t, s, VisitedStyle("#FFFFFF", "#000000"))
}

func TestFilterMatchedStyle(t *testing.T) {
	s := lipgloss.NewStyle().
		Background(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}).
		Underline(true)

	assert.Equal(t, s, FilterMatchedStyle("#FFFFFF", "#000000", "#FFFFFF", "#000000"))
}

func TestTabGapBorderWithBottom(t *testing.T) {

	tt := []struct {
		name           string
		left, border   string
		expectedBorder lipgloss.Border
	}{
		{
			name:           "default border",
			left:           "┴",
			border:         "",
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "double border",
			left:           "┴",
			border:         constants.DoubleBorder,
			expectedBorder: lipgloss.DoubleBorder(),
		},
		{
			name:           "rounded border",
			left:           "┴",
			border:         constants.RoundedBorder,
			expectedBorder: lipgloss.RoundedBorder(),
		},
		{
			name:           "thick border",
			left:           "┴",
			border:         constants.ThickBorder,
			expectedBorder: lipgloss.ThickBorder(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tabStyle := tabGapBorderWithBottom(tc.left, tc.border)

			assert.Equal(t, tc.left, tabStyle.BottomLeft)
			assert.Equal(t, tc.expectedBorder.TopRight, tabStyle.BottomRight)
			assert.Equal(t, "", tabStyle.Right)
		})
	}
}
