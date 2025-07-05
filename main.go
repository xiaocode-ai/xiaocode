package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/xiaocode-ai/xiaocode/internal/tui"
)

func main() {
	custom := tui.GetCustom()

	// 创建一个初始的 Model 实例
	tui := tea.NewProgram(
		tui.New(custom),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	// 模拟加载完成
	go func() {
		time.Sleep(time.Second * 1)
		custom.SetLoading(false)
	}()

	// 运行程序
	if _, err := tui.Run(); err != nil {
		panic(err)
	}
}
