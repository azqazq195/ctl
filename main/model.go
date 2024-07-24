package manager

import (
	"ctl/config"
	"ctl/download"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			BorderStyle(lipgloss.NormalBorder()).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type item struct {
	title       string
	description string
	view        string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list         list.Model
	delegateKeys *delegateKeyMap
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	m.list.SetShowStatusBar(false)
	m.list.SetFilteringEnabled(false)

	// KeyMap
	m.list.KeyMap.Quit.SetEnabled(false)
	m.list.KeyMap.ForceQuit.SetEnabled(false)
	m.list.KeyMap.CloseFullHelp.SetEnabled(false)
	m.list.KeyMap.ShowFullHelp.SetEnabled(false)
	m.list.KeyMap.CancelWhileFiltering.SetEnabled(false)
	m.list.KeyMap.AcceptWhileFiltering.SetEnabled(false)

	return appStyle.Render(m.list.View())
}

func New(cfg *config.Config) Model {
	var (
		delegateKeys = newDelegateKeyMap()
	)

	var items []list.Item
	if len(cfg.Services.Downloads) > 0 {
		items = append(items, item{
			title:       "Downloads",
			description: "Download dependency",
			view:        download.New().View(),
		})
	}
	//if len(cfg.Services.Installs) > 0 {
	//	items = append(items, item{
	//		title:       "Installs",
	//		description: "Install application",
	//	})
	//}
	//if len(cfg.Services.Runs) > 0 {
	//	items = append(items, item{
	//		title:       "Runs",
	//		description: "Run application",
	//	})
	//}

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	groceryList := list.New(items, delegate, 0, 0)
	groceryList.Title = cfg.Name
	groceryList.Styles.Title = titleStyle

	return Model{
		list:         groceryList,
		delegateKeys: delegateKeys,
	}
}
