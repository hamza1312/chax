package main

import (
	// "fmt"
	// "strconv"
	// "encoding/json"

	// "golang.org/x/crypto/ssh/terminal"
	// "github.com/kopoli/go-terminal-size"
	"os"

	"golang.org/x/term"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/gofiber/fiber/v2"
)

type model struct {
	message      string
	currentState string
	states       [4]string
	cursor       int
}

type ReqBody struct {
	Status bool   `json:"success"`
	Msg    string `json:"message"`
	Token  string `json:"token"`
	Err    string `json:"error"`
}

var selectedButton = lipgloss.NewStyle().
	Foreground(lipgloss.Color("0")).
	Background(lipgloss.Color("#4F9852")).
	Padding(1, 2, 1, 2)

var button = lipgloss.NewStyle().
	Foreground(lipgloss.Color("0")).
	Background(lipgloss.Color("#9E9E9E")).
	Padding(1, 2, 1, 2)

func initialModel() model {
	return model{
		message: "",
		states: [...]string{
			"auth",
			"servers",
			"channels",
			"channel",
		},
		currentState: "auth",
		cursor:       0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// TODO HTTP request handler
// agent := fiber.Get("http://127.0.0.1:3000/");
// statusCode, body, errs := agent.Bytes()
// var reqBody ReqBody
// json.Unmarshal(body, &reqBody)
// str := strconv.Itoa(statusCode) + " " + strconv.FormatBool(reqBody.Status) + " " + strconv.Itoa(len(errs))
// return model{message: str}, nil

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return model{}, tea.Quit

		case "t":
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.currentState {
	case "auth":
		w, h, e := term.GetSize(int(os.Stdin.Fd()))
		if e == nil {
			return lipgloss.Place(
				w,
				h,
				lipgloss.Center,
				lipgloss.Center,
				lipgloss.NewStyle().
					Foreground(lipgloss.Color("0")).
					Background(lipgloss.Color("#4F9852")).
					Padding(1, 2, 1, 2).
					BorderStyle(lipgloss.RoundedBorder()).
					Render(m.states[m.cursor]),
			) //selectedButton.Render("register")
		}
	}
	return ""
}

func main() {
	//m := model{}
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	p.Run()
}
