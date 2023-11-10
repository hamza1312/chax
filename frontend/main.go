package main

import (
	// "fmt"
	"strconv"
	"encoding/json"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gofiber/fiber/v2"
)

type model struct {
	message string
	currentScreen string
	screens [6]string
	cursor int
}

type ReqBody struct {
	Status bool `json:"success"`
	Msg string `json:"message"`
	Token string `json:"token"`
	Err string `json:"error"`
}

var selectedButton = lipgloss.NewStyle().
	Foreground(lipgloss.Color("0")).
	Background(lipgloss.Color("#4F9852")).
	Padding(1,2,1,2)

var button = lipgloss.NewStyle().
	Foreground(lipgloss.Color("0")).
	Background(lipgloss.Color("#9E9E9E")).
	Padding(1,2,1,2)

func initialModel() model {
	return model{
		screens: [...]string{
			"auth",
			"register",
			"login",
			"servers",
			"channels",
			"channel",
		},
		currentScreen: "auth",
		cursor: 0,
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
				
				case "t":
					agent := fiber.Get("http://127.0.0.1:3000/");
					statusCode, body, errs := agent.Bytes()
					var reqBody ReqBody
					json.Unmarshal(body, &reqBody)
					str := strconv.Itoa(statusCode) + " " + strconv.FormatBool(reqBody.Status) + " " + strconv.Itoa(len(errs))
					return model{message: str}, nil
			}
  }
  return m, nil
}

func (m model) View() string {
	switch m.currentScreen {
		case "auth":
			return selectedButton.Render("register")
	}
   return ""
}

func main() {
	//m := model{}
  p := tea.NewProgram(initialModel(), tea.WithAltScreen())
  p.Run()
}