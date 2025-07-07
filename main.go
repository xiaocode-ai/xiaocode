package main

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/xiaocode-ai/xiaocode/internal/app/setup"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	indexTui "github.com/xiaocode-ai/xiaocode/internal/tui/index"
	logTui "github.com/xiaocode-ai/xiaocode/internal/tui/log"
)

func main() {
	// 初始化配置
	su := setup.New()
	su.CheckAndCreateSystemProfile()

	// 当前 TUI 页面
	var tuiPage = map[string]tea.Model{
		consts.TuiMain: indexTui.NewTui(),
		consts.TuiLog:  logTui.NewTui(),
	}

	for {
		if consts.SystemTuiPage == consts.TuiNil {
			break
		}

		program := tea.NewProgram(
			tuiPage[consts.SystemTuiPage],
			tea.WithAltScreen(),
			tea.WithMouseCellMotion(),
		)

		if _, err := program.Run(); err != nil {
			panic(err)
		}
	}

}
