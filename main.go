package main

import (
	"ctl/config"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
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
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{}
}

type model struct {
	list         list.Model
	delegateKeys *delegateKeyMap
}

func newModel() model {
	var (
		delegateKeys = newDelegateKeyMap()
	)

	// Make initial list of items
	const numItems = 2
	items := make([]list.Item, numItems)
	items[0] = item{title: "Download", description: "Download dependency"}
	items[1] = item{title: "Install", description: "Install application"}
	//items[2] = item{title: "App Manage", description: "Application processing"}

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	groceryList := list.New(items, delegate, 0, 0)
	groceryList.Title = "EzisCtl"
	groceryList.Styles.Title = titleStyle

	return model{
		list:         groceryList,
		delegateKeys: delegateKeys,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	m.list.SetShowStatusBar(false)
	return appStyle.Render(m.list.View())
}

//func main() {
//	config := config.LoadConfig()
//	fmt.Println(config)
//	//if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
//	//	fmt.Println("Error running program:", err)
//	//	os.Exit(1)
//	//}
//}

func main() {
	cfg := config.LoadConfig()

	fmt.Println("Config Name:", cfg.Name)
	fmt.Println("Downloads:")
	for name, service := range cfg.Services.Downloads {
		fmt.Printf("  %s:\n", name)
		fmt.Printf("    Description: %s\n", service.Description)
		fmt.Println("    URLs:")
		for _, url := range service.URLs {
			fmt.Printf("      - %s\n", url)
		}
		fmt.Printf("    Required: %t\n", service.Required)
	}

	fmt.Println("Installs:")
	for name, service := range cfg.Services.Installs {
		fmt.Printf("  %s:\n", name)
		fmt.Printf("    Path: %s\n", service.Path)
	}

	fmt.Println("Runs:")
	for name, service := range cfg.Services.Runs {
		fmt.Printf("  %s:\n", name)
		fmt.Printf("    Path: %s\n", service.Path)
	}
}
