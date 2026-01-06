package gui

import tea "github.com/charmbracelet/bubbletea"

type MainModel struct {
	focused bool
}

func NewMainView() MainModel {
	return MainModel{
		focused: true,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(message tea.Msg) (MainModel, tea.Cmd) {
	return m, nil
}

func (m MainModel) View() string {
	return "Main"
}
