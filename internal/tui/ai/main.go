package ai

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	"gorm.io/gorm"
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
	showCreateErr     string         // 创建对话框错误信息
	createInputSelect int            // 创建对话框输入框选择索引
	createModel       *CreateModel   // 创建对话框模型
	db                *gorm.DB       // 数据库连接
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
			if m.showCreate {
				if m.createInputSelect > 0 {
					m.createInputSelect--
				}
			} else {
				if m.selectedAI > 0 {
					m.selectedAI--
				}
			}
		case tea.KeyDown:
			if m.showCreate {
				if m.createInputSelect < 2 {
					m.createInputSelect++
				}
			} else {
				if m.selectedAI < consts.AiCount-1 {
					m.selectedAI++
				}
			}
		case tea.KeyTab:
			if m.showCreate {
				m.createInputSelect = (m.createInputSelect + 1) % 3
			}
		case tea.KeyEnter:
			if m.showCreate {
				// 处理创建AI链接的逻辑
				returnErr := m.AddLinkAPI()
				if !returnErr {
					// 重置创建对话框状态
					m.showCreate = false
					m.createModel.AiLink.Reset()
					m.createModel.AiKey.Reset()
					m.createModel.AiDescription.Reset()
				}
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
						m.createModel.AiLink.Focus()
						m.createModel.AiLink, cmd = m.createModel.AiLink.Update(msg)
					case 1:
						m.createModel.AiKey.Focus()
						m.createModel.AiKey, cmd = m.createModel.AiKey.Update(msg)
					case 2:
						m.createModel.AiDescription.Focus()
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

func NewTui(db *gorm.DB) *Tui {
	// 初始化viewport
	vp := viewport.New(80, 20)
	vp.MouseWheelEnabled = true

	// 创建输入框
	aiLinkInput := textinput.New()
	aiLinkInput.Prompt = ""
	aiLinkInput.Width = 40

	aiKeyInput := textinput.New()
	aiKeyInput.Prompt = ""
	aiKeyInput.Width = 40

	aiDescriptionInput := textinput.New()
	aiDescriptionInput.Prompt = ""
	aiDescriptionInput.Width = 40

	return &Tui{
		selectedAI:        0,
		createInputSelect: 0,
		viewport:          vp,
		db:                db,
		showHelp:          false,
		showCreate:        false,
		createModel: &CreateModel{
			AiLink:        aiLinkInput,
			AiKey:         aiKeyInput,
			AiDescription: aiDescriptionInput,
		},
	}
}
