package gui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ResultModel struct {
	viewport viewport.Model
	focused  bool
}

const data = `
	# Result
	## Test
	`

func NewResult() ResultModel {
	const width = 78

	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	vp.SetContent(data)

	return ResultModel{
		viewport: vp,
	}
}

func (m ResultModel) Init() tea.Cmd {
	return nil
}

func (m ResultModel) Update(message tea.Msg) (ResultModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.viewport, cmd = m.viewport.Update(message)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m ResultModel) View() string {
	return m.viewport.View()
}

func (m *ResultModel) SetFocused(focused bool) {
	m.focused = focused
}
