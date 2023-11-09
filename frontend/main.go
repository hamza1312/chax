package main

import (
	// "fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	message string
}

func initialModel() model {
	return model{
		message: "Hello World!",
	}
}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c":
      return model{}, tea.Quit
    }
  }
  return m, nil
}

func (m model) View() string {
  return m.message
}

func main() {
	//m := model{}
  p := tea.NewProgram(initialModel(), tea.WithAltScreen())
  p.Run()
}