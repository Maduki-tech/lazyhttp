package main

import (
	"fmt"
	"log"
	"os"

	"lazyhttp/pkg/gui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	sidebar  gui.SidebarModel
	mainview gui.MainModel
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "enter":
			log.Println("Enter Pressed")
		}
	}
	return m, nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.sidebar.View(), m.mainview.View())
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	m := model{
		sidebar:  gui.NewSidebar(),
		mainview: gui.NewMainView(),
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
