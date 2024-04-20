package model

import (
	"context"
	"runtime"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/KarolosLykos/hackertea/internal/config"
	"github.com/KarolosLykos/hackertea/internal/constants"
	"github.com/KarolosLykos/hackertea/internal/hn"
	"github.com/KarolosLykos/hackertea/internal/item"
	"github.com/KarolosLykos/hackertea/internal/tui/keys"
	"github.com/KarolosLykos/hackertea/internal/tui/theme"
	"github.com/KarolosLykos/hackertea/internal/utils"
)

type model struct {
	ctx           context.Context
	cancel        context.CancelFunc
	cfg           *config.Config
	theme         *theme.Theme
	tabs          []string
	TabContent    []list.Model
	activeTab     int
	loading       bool
	client        *hn.HN
	spinner       spinner.Model
	ids           [][]int
	visited       []map[int]bool
	width, height int
}

func New(ctx context.Context, client *hn.HN) (*model, error) {
	newCtx, cancel := context.WithCancel(ctx)

	cfg, err := config.LoadConfig()
	if err != nil {
		cancel()
		return nil, err
	}

	th, err := theme.NewTheme()
	if err != nil {
		cancel()
		return nil, err
	}

	s := spinner.New()
	s.Spinner = spinner.Points

	topStories, err := client.GetItems(newCtx, constants.Items.TopItems)
	if err != nil {
		cancel()
		return nil, err
	}

	bestStories, err := client.GetItems(newCtx, constants.Items.BestItems)
	if err != nil {
		cancel()
		return nil, err
	}

	newStories, err := client.GetItems(newCtx, constants.Items.NewItems)
	if err != nil {
		cancel()
		return nil, err
	}

	askStories, err := client.GetItems(newCtx, constants.Items.AskItems)
	if err != nil {
		cancel()
		return nil, err
	}

	m := &model{
		cfg:     cfg,
		ctx:     newCtx,
		cancel:  cancel,
		theme:   th,
		ids:     [][]int{topStories, newStories, bestStories, askStories},
		client:  client,
		spinner: s,
		visited: make([]map[int]bool, 4),
		tabs:    []string{constants.TabTop, constants.TabNew, constants.TabBest, constants.TabAsk},
	}

	m.TabContent = m.createTabContent(4)

	listKeys := keys.NewListKeyMap()

	for i := 0; i < len(m.TabContent); i++ {
		m.TabContent[i].AdditionalShortHelpKeys = listKeys.KeyBindings()
		m.TabContent[i].AdditionalFullHelpKeys = listKeys.KeyBindings()
	}

	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case next:
		m.loading = false
		items := append(m.TabContent[m.activeTab].Items(), msg.items...)
		m.TabContent[m.activeTab].SetItems(items)
		m.visited[m.activeTab][m.TabContent[m.activeTab].Paginator.Page] = true
		m.TabContent[m.activeTab].Paginator.NextPage()
	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.TabContent[m.activeTab].FilterState() == list.Filtering {
			break
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case tea.KeyEnter.String():
			if v, ok := m.TabContent[m.activeTab].SelectedItem().(*item.Item); ok {
				if err := utils.Open(v.URL, runtime.GOOS); err != nil {
					return m, nil
				}
				v.Visited = true
			}
		case "n":
			if !m.visited[m.activeTab][m.TabContent[m.activeTab].Paginator.Page] {
				m.loading = true
				return m, m.next(m.activeTab, m.TabContent[m.activeTab].Paginator.PerPage, m.TabContent[m.activeTab].Paginator.Page+1)
			}
		case "t", "tab":
			m.activeTab = utils.Min(m.activeTab+1, len(m.tabs)-1)
		case "T", "shift+tab":
			m.activeTab = utils.Max(m.activeTab-1, 0)
		}
	case initMsg:
		m.loading = false

	case spinner.TickMsg:
		if !m.loading {
			return m, nil
		}

		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.loading = true
		return m, tea.Batch(
			m.spinner.Tick,
			m.initCmd(),
		)
	}

	m.TabContent[m.activeTab], cmd = m.TabContent[m.activeTab].Update(msg)

	return m, cmd
}

func (m model) View() string {
	doc := strings.Builder{}
	doc.Reset()

	windowFrameSize := m.theme.Window.GetHorizontalFrameSize()
	docFrameSize := m.theme.Doc.GetHorizontalFrameSize()
	var renderedTabs []string

	renderedTabs = append(renderedTabs, m.theme.TitleTab.Render("HackerTea"))

	for i, t := range m.tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.tabs)-1, i == m.activeTab
		if isActive {
			style = m.theme.ActiveTab.Copy()
		} else {
			style = m.theme.InActiveTab.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "┘"
		} else if isFirst && !isActive {
			border.BottomLeft = "┴"
		} else if isLast && isActive {
			border.BottomRight = "└"
		} else if isLast && !isActive {
			border.BottomRight = "┴"
		}

		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	gap := m.theme.GapTab.Render(
		strings.Repeat(" ", utils.Max(0, m.width-windowFrameSize-docFrameSize-lipgloss.Width(row))),
	)

	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
	doc.WriteString(row)
	doc.WriteString("\n")

	m.theme.Window.Width(m.width - windowFrameSize - docFrameSize - 1)

	if m.loading {
		doc.WriteString(m.theme.Window.Render(m.spinner.View()))
	} else {
		doc.WriteString(m.theme.Window.Render(m.TabContent[m.activeTab].View()))
	}

	return m.theme.Doc.Render(doc.String())
}

func (m model) createTabContent(tabs int) []list.Model {
	tabContent := make([]list.Model, tabs)

	delegate := list.NewDefaultDelegate()
	delegate.Styles = list.DefaultItemStyles{
		NormalTitle:   m.theme.NormalTitle,
		NormalDesc:    m.theme.NormalDesc,
		SelectedTitle: m.theme.SelectedTitle,
		SelectedDesc:  m.theme.SelectedDesc,
		DimmedTitle:   m.theme.DimmedTitle,
		DimmedDesc:    m.theme.DimmedDesc,
		FilterMatch:   m.theme.FilterMatch,
	}

	for i := 0; i < tabs; i++ {
		l := list.New(make([]list.Item, 0), delegate, 0, 0)
		l.SetShowTitle(false)
		l.SetShowStatusBar(false)

		tabContent[i] = l
	}

	return tabContent
}
