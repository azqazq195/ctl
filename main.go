package main

import (
	"ctl/run"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"
)

type TickMsg time.Time
type Model struct {
	name string
	num  int
}

func main() {
	//cfg := config.LoadConfig()
	//p := tea.NewProgram(main.New(&cfg), tea.WithAltScreen())

	p := tea.NewProgram(run.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
