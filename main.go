package main

import (
	"ctl/config"
	"ctl/manager"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	cfg := config.LoadConfig()

	//fmt.Println("Config Name:", cfg.Name)
	//fmt.Println("Downloads:")
	//for name, service := range cfg.Services.Downloads {
	//	fmt.Printf("  %s:\n", name)
	//	fmt.Printf("    Description: %s\n", service.Description)
	//	fmt.Println("    URLs:")
	//	for _, url := range service.URLs {
	//		fmt.Printf("      - %s\n", url)
	//	}
	//	fmt.Printf("    Required: %t\n", service.Required)
	//}
	//
	//fmt.Println("Installs:")
	//for name, service := range cfg.Services.Installs {
	//	fmt.Printf("  %s:\n", name)
	//	fmt.Printf("    Path: %s\n", service.Path)
	//}
	//
	//fmt.Println("Runs:")
	//for name, service := range cfg.Services.Runs {
	//	fmt.Printf("  %s:\n", name)
	//	fmt.Printf("    Path: %s\n", service.Path)
	//}

	p := tea.NewProgram(manager.New(&cfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
