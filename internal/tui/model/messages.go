package model

import (
	"github.com/charmbracelet/bubbles/list"
)

type initMsg struct{}

type next struct {
	items []list.Item
}
