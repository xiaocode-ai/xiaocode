package wait

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	"strings"
)

// UiWait
// 生成 UI 等待状态的渲染内容。
//
// 该方法组织并渲染导航栏、中间内容及底部组件的布局，返回完整的待显示字符串。
func (m *Tui) UiWait() string {
	return lipgloss.NewStyle().Padding(1, 2).Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.navBar(),
			m.waitInfo(),
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
		lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorAccent)).Padding(0, 2).Render(
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorText)).Render(consts.GlobalVersion),
		),
	)

	// 右侧内容：作者信息
	rightContent := lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorDivider)).Padding(0, 2).Render(
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorPrimaryText)).Render(consts.GlobalAuthor),
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

func (m *Tui) waitInfo() string {
	waitInfo := lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.BlockBorder(), false, false, false, true).
		BorderForeground(lipgloss.Color(consts.ColorDarkGreen)).
		Render(fmt.Sprintf("%s 等待系统加载，请稍后...", m.spinner.View()))
	return lipgloss.NewStyle().
		Padding(1, 0).
		Render(waitInfo)
}
