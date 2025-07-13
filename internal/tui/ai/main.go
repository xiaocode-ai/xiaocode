package ai

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	"strings"
)

//
// Tui 结构
//

type Tui struct {
	width             int            // 窗口宽度
	height            int            // 窗口高度
	viewport          viewport.Model // 用于滚动显示内容的viewport
	selectedAI        int            // 当前选中的AI索引
	showHelp          bool           // 是否显示帮助信息
	showCreate        bool           // 是否显示创建对话框
	createInputSelect int            // 创建对话框输入框选择索引
	createModel       *CreateModel   // 创建对话框模型
}

type CreateModel struct {
	AiLink        textinput.Model // AI链接
	AiKey         textinput.Model // AI密钥
	AiDescription textinput.Model // AI描述
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
		case tea.KeyEsc, tea.KeyCtrlC:
			if m.showHelp {
				m.showHelp = false
			} else if m.showCreate {
				m.showCreate = false
			} else {
				consts.SystemTuiPage = consts.TuiMain
				return m, tea.Quit
			}
		case tea.KeyUp:
			if m.selectedAI > 0 {
				m.selectedAI--
			}
		case tea.KeyDown:
			if m.selectedAI < consts.AiCount-1 {
				m.selectedAI++
			}
		default:
			switch strings.ToLower(msg.String()) {
			case "?", "shift+/", "shift+?":
				m.showHelp = !m.showHelp
			case "n":
				if !m.showCreate {
					m.showCreate = true
				}
			default:
				if m.showCreate {
					var cmd tea.Cmd
					switch m.createInputSelect {
					case 0:
						m.createModel.AiLink, cmd = m.createModel.AiLink.Update(msg)
					case 1:
						m.createModel.AiKey, cmd = m.createModel.AiKey.Update(msg)
					case 2:
						m.createModel.AiDescription, cmd = m.createModel.AiDescription.Update(msg)
					}
					return m, cmd
				}
			}
		}
	default:
		return m, nil
	}

	return m, nil
}

func (m *Tui) View() string {
	if !m.showHelp {
		if !m.showCreate {
			return lipgloss.NewStyle().
				Width(m.width).
				Height(m.height).Render(m.UiAi())
		} else {
			return m.UiShowCreate()
		}
	} else {
		return m.UiShowQuestion()
	}
}

//
// 客制化内容
//

func NewTui() *Tui {
	// 初始化viewport
	vp := viewport.New(80, 20)
	vp.MouseWheelEnabled = true

	// 创建输入框
	aiLinkInput := textinput.New()
	aiLinkInput.Placeholder = ""
	aiLinkInput.Focus()

	aiKeyInput := textinput.New()
	aiKeyInput.Placeholder = ""
	aiKeyInput.Focus()

	aiDescriptionInput := textinput.New()
	aiDescriptionInput.Placeholder = ""
	aiDescriptionInput.Focus()

	return &Tui{
		selectedAI: 0,
		viewport:   vp,
		showHelp:   false,
		showCreate: false,
		createModel: &CreateModel{
			AiLink:        aiLinkInput,
			AiKey:         aiKeyInput,
			AiDescription: aiDescriptionInput,
		},
	}
}
