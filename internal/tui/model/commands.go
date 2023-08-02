package model

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/KarolosLykos/hackertea/internal/utils"
)

func (m model) initCmd() tea.Cmd {
	return func() tea.Msg {
		docH, docV := m.theme.Doc.GetFrameSize()
		winH, _ := m.theme.Window.GetFrameSize()
		contH, contV := m.theme.ListContent.GetFrameSize()
		for i := range m.tabs {
			m.TabContent[i].SetSize(
				m.width-docH-winH-contH,
				m.height-docV-contV,
			)
			m.visited[i] = map[int]bool{}
			m.TabContent[i].SetItems(
				utils.FetchStories(m.ctx, m.client, m.ids, m.cfg.Workers, i, 0, m.TabContent[i].Paginator.PerPage),
			)
		}

		return initMsg{}
	}
}

func (m model) next(tabID, perPage, page int) tea.Cmd {
	return func() tea.Msg {
		n := next{}
		n.items = utils.FetchStories(m.ctx, m.client, m.ids, m.cfg.Workers, tabID, perPage*page, perPage*page+perPage)

		return n
	}
}
