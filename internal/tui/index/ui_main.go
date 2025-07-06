package index

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

// UiMain
// 主界面
//
// 返回主界面的内容
func (m Tui) UiMain() string {

	// 组件引入
	navBar := m.navBar()
	bottomContent := m.footer()
	middleContent := m.middleContent(navBar, bottomContent)

	return lipgloss.NewStyle().Padding(1, 2).Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			navBar,
			middleContent,
			bottomContent,
		),
	)
}

// navBar
// 导航栏
//
// 返回导航栏的内容
func (m Tui) navBar() string {
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
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		leftContent,
		middleSpace,
		rightContent,
	)
}

// middleContent
// 中间内容
//
// 返回中间内容的内容
func (m Tui) middleContent(topContent string, bottomContent string) string {
	topHeight := lipgloss.Height(topContent)
	bottomHeight := lipgloss.Height(bottomContent)
	paddingHeight := 2
	myselfHeight := 2

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
		Padding(1, 1).
		Render(middleContent)
}

// footer
// 底部内容
//
// 返回底部内容的内容
func (m Tui) footer() string {
	m.textInput.Prompt = ""

	// 前置输入内容
	prefixRender := lipgloss.NewStyle().Bold(true).Padding(0, 1, 0, 0).Foreground(lipgloss.Color(consts.ColorAccent)).Render(">")
	inputWidth := m.width - lipgloss.Width(prefixRender) - 4
	inputRender := lipgloss.NewStyle().Width(inputWidth).Render(m.textInput.View())

	content := lipgloss.JoinHorizontal(
		lipgloss.Left,
		prefixRender,
		inputRender,
	)

	return lipgloss.NewStyle().
		BorderTop(true).
		BorderTopForeground(lipgloss.Color("#FFFFFF")).
		Render(content)
}
