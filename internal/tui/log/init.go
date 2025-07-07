package log

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

//
// Tui 结构
//

type Tui struct {
	width    int       // 窗口宽度
	height   int       // 窗口高度
	keyboard *Keyboard // 键盘
}

type Keyboard struct {
	esc bool // 是否按下 ESC
}

func (m *Tui) Init() tea.Cmd {
	return nil
}

func (m *Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.width != msg.Width || m.height != msg.Height {
			m.width = msg.Width
			m.height = msg.Height
			return m, tea.ClearScreen
		}
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			consts.SystemTuiPage = consts.TuiMain
			return m, tea.Quit
		default:
			return m, nil
		}
	default:
		return m, nil
	}

	return m, nil
}

func (m *Tui) View() string {
	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Render(m.UiLogger())
}

//
// 客制化内容
//

func NewTui() *Tui {
	return &Tui{
		keyboard: &Keyboard{
			esc: false,
		},
	}
}
