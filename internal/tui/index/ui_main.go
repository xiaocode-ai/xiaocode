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
func (m *Tui) UiMain() string {

	// 组件引入
	navBar := m.navBar()
	bottomContent := m.footer()
	middleContent := m.middleContent(navBar, bottomContent)

	return lipgloss.NewStyle().Padding(1, 2).Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
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
func (m *Tui) navBar() string {
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
func (m *Tui) middleContent(topContent string, bottomContent string) string {
	topHeight := lipgloss.Height(topContent)
	bottomHeight := lipgloss.Height(bottomContent)
	paddingHeight := 2

	middleHeight := m.height - topHeight - bottomHeight - paddingHeight
	middleWidth := m.width - 6
	if middleHeight < 0 {
		middleHeight = 0
	}

	// 创建 AI 交流对话框
	middleContent := m.renderChatContent(middleHeight, middleWidth)
	return lipgloss.NewStyle().
		Padding(1, 0).
		Height(middleHeight).
		Width(middleWidth).
		Render(middleContent)
}

// renderChatContent
// 渲染聊天内容
//
// 返回聊天内容的内容
func (m *Tui) renderChatContent(maxHeight, maxWidth int) string {
	// 欢迎介绍语
	if len(m.chat.chat) == 0 {
		return lipgloss.NewStyle().
			Border(lipgloss.BlockBorder(), false, false, false, true).
			BorderForeground(lipgloss.Color(consts.ColorDarkOrange)).
			Width(maxWidth).
			Render(
				lipgloss.NewStyle().
					Padding(1, 0, 1, 1).
					Align(lipgloss.Left).
					Foreground(lipgloss.Color(consts.ColorText)).
					Render("你好这里是 XiaoCode 的聊天界面，你可以输入对话内容，也可以按下 ESC 键来查看帮助菜单。\n我可以帮你写代码，也可以帮你写文章，还可以帮你写诗，甚至可以帮你写小说。\n\n你可以输入任何内容，我会尽力回答你。"),
			)
	}

	chatContent := []string{}

	for _, chat := range m.chat.chat {
		if chat.user == "user" {
			// 用户消息：右对齐显示
			messageWidth := maxWidth * 2 / 3 // 用户消息宽度为容器的2/3

			userMessage := lipgloss.NewStyle().
				Border(lipgloss.BlockBorder(), false, true, false, false).
				BorderForeground(lipgloss.Color(consts.ColorPrimary)).
				Width(messageWidth).
				Align(lipgloss.Right).
				MarginBottom(1).
				Render(
					lipgloss.NewStyle().
						PaddingRight(1).
						Foreground(lipgloss.Color(consts.ColorText)).
						Render(chat.content),
				)

			// 使用 Place 将用户消息放置到右侧
			rightAlignedMessage := lipgloss.Place(
				maxWidth, lipgloss.Height(userMessage),
				lipgloss.Right, lipgloss.Top,
				userMessage,
			)

			chatContent = append(chatContent, rightAlignedMessage)
		} else {
			// AI消息：左对齐显示
			messageWidth := maxWidth * 2 / 3 // AI消息宽度为容器的2/3

			aiMessage := lipgloss.NewStyle().
				Border(lipgloss.BlockBorder(), false, false, false, true).
				BorderForeground(lipgloss.Color(consts.ColorAccent)).
				Width(messageWidth).
				Align(lipgloss.Left).
				MarginBottom(1).
				Render(
					lipgloss.NewStyle().
						PaddingLeft(1).
						Foreground(lipgloss.Color(consts.ColorText)).
						Render(chat.content),
				)

			// 使用 Place 将 AI 消息放置到左侧
			leftAlignedMessage := lipgloss.Place(
				maxWidth, lipgloss.Height(aiMessage),
				lipgloss.Left, lipgloss.Top,
				aiMessage,
			)

			chatContent = append(chatContent, leftAlignedMessage)
		}
	}

	return lipgloss.NewStyle().
		Width(maxWidth).
		Render(lipgloss.JoinVertical(lipgloss.Left, chatContent...))
}

// footer
// 底部内容
//
// 返回底部内容的内容
func (m *Tui) footer() string {
	// 前置输入内容
	prefixRender := lipgloss.NewStyle().Bold(true).Padding(0, 1, 0, 0).Foreground(lipgloss.Color(consts.ColorAccent)).Render(">")
	inputWidth := m.width - lipgloss.Width(prefixRender) - 4

	// 设置 textarea 的宽度
	m.textArea.SetWidth(inputWidth)
	inputRender := lipgloss.NewStyle().Width(inputWidth).Render(m.textArea.View())

	content := lipgloss.JoinHorizontal(
		lipgloss.Left,
		prefixRender,
		inputRender,
	)

	inputContent := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(lipgloss.Color(consts.ColorDivider)).
		Render(content)

	if m.custom.plzInputContent {
		plzInputContent := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(consts.ColorText)).
			Background(lipgloss.Color(consts.ColorDarkPurple)).
			Padding(1, 4).
			Align(lipgloss.Center).
			Render("♦︎ 请输入内容 ♦︎")

		return lipgloss.JoinVertical(
			lipgloss.Center,
			plzInputContent,
			inputContent,
		)
	} else {
		return inputContent
	}
}
