package index

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

//
// Tui 结构
//

type Tui struct {
	spinner  spinner.Model  // 加载动画
	width    int            // 窗口宽度
	height   int            // 窗口高度
	keyboard *Keyboard      // 键盘
	textArea textarea.Model // 文本框
	chat     *Chat          // 聊天
	custom   *Custom        // 客制化内容
}

type Keyboard struct {
	esc bool // 是否按下 ESC
}

type Custom struct {
	plzInputContent bool // 请输入内容
	aiReady         bool // AI 是否准备好
}

type Chat struct {
	chatId string         // 聊天ID
	chat   []*CurrentChat // 聊天内容
}

type CurrentChat struct {
	user    string // 交流对象
	content string // 交流内容
}

func (m *Tui) Init() tea.Cmd {
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
				case 3:
					consts.SystemTuiPage = consts.TuiAi
					return m, tea.Quit
				case 5:
					consts.SystemTuiPage = consts.TuiLog
					return m, tea.Quit
				default:
					consts.SystemTuiPage = consts.TuiNil
					return m, tea.Quit
				}
			default:
				return m, nil
			}
		} else {
			switch msg.Type {
			case tea.KeyShiftRight:
				currentText := m.textArea.Value()
				newText := currentText + "\n"
				m.textArea.SetValue(newText)
				m.textArea.SetHeight(m.textArea.Height() + 1)
				return m, nil
			case tea.KeyEnter:
				m.textArea.SetHeight(1)
				if m.textArea.Value() != "" && m.custom.aiReady {
					m.chat.chat = append(m.chat.chat, &CurrentChat{
						user:    "user",
						content: m.textArea.Value(),
					})
				} else {
					m.DisplayEnterContent()
				}
				m.textArea.Reset()
				return m, nil
			case tea.KeyEsc, tea.KeyCtrlC:
				m.keyboard.esc = !m.keyboard.esc
				return m, nil
			default:
				var cmd tea.Cmd
				m.textArea, cmd = m.textArea.Update(msg)
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

func (m *Tui) View() string {
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

//
// 客制化内容
//

func NewTui() *Tui {
	spin := spinner.New()
	spin.Spinner = spinner.MiniDot
	spin.Style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorDarkGreen))

	ta := textarea.New()
	ta.Focus()
	ta.CharLimit = 0
	ta.SetHeight(1)
	ta.Prompt = ""
	ta.ShowLineNumbers = false

	return &Tui{
		spinner:  spin,
		textArea: ta,
		keyboard: &Keyboard{
			esc: false,
		},
		chat: &Chat{
			chatId: uuid.New().String(),
			chat:   []*CurrentChat{},
		},
		custom: &Custom{
			plzInputContent: false,
			aiReady:         false,
		},
	}
}
