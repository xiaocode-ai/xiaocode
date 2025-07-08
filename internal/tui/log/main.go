package log

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	"github.com/xiaocode-ai/xiaocode/pkg/xlog"
)

//
// Tui 结构
//

type Tui struct {
	width       int            // 窗口宽度
	height      int            // 窗口高度
	keyboard    *Keyboard      // 键盘
	logSelected int            // 当前选中的日志索引
	viewport    viewport.Model // 用于滚动显示内容的viewport
	showHelp    bool           // 是否显示帮助信息
}

type Keyboard struct {
	esc bool // 是否按下 ESC
}

func (m *Tui) Init() tea.Cmd {
	// 初始化viewport
	return nil
}

func (m *Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.width != msg.Width || m.height != msg.Height {
			m.width = msg.Width
			m.height = msg.Height

			// 更新viewport尺寸
			logContentHeight := m.height - 10 // 预留导航栏和其他元素的空间
			m.viewport.Width = m.width - 4
			m.viewport.Height = logContentHeight

			return m, tea.ClearScreen
		}
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			if m.showHelp {
				m.showHelp = false
				return m, nil
			} else {
				consts.SystemTuiPage = consts.TuiMain
				return m, tea.Quit
			}
		case tea.KeyUp:
			// 基本选择功能
			if m.logSelected > 0 {
				m.logSelected--
			}
			// 视图向上滚动
			m.viewport.ScrollUp(1)
		case tea.KeyDown:
			// 基本选择功能
			if m.logSelected < len(xlog.CustomLogs)-1 {
				m.logSelected++
			}
			// 视图向下滚动
			m.viewport.ScrollDown(1)
		case tea.KeyPgUp: // 添加翻页功能
			// 向上翻页
			m.viewport.PageUp()
			// 同时更新选择的日志项
			m.logSelected -= 10
			if m.logSelected < 0 {
				m.logSelected = 0
			}
		case tea.KeyPgDown: // 添加翻页功能
			// 向下翻页
			m.viewport.PageDown()
			// 同时更新选择的日志项
			m.logSelected += 10
			if m.logSelected >= len(xlog.CustomLogs) {
				m.logSelected = len(xlog.CustomLogs) - 1
			}
		case tea.KeyHome: // 添加回到顶部功能
			m.viewport.GotoTop()
			m.logSelected = 0
		case tea.KeyEnd: // 添加到底部功能
			m.viewport.GotoBottom()
			m.logSelected = len(xlog.CustomLogs) - 1
			m.showHelp = !m.showHelp
		default:
			switch msg.String() {
			case "?", "shift+/", "shift+?":
				m.showHelp = !m.showHelp
			default:
				// 将其他按键事件也传递给viewport处理
				var cmd tea.Cmd
				m.viewport, cmd = m.viewport.Update(msg)
				return m, cmd
			}
		}
	default:
		// 将其他类型的消息也传递给viewport处理
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *Tui) View() string {
	if !m.showHelp {
		return lipgloss.NewStyle().
			Width(m.width).
			Height(m.height).Render(m.UiLogger())
	} else {
		return m.UiShowQuestion()
	}
}

//
// 客制化内容
//

func NewTui() *Tui {
	vp := viewport.New(80, 20)  // 初始尺寸，会在窗口调整时更新
	vp.MouseWheelEnabled = true // 启用鼠标滚轮支持

	return &Tui{
		keyboard: &Keyboard{
			esc: false,
		},
		viewport: vp,
		showHelp: false,
	}
}
