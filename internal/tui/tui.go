package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Tui
type Tui struct {
	spinner   spinner.Model
	width     int
	height    int
	custom    *CustomTui
	textInput textinput.Model
}

type CustomTui struct {
	loading bool
}

func (m Tui) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.width != msg.Width || m.height != msg.Height {
			m.width = msg.Width
			m.height = msg.Height
			return m, tea.ClearScreen
		}
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			m.textInput.Reset()
		default:
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Tui) View() string {
	// 设置全局底色，覆盖整个屏幕
	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Render(m.UiMain())
}

func New(custom *CustomTui) tea.Model {
	spin, input := GetInit(custom)

	return Tui{
		custom:    custom,
		spinner:   spin,
		textInput: input,
		width:     0,
		height:    0,
	}
}
