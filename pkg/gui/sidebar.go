// Package gui is for managing the UI of the TUI
package gui

import tea "github.com/charmbracelet/bubbletea"

type SidebarModel struct {
	focused bool
}

func NewSidebar() SidebarModel {
	return SidebarModel{
		focused: true,
	}
}

func (m SidebarModel) Init() tea.Cmd {
	return nil
}

func (m SidebarModel) Update(message tea.Msg) (SidebarModel, tea.Cmd) {
	return m, nil
}

func (m SidebarModel) View() string {
	return "Sidebar"
}
