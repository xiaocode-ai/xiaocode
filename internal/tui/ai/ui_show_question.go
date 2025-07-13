package ai

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

// UiShowQuestion
// 生成 UI 日志使用帮助的渲染内容。
//
// 该方法组织并渲染头部、左右两侧的帮助内容，返回完整的待显示字符串。
func (m *Tui) UiShowQuestion() string {
	// 计算宽度
	calcWidth := 60
	if m.width < calcWidth {
		calcWidth = m.width - 4
		if calcWidth < 20 {
			calcWidth = 20
		}
	}

	// 构建头部和帮助内容
	header := lipgloss.NewStyle().
		Width(calcWidth-4).
		Padding(1, 2).
		Align(lipgloss.Center).
		Render(fmt.Sprintf("%s AI 配置使用帮助", consts.GlobalName))
	// 构建帮助内容
	leftContent := lipgloss.NewStyle().
		Width(calcWidth/2-4).
		Padding(0, 2).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				"↑      向上选中",
				"↓      向下选中",
				"PgUp   向上翻页",
				"PgDown 向下翻页",
			),
		)
	rightContent := lipgloss.NewStyle().
		Width(calcWidth/2-4).
		Padding(0, 2).
		Align(lipgloss.Right).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				"N      创建AI组",
				"Esc    退出帮助",
				"Home   首个日志",
				"End    末尾日志",
			),
		)
	var index string
	if calcWidth > 50 {
		// 组合内容
		index = lipgloss.NewStyle().
			Width(calcWidth).
			Align(lipgloss.Center).
			Padding(0, 0, 1, 0).
			Render(
				lipgloss.JoinHorizontal(
					lipgloss.Center,
					leftContent,
					rightContent,
				),
			)
	} else {
		// 如果宽度过小，直接使用垂直放置
		index = lipgloss.JoinVertical(
			lipgloss.Center,
			leftContent,
			rightContent,
		)
	}

	// 组合使用帮助内容
	horizontal := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		index,
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
