package index

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

var (
	menuItem = []string{
		"设置",
		"帮助",
		"历史",
		"AI[LLM]",
		"工具[MCP]",
		"日志",
		"退出",
	}
)

// Style 样式表
var (
	selectedMenuIndex = 0
	menuItemStyle     = lipgloss.NewStyle().
				Width(14).
				Align(lipgloss.Center).
				Background(lipgloss.Color(consts.ColorBackground)).
				Foreground(lipgloss.Color(consts.ColorText)).
				Padding(0, 1)
	menuSelectedItemStyle = lipgloss.NewStyle().
				Width(14).
				Align(lipgloss.Center).
				Background(lipgloss.Color(consts.ColorDarkPrimary)).
				Foreground(lipgloss.Color(consts.ColorText)).
				Padding(0, 1)
)

func (m *Tui) UiPushEsc() string {

	menuStyleList := []string{}
	for i, item := range menuItem {
		if i == selectedMenuIndex {
			menuStyleList = append(menuStyleList, menuSelectedItemStyle.Render(item))
		} else {
			menuStyleList = append(menuStyleList, menuItemStyle.Render(item))
		}
	}

	// 菜单头部UiPushEsc
	menuHeader := lipgloss.NewStyle().
		Padding(0, 2, 1, 2).
		Foreground(lipgloss.Color(consts.ColorText)).
		Render(fmt.Sprintf("%s Menu [%s]", consts.GlobalName, consts.GlobalAuthor))

	// 菜单底部UiPushEsc
	menuFooter := lipgloss.NewStyle().
		Padding(1, 2, 0, 2).
		Foreground(lipgloss.Color(consts.ColorText)).
		Render(fmt.Sprintf(
			"%s 关闭 | %s 选择 | %s 上移 | %s 下移",
			strings.ToUpper(tea.KeyEsc.String()),
			strings.ToUpper(tea.KeyEnter.String()),
			strings.ToUpper(tea.KeyUp.String()),
			strings.ToUpper(tea.KeyDown.String()),
		))

	// 构建菜单文字内容
	menuStringInfo := lipgloss.JoinVertical(
		lipgloss.Center,
		menuHeader,
		lipgloss.JoinVertical(lipgloss.Center, menuStyleList...),
		menuFooter,
	)

	// 创建 ESC 覆盖层，设置明确的宽高
	escOverlay := lipgloss.NewStyle().
		Width(50).
		Background(lipgloss.Color(consts.ColorBackground)).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(consts.ColorDivider)).
		Padding(1, 2).
		Render(menuStringInfo)

	// 将 ESC 菜单放置在半透明背景的中心
	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		escOverlay,
	)
}
