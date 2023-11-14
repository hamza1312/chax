package main

import (
	// "fmt"
	// "strconv"
	// "encoding/json"

	// "golang.org/x/crypto/ssh/terminal"
	// "github.com/kopoli/go-terminal-size"
	"os"
	// "strings"

	"golang.org/x/term"
	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/gofiber/fiber/v2"
)

type model struct {
	input        textinput.Model
	hint				 string
	fullInput    string
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
	ti := textinput.New()
	ti.Focus()

	return model{
		input: ti,
		hint: "",
		fullInput: "",
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
	return textinput.Blink
}

// func (m model) addChar(char string) model {
// 	if (len(char) == 1) {
// 		m.input = m.input[:m.cursor] + char + m.input[m.cursor:]
// 		m.cursor++
// 	}
// 	return m
// }

// TODO HTTP request handler
// agent := fiber.Get("http://127.0.0.1:3000/");
// statusCode, body, errs := agent.Bytes()
// var reqBody ReqBody
// json.Unmarshal(body, &reqBody)
// str := strconv.Itoa(statusCode) + " " + strconv.FormatBool(reqBody.Status) + " " + strconv.Itoa(len(errs))
// return model{message: str}, nil

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
			case "ctrl+c":
				return model{}, tea.Quit
			
			case "enter":
				m.input.SetValue("")

		// case "left":
		// 	if m.cursor > 0 {
    //     m.cursor--
    //   }
    //   return m, nil

    // case "right":
		// 	if m.cursor < len(m.input) {
    //     m.cursor++
    //   }
		// 	return m, nil

		// case "backspace":
		// 	if m.cursor > 0 {
    //     m.input = m.input[:m.cursor-1] + m.input[m.cursor:]
    //     m.cursor--
    //   }
    //   return m, nil

    // case "delete":
		// 	if m.cursor < len(m.input) {
    //     m.input = m.input[:m.cursor] + m.input[m.cursor+1:]
    //   }
		// 	return m, nil

		// default:
		// 	m = m.addChar(msg.String())
		// 	return m, nil
		}
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	w, h, e := term.GetSize(int(os.Stdin.Fd()))
	// switch m.currentState {
	// case "auth":
	// 	w, h, e := term.GetSize(int(os.Stdin.Fd()))
	// 	if e == nil {
	// 		return lipgloss.Place(
	// 			w,
	// 			h,
	// 			lipgloss.Center,
	// 			lipgloss.Center,
	// 			lipgloss.NewStyle().
	// 				Foreground(lipgloss.Color("0")).
	// 				Background(lipgloss.Color("#4F9852")).
	// 				Padding(1, 2, 1, 2).
	// 				BorderStyle(lipgloss.RoundedBorder()).
	// 				Render(m.input),
	// 		)
	// 	}
	// }
	if e == nil {
		return lipgloss.Place(w, h-1, lipgloss.Left, lipgloss.Bottom, m.input.View())
	}
	return ""
}

func main() {
	//m := model{}
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	p.Run()
}
