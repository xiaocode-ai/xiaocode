package ai

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

var (
	inputRender = lipgloss.NewStyle().Width(40).Align(lipgloss.Left).Render
	errFooter   string
)

func (m *Tui) UiShowCreate() string {
	// 计算宽度
	calcWidth := 60
	if m.width < calcWidth {
		calcWidth = m.width - 4
		if calcWidth < 20 {
			calcWidth = 20
		}
	}

	// 构建头部和尾部
	header := lipgloss.NewStyle().
		Width(calcWidth-4).
		Padding(1, 2).
		Align(lipgloss.Center).
		Render("创建 OpenAI 配置组")
	errFooterStyle := lipgloss.NewStyle().
		Width(calcWidth-4).
		Padding(1, 2, 0, 2).
		Align(lipgloss.Center).
		Background(lipgloss.Color(consts.ColorBackground)).
		Foreground(lipgloss.Color(consts.ColorDarkRed)).
		Bold(true)
	if m.showCreateErr != "" {
		errFooter = errFooterStyle.Render(m.showCreateErr)
	} else {
		errFooter = errFooterStyle.Render("")
	}
	footer := lipgloss.NewStyle().
		Width(calcWidth-4).
		Padding(0, 2, 1, 2).
		Align(lipgloss.Center).
		Render("Esc/Ctrl+C 退出创建 | Enter 确认创建")
	// 构建输入表单
	link := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.matchOutput(0, "链接 >"),
		inputRender(m.createModel.AiLink.View()),
	)
	key := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.matchOutput(1, "密钥 >"),
		inputRender(m.createModel.AiKey.View()),
	)
	description := lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.matchOutput(2, "描述 >"),
		inputRender(m.createModel.AiDescription.View()),
	)
	content := lipgloss.NewStyle().
		Width(calcWidth-4).
		Padding(0, 2).
		Render(lipgloss.JoinVertical(
			lipgloss.Center,
			link,
			key,
			description,
		))

	// 组合使用帮助内容
	horizontal := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		content,
		errFooter,
		footer,
	)
	// 组合最后结果
	var renderQuestion = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Background(lipgloss.Color(consts.ColorBackground)).
		Margin(1, 2).
		Render(horizontal)

	return lipgloss.Place(
		m.width, m.height, lipgloss.Center, lipgloss.Center,
		renderQuestion,
	)
}

func (m *Tui) matchOutput(selectedIndex int, format string) string {
	var baseStyle = lipgloss.NewStyle().Width(10).Padding(0, 2).Align(lipgloss.Left)
	if m.createInputSelect == selectedIndex {
		return baseStyle.Bold(true).Foreground(lipgloss.Color(consts.ColorPrimary)).Render(format)
	} else {
		return baseStyle.Render(format)
	}
}
