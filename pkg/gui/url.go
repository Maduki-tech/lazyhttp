package gui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type URLModel struct {
	focused   bool
	textInput textinput.Model
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
	return m.textInput.View()
}

func (m *URLModel) SetFocused(focused bool) {
	m.focused = focused
	if focused {
		m.textInput.Focus()
	} else {
		m.textInput.Blur()
	}
}
