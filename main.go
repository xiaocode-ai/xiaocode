package main

import (
	tea "github.com/charmbracelet/bubbletea"

	_ "github.com/xiaocode-ai/xiaocode/internal/packed"

	"github.com/xiaocode-ai/xiaocode/internal/app/setup"
	"github.com/xiaocode-ai/xiaocode/internal/tui"
	"github.com/xiaocode-ai/xiaocode/internal/tui/index"
)

func main() {
	// 初始化配置
	su := setup.New()
	su.CheckAndCreateSystemProfile()

	// TUI 渲染部分
	keyboard := index.NewKeyboard()
	tui := tea.NewProgram(
		tui.New(keyboard),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := tui.Run(); err != nil {
		panic(err)
	}
}
