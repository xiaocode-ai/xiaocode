package ai

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	"strings"
)

// UiAi
// 代表 TUI 界面中的 AI 功能
//
// 用于 TUI 展示 AI 配置相关页面，用于配置 OpenAPI 链接地址；
// 客制化 AI 的链接地址和API密钥等信息。
func (m *Tui) UiAi() string {
	var (
		navBar  = m.navBar()
		content = m.content(navBar)
	)

	return lipgloss.NewStyle().Padding(1, 2).Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			navBar,
			content,
		),
	)
}

// navBar
// 导航栏
//
// 返回导航栏的内容
func (m *Tui) navBar() string {
	// 左侧内容：logo 和版本号
	leftContent := lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorDarkPrimary)).Padding(0, 2).Render(
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorText)).Render(consts.GlobalName),
		),
		lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorDarkPurple)).Padding(0, 2).Render(
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorText)).Render("AI 配置"),
		),
	)

	// 右侧内容：作者信息
	rightContent := lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorDivider)).Padding(0, 2).Render(
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorPrimaryText)).Render("?(Ctrl+/) 查看帮助"),
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
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		leftContent,
		middleSpace,
		rightContent,
	)
}

func (m *Tui) content(navbar string) string {
	spaceWidth := m.width - 4
	spaceHeight := m.height - lipgloss.Height(navbar) - 2

	newTable := table.New().
		Border(lipgloss.NormalBorder()).
		Headers("描述", "链接地址", "启用").
		Width(spaceWidth).
		Height(spaceHeight).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lipgloss.NewStyle().Foreground(lipgloss.Color(consts.ColorText)).Bold(true).Align(lipgloss.Center)
			default:
				baseStyle := lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Left)
				if m.selectedAI == row {
					return baseStyle.Bold(true).Foreground(lipgloss.Color(consts.ColorAccent))
				}
				return lipgloss.NewStyle().Padding(0, 1)
			}
		})
	newTable.Row("测试", "https://api.xiaocode.com/", "启用")
	newTable.Row("测试", "https://api.xiaocode.com/", "禁用")

	return newTable.String()
}
