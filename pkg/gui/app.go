package gui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ActiveView int

const (
	SidebarView ActiveView = iota
	Url
	Body
	Result
)

type AppModel struct {
	sidebar    SidebarModel
	url        URLModel
	body       BodyModel
	result     ResultModel
	activeView ActiveView
}

func NewAppModel() AppModel {
	return AppModel{
		sidebar:    NewSidebar(),
		url:        NewURL(),
		body:       NewBody(),
		result:     NewResult(),
		activeView: SidebarView,
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "tab":
			// Switch between views
			switch m.activeView {
			case SidebarView:
				m.activeView = Url
				m.sidebar.SetFocused(false)
				m.url.SetFocused(true)
			case Url:
				m.activeView = Body
				m.body.SetFocused(true)
				m.url.SetFocused(false)
			case Body:
				m.activeView = Result
				m.result.SetFocused(true)
				m.url.SetFocused(false)
			case Result:
				m.activeView = SidebarView
				m.result.SetFocused(false)
				m.sidebar.SetFocused(true)
			}
			return m, nil
		}
	}

	switch m.activeView {
	case SidebarView:
		m.sidebar, cmd = m.sidebar.Update(message)
		cmds = append(cmds, cmd)
	case Url:
		m.url, cmd = m.url.Update(message)
		cmds = append(cmds, cmd)
	case Body:
		m.body, cmd = m.body.Update(message)
		cmds = append(cmds, cmd)
	case Result:
		m.result, cmd = m.result.Update(message)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	main := lipgloss.JoinVertical(lipgloss.Top, m.url.View(), m.body.View(), m.result.View())
	return lipgloss.JoinHorizontal(lipgloss.Top, m.sidebar.View(), main)
}

func (m AppModel) GetSidebarView() string {
	return m.sidebar.View()
}

func (m AppModel) GetMainView() string {
	return m.url.View()
}

func (m AppModel) GetActiveView() ActiveView {
	return m.activeView
}
