package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type listKeyMap struct {
	nextPage      key.Binding
	previousPage  key.Binding
	nextTab       key.Binding
	previousTab   key.Binding
	fetchNextPage key.Binding
}

func NewListKeyMap() *listKeyMap {
	return &listKeyMap{
		previousPage: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "left"),
		),
		nextPage: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "right"),
		),
		nextTab: key.NewBinding(
			key.WithKeys("t", "tab"),
			key.WithHelp("t/Tab", "next tab"),
		),
		previousTab: key.NewBinding(
			key.WithKeys("T", "shift+tab"),
			key.WithHelp("T/Shift+Tab", "previous tab"),
		),
		fetchNextPage: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "next page"),
		),
	}
}

func (l *listKeyMap) KeyBindings() func() []key.Binding {
	return func() []key.Binding {
		return []key.Binding{
			l.nextPage,
			l.previousPage,
			l.nextTab,
			l.previousTab,
			l.fetchNextPage,
		}
	}
}
