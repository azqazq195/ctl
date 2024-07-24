package run

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	title string
}

func New() Model {
	return Model{
		title: "Run",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.title
}
