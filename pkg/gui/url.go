package gui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var urlStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("75"))

type URLModel struct {
	focused   bool
	textInput textinput.Model
	width     int
}

func initModel() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "URL"
	ti.CharLimit = 256
	ti.Width = 30

	return ti
}

func NewURL() URLModel {
	return URLModel{
		focused:   true,
		textInput: initModel(),
	}
}

func (m URLModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m URLModel) Update(message tea.Msg) (URLModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(message)
	return m, cmd
}

func (m URLModel) View() string {
	return urlStyle.
		Width(m.width - 2).
		Render(m.textInput.View())
}

func (m *URLModel) SetFocused(focused bool) {
	m.focused = focused
	if focused {
		m.textInput.Focus()
	} else {
		m.textInput.Blur()
	}
}

func (m *URLModel) SetSize(widht int) {
	m.width = widht
}
