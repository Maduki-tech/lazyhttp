package gui

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var bodyStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("75"))

type errMsg error

type BodyModel struct {
	textarea      textarea.Model
	err           error
	focused       bool
	width, height int
}

func NewBody() BodyModel {
	ti := textarea.New()
	ti.Placeholder = "Body Info here"

	return BodyModel{
		textarea: ti,
		err:      nil,
	}
}

func (m BodyModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m BodyModel) Update(message tea.Msg) (BodyModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			tea.Println("Pressed ESC from Body")
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(message)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m BodyModel) View() string {
	return bodyStyle.
		Width(m.width).
		Height(m.height).
		Render(m.textarea.View())
}

func (m *BodyModel) SetFocused(focused bool) {
	m.focused = focused
	if focused {
		m.textarea.Focus()
	} else {
		m.textarea.Blur()
	}
}

func (m *BodyModel) SetSize(widht, height int) {
	m.height = height
	m.width = widht
}
