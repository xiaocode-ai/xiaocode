package index

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
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

// UiMain
// 主界面
//
// 返回主界面的内容
func (m Tui) UiMain() string {
	// 左侧内容：logo 和版本号
	leftContent := lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorDarkPrimary)).Padding(0, 2).Render(
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorText)).Render("XiaoCode"),
		),
		lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorAccent)).Padding(0, 2).Render(
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorText)).Render("v1.0.0"),
		),
	)

	// 右侧内容：作者信息
	rightContent := lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorDivider)).Padding(0, 2).Render(
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorPrimaryText)).Render("XiaoLFeng"),
	)

	// 计算中间需要的空间来实现 space-between 效果
	leftWidth := lipgloss.Width(leftContent)
	rightWidth := lipgloss.Width(rightContent)
	totalUsedWidth := leftWidth + rightWidth

	// 如果终端宽度足够，计算中间空间
	var middleSpace string
	if m.width > totalUsedWidth {
		spaceWidth := m.width - totalUsedWidth - 4
		if spaceWidth > 0 {
			middleSpace = strings.Repeat(" ", spaceWidth)
		}
	}

	// 顶部内容
	topContent := lipgloss.JoinHorizontal(
		lipgloss.Left,
		leftContent,
		middleSpace,
		rightContent,
	)

	// 底部输入框
	bottomContent := input(m)
	middleContent := middleContent(m, topContent, bottomContent)

	return lipgloss.NewStyle().Padding(1, 2).Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			topContent,
			middleContent,
			bottomContent,
		),
	)
}

// middleContent 生成 TUI 界面的中间内容区域
//
// 参数:
//   - m: Tui 模型实例，包含当前 TUI 的状态信息
//   - topContent: 顶部内容的字符串表示
//   - bottomContent: 底部内容的字符串表示
//
// 返回值:
//   - string: 中间内容区域的渲染结果
//
// 功能描述:
//  1. 如果正在加载，显示加载动画和退出提示
//  2. 计算中间区域的高度，确保界面布局合理
//  3. 根据可用高度生成适当的空白填充
func middleContent(m Tui, topContent string, bottomContent string) string {
	if m.custom.loading {
		return lipgloss.NewStyle().Padding(1, 0).Render(
			fmt.Sprintf("%s 正在加载...「按 %s 退出」", m.spinner.View(), tea.KeyCtrlC),
		)
	}

	topHeight := lipgloss.Height(topContent)
	bottomHeight := lipgloss.Height(bottomContent)
	paddingHeight := 2 // 上下各2行边距
	myselfHeight := 2  // 自己占用的行数

	middleHeight := m.height - topHeight - bottomHeight - paddingHeight - myselfHeight
	middleWidth := m.width - 6
	if middleHeight < 0 {
		middleHeight = 0
	}

	// 创建中间的填充空间
	middleContent := "你好世界"
	return lipgloss.NewStyle().
		Height(middleHeight).
		Width(middleWidth).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(consts.ColorDivider)).
		Render(middleContent)
}

// UiInput
// 输入框
//
// 如果正在加载，则返回空字符串
// 否则返回输入框的内容
func input(m Tui) string {
	if m.custom.loading {
		return ""
	}
	m.textInput.Prompt = ""

	// 前置输入内容
	prefixRender := lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(lipgloss.Color(consts.ColorDarkPrimary)).Render(">>")
	inputWidth := m.width - lipgloss.Width(prefixRender) - 4
	inputRender := lipgloss.NewStyle().Width(inputWidth).Padding(0, 1).Background(lipgloss.Color(consts.ColorDarkBrown)).Render(m.textInput.View())

	return lipgloss.NewStyle().Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			prefixRender,
			inputRender,
		),
	)
}
