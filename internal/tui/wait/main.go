package wait

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

//
// Tui 结构
//

type Tui struct {
	width   int // 窗口宽度
	height  int // 窗口高度
	spinner spinner.Model
}

type Keyboard struct {
	esc bool // 是否按下 ESC
}

func (m *Tui) Init() tea.Cmd {
	// 初始化viewport
	return m.spinner.Tick
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
		case tea.KeyEsc, tea.KeyCtrlC:
			consts.GlobalWaitQuit = true
			return m, tea.Quit
		default:
			return m, nil
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *Tui) View() string {
	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).Render(m.UiWait())
}

//
// 客制化内容
//

func NewTui() *Tui {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(consts.ColorDarkGreen))

	return &Tui{
		spinner: s,
	}
}
