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
	custom    *CustomOperate
	keyboard  *CustomKeyboard
	textInput textinput.Model
}

type CustomOperate struct {
	loading bool
}

type CustomKeyboard struct {
	esc bool
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
		if m.keyboard.esc {
			switch msg.Type {
			case tea.KeyEsc:
				m.keyboard.esc = false
			case tea.KeyDown:
				if selectedMenuIndex < len(menuItem)-1 {
					selectedMenuIndex++
				}
			case tea.KeyUp:
				if selectedMenuIndex > 0 {
					selectedMenuIndex--
				}
			case tea.KeyEnter:
				switch selectedMenuIndex {
				case 0:
					m.textInput.Reset()
				case 1:
					m.textInput.Reset()
				case 4:
					return m, tea.Quit
				}
			default:
				return m, nil
			}
		} else {
			switch msg.Type {
			case tea.KeyCtrlC:
				return m, tea.Quit
			case tea.KeyEnter:
				m.textInput.Reset()
			case tea.KeyEsc:
				m.keyboard.esc = !m.keyboard.esc
			default:
				var cmd tea.Cmd
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Tui) View() string {
	// 首先渲染主界面作为底层
	mainUI := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Render(m.UiMain())

	// 如果按下 ESC，在主界面之上显示 ESC 菜单
	if m.keyboard.esc {
		return m.UiPushEsc()
	}

	// 没有按 ESC，只显示主界面
	return mainUI
}

func New(custom *CustomOperate, keyboard *CustomKeyboard) tea.Model {
	spin, input := GetInit(custom)

	return Tui{
		custom:    custom,
		keyboard:  keyboard,
		spinner:   spin,
		textInput: input,
		width:     0,
		height:    0,
	}
}
