package theme

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/KarolosLykos/hackertea/internal/config"
	"github.com/KarolosLykos/hackertea/internal/tui/style"
)

var instance *Theme

type Theme struct {
	TitleTab      lipgloss.Style
	NormalTitle   lipgloss.Style
	NormalDesc    lipgloss.Style
	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style
	DimmedTitle   lipgloss.Style
	DimmedDesc    lipgloss.Style
	FilterMatch   lipgloss.Style
	Visited       lipgloss.Style
	ActiveTab     lipgloss.Style
	InActiveTab   lipgloss.Style
	GapTab        lipgloss.Style
	Window        lipgloss.Style
	Doc           lipgloss.Style
	ListContent   lipgloss.Style
}

func GetTheme() (*Theme, error) {
	if instance != nil {
		return instance, nil
	}

	instance, _ = NewTheme()

	return instance, nil
}

func NewTheme() (*Theme, error) {
	cfg, _ := config.LoadConfig()

	return &Theme{
		TitleTab: style.TitleTabStyle(cfg.Style.Tab.Color.Light, cfg.Style.Tab.Color.Dark, cfg.Style.Window.Border),
		NormalTitle: style.ItemNormalTitleStyle(
			cfg.Style.ListItem.NormalTitle.Light,
			cfg.Style.ListItem.NormalTitle.Dark,
		),
		NormalDesc: style.ItemNormalDescStyle(cfg.Style.ListItem.NormalDesc.Light, cfg.Style.ListItem.NormalDesc.Dark),
		SelectedTitle: style.ItemSelectedTitleStyle(
			cfg.Style.ListItem.SelectedTitle.BorderForeground.Light,
			cfg.Style.ListItem.SelectedTitle.BorderForeground.Dark,
			cfg.Style.ListItem.SelectedTitle.Foreground.Light,
			cfg.Style.ListItem.SelectedTitle.Foreground.Dark,
		),
		SelectedDesc: style.ItemSelectedDescStyle(
			cfg.Style.ListItem.SelectedTitle.BorderForeground.Light,
			cfg.Style.ListItem.SelectedTitle.BorderForeground.Dark,
			cfg.Style.ListItem.SelectedDesc.Light,
			cfg.Style.ListItem.SelectedDesc.Dark,
		),
		DimmedTitle: style.ItemDimmedTitleStyle(cfg.Style.ListItem.DimmedTitle.Light, cfg.Style.ListItem.DimmedTitle.Dark),
		DimmedDesc:  style.ItemDimmedDescStyle(cfg.Style.ListItem.DimmedDesc.Light, cfg.Style.ListItem.DimmedDesc.Dark),
		Visited:     style.VisitedStyle(cfg.Style.Visited.Light, cfg.Style.Visited.Dark),
		FilterMatch: style.FilterMatchedStyle(
			cfg.Style.ListItem.FilterMatch.BorderForeground.Light,
			cfg.Style.ListItem.FilterMatch.BorderForeground.Dark,
			cfg.Style.ListItem.FilterMatch.Foreground.Light,
			cfg.Style.ListItem.FilterMatch.Foreground.Dark,
		),
		ActiveTab:   style.ActiveTabStyle(cfg.Style.Tab.Color.Light, cfg.Style.Tab.Color.Dark, cfg.Style.Window.Border),
		InActiveTab: style.InActiveTabStyle(cfg.Style.Tab.Color.Light, cfg.Style.Tab.Color.Dark, cfg.Style.Window.Border),
		GapTab:      style.TabGapStyle(cfg.Style.Tab.Color.Light, cfg.Style.Tab.Color.Dark, cfg.Style.Window.Border),
		Window:      style.WindowStyle(cfg.Style.Window.Color.Light, cfg.Style.Window.Color.Dark, cfg.Style.Window.Border),
		Doc:         style.DocStyle(),
		ListContent: lipgloss.NewStyle().Padding(5),
	}, nil
}
