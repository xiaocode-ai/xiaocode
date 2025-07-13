package ai

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
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
	footer := lipgloss.NewStyle().
		Width(calcWidth-4).
		Padding(1, 2).
		Align(lipgloss.Center).
		Render("Esc/Ctrl+C 退出创建 | Enter 确认创建")
	// 构建输入表单
	link := lipgloss.JoinHorizontal(
		lipgloss.Left,
		"链接: ",
		m.createModel.AiLink.View(),
	)
	key := lipgloss.JoinHorizontal(
		lipgloss.Left,
		"链接: ",
		m.createModel.AiKey.View(),
	)
	description := lipgloss.JoinHorizontal(
		lipgloss.Left,
		"链接: ",
		m.createModel.AiDescription.View(),
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
