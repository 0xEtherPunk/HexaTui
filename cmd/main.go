package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"hexatui/internal/ui/model"
)

func main() {
	m := model.InitialModel()
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Ошибка при запуске программы: %v", err)
		os.Exit(1)
	}
} 