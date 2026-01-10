// Package gui is for managing the UI of the TUI
package gui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("75"))

type SidebarModel struct {
	focused       bool
	table         table.Model
	width, height int
}

func NewSidebar() SidebarModel {
	table := newTable()
	return SidebarModel{
		focused: true,
		table:   table,
	}
}

func newTable() table.Model {
	columns := []table.Column{
		{Title: "Method", Width: 6},
		{Title: "Name", Width: 20},
	}

	rows := []table.Row{
		{"GET", "/api/v1/users"},
		{"POST", "/api/v1/users"},
		{"DELETE", "/api/v1/users/:id"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("81")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	return t
}

func (m SidebarModel) Init() tea.Cmd {
	return nil
}

func (m SidebarModel) Update(message tea.Msg) (SidebarModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Batch(
				tea.Printf("Lets do this %s", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(message)
	return m, cmd
}

func (m SidebarModel) View() string {
	return style.
		Height(m.height - 2).
		Width(m.width - 2).
		Render(m.table.View())
}

func (m *SidebarModel) SetFocused(focused bool) {
	m.focused = focused
	m.table.Focus()
	if !focused {
		m.table.Blur()
	}
}

func (m *SidebarModel) SetSize(widht, height int) {
	m.height = height
	m.width = widht
}
