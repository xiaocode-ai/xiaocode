package log

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	"github.com/xiaocode-ai/xiaocode/pkg/xerr"
	"github.com/xiaocode-ai/xiaocode/pkg/xlog"
	"github.com/xiaocode-ai/xiaocode/pkg/xtui"
	"strings"
)

// UiLogger
// 日志界面
//
// 返回日志界面的内容
func (m *Tui) UiLogger() string {
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
		lipgloss.NewStyle().Background(lipgloss.Color(consts.ColorDarkBrown)).Padding(0, 2).Render(
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorText)).Render("日志"),
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

func (m *Tui) content(navBar string) string {
	// 计算宽高
	logContentHeight := m.height - lipgloss.Height(navBar) - 4
	logContentWidth := m.width - 4

	leftContentWidth := 71
	rightContentWidth := logContentWidth - leftContentWidth

	// 渲染日志列表
	logList, selectedLog := m.logList(logContentHeight, leftContentWidth)
	logContent := m.logContent(selectedLog, logContentHeight, rightContentWidth)

	// 渲染内容
	content := lipgloss.NewStyle().Width(logContentWidth).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			logList,
			logContent,
		),
	)

	// 返回渲染后的内容
	return content
}

// logList
// 渲染日志列表
//
// 根据选中的索引返回已渲染的日志列表字符串和选中的日志内容字符串
func (m *Tui) logList(height, width int) (string, string) {
	var lightStyle []string
	var selectedLog string
	var reverseLog []*xerr.Log

	// 反转日志列表，最新的日志在最上面
	if len(xlog.CustomLogs) > 0 {
		reverseLog = make([]*xerr.Log, len(xlog.CustomLogs))
		for i := len(xlog.CustomLogs) - 1; i >= 0; i-- {
			reverseLog[len(xlog.CustomLogs)-1-i] = xlog.CustomLogs[i]
		}
	} else {
		m.viewport.SetContent(lipgloss.NewStyle().Render("暂无日志"))
		return lipgloss.NewStyle().Render("暂无日志"), ""
	}

	for i, singleLog := range reverseLog {
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color(consts.ColorText)).
			Padding(0, 1)
		if i == m.logSelected {
			style = style.Background(lipgloss.Color(consts.ColorDarkPrimary)).
				Bold(true)
			selectedLog = singleLog.Message
		}

		// 对不同日志等级区分颜色
		var levelLight = lipgloss.NewStyle().Width(10)
		if i == m.logSelected {
			levelLight = levelLight.Bold(true).Background(lipgloss.Color(consts.ColorDarkPrimary))
		}
		switch singleLog.Level {
		case xerr.XLevelDebug:
			levelLight = levelLight.Foreground(lipgloss.Color(consts.ColorLightGreen))
		case xerr.XLevelInfo:
			levelLight = levelLight.Foreground(lipgloss.Color(consts.ColorLightBlue))
		case xerr.XLevelNotice:
			levelLight = levelLight.Foreground(lipgloss.Color(consts.ColorLightCyan))
		case xerr.XLevelWarn:
			levelLight = levelLight.Foreground(lipgloss.Color(consts.ColorLightYellow))
		case xerr.XLevelError:
			levelLight = levelLight.Foreground(lipgloss.Color(consts.ColorLightRed))
		case xerr.XLevelPanic:
			levelLight = levelLight.Foreground(lipgloss.Color(consts.ColorLightMagenta))
		}
		tableRender := lipgloss.JoinHorizontal(
			lipgloss.Left,
			style.Width(25).Render(singleLog.Time.Format("2006-01-02 15:04:05")),
			levelLight.Render(singleLog.Level),
			style.Width(15).Render(singleLog.Tag),
			style.Width(15).Render(singleLog.Status),
		)
		lightStyle = append(lightStyle, tableRender)
	}

	// 设置viewport的内容
	logContent := lipgloss.JoinVertical(lipgloss.Left, lightStyle...)
	m.viewport.SetContent(logContent)

	return lipgloss.NewStyle().
		Width(width-2).
		Border(xtui.RightConnectNormalTable()).
		Height(height).
		Padding(1, 2).
		Render(
			m.viewport.View(),
		), selectedLog
}

// logContent
// 渲染选中的日志内容
//
// 返回渲染后的日志内容
func (m *Tui) logContent(selectedLog string, height, width int) string {
	if selectedLog == "" {
		return lipgloss.NewStyle().Render("请选择日志查看详情")
	}
	// 渲染选中的日志内容
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, true, true, false).
		Padding(1, 2).
		Width(width - 1).
		Height(height).
		Render(selectedLog)
}
