package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// common behavior for a Tui based on Bubbletea
// Value to get a state value
// Update to update a state value
// View to display a state value
type BBTui interface {
	Blink() tea.Msg
	Blur() tea.Msg
	Value() string
	Update(tea.Msg) (BBTui, tea.Cmd)
	View() string
}

type Question struct {
	question string
	answer   string
	tui      BBTui
}

// manage the CLI flow
type Container struct {
	index     int
	questions []Question
	done      bool
}

// **********************************
// ***** 		CONTAINER		*****
// **********************************

func (c Container) Init() tea.Cmd {
	return c.questions[c.index].tui.Blink
}

func (m Container) View() string {
	var output string
	current := m.questions[m.index]
	if m.done {
		for _, q := range m.questions {
			output += fmt.Sprintf("%s: %s\n", q.question, q.answer)
		}
		return output
	}
	inputView := fmt.Sprintf("%s \n %s", current.question, current.tui.View())

	return inputView
}

func (m Container) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.questions[m.index].answer != "" {
		m.Next()
	}
	if m.index == len(m.questions)-1 {
		m.done = true
	}
	current := &m.questions[m.index]
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.index == len(m.questions)-1 {
				m.done = true
			}
			current.answer = current.tui.Value()
			m.Next()
			return m, current.tui.Blur
		}
	}
	current.tui, cmd = current.tui.Update(msg)
	return m, cmd
}
func New(q []Question) *Container {
	return &Container{
		questions: q,
	}
}

func (m *Container) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func (m Container) GetQuestions() string {
	return m.questions[0].answer
}

// **********************************
// ***** 		TEXTFIELD		*****
// **********************************

type TextField struct {
	input textinput.Model
}

func (i *TextField) Blink() tea.Msg {
	return textinput.Blink()
}

func (i *TextField) Value() string {
	return i.input.Value()
}

func (i *TextField) View() string {
	return i.input.View()
}

func (i *TextField) Blur() tea.Msg {
	return i.input.Blur
}

func newTextField(placeholder string) *TextField {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	textField := TextField{}
	textField.input = ti

	return &textField
}

func NewQuestion(q string, ans string) Question {
	return Question{
		question: q,
		answer:   ans,
		tui:      newTextField(""),
	}
}

func (i *TextField) Update(msg tea.Msg) (BBTui, tea.Cmd) {
	var cmd tea.Cmd
	i.input, cmd = i.input.Update(msg)
	return i, cmd
}
