package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/xiaocode-ai/xiaocode/internal/tui"
	"github.com/xiaocode-ai/xiaocode/internal/tui/index"
)

func main() {
	keyboard := index.NewKeyboard()

	// 创建一个初始的 Model 实例
	tui := tea.NewProgram(
		tui.New(keyboard),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	// 运行程序
	if _, err := tui.Run(); err != nil {
		panic(err)
	}
}
