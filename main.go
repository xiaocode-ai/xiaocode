package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/xiaocode-ai/xiaocode/internal/app/setup"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	indexTui "github.com/xiaocode-ai/xiaocode/internal/tui/index"
	logTui "github.com/xiaocode-ai/xiaocode/internal/tui/log"
	waitTui "github.com/xiaocode-ai/xiaocode/internal/tui/wait"
	"github.com/xiaocode-ai/xiaocode/pkg/xerr"
	"github.com/xiaocode-ai/xiaocode/pkg/xlog"
)

func main() {
	xlog.Logger(xerr.XLevelDebug, xerr.XTagSetup, xerr.XSUCCESS, "系统初始化中......")

	// 等待页面
	wTui := tea.NewProgram(
		waitTui.NewTui(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	// 初始化配置
	su := setup.New(wTui)
	go func() {
		su.CheckAndCreateSystemProfile()
		su.CheckAndCreateProjectProfile()
		su.ConnectSystemDatabase()
		su.ConnectProjectDatabase()
		su.SystemDatabaseDataPrepare()
		su.Final()
	}()

	if _, err := wTui.Run(); err != nil {
		panic(err)
	}
	if consts.GlobalWaitQuit {
		return
	}

	// 当前 TUI 页面
	var tuiPage = map[string]tea.Model{
		consts.TuiMain: indexTui.NewTui(),
		consts.TuiLog:  logTui.NewTui(),
	}
	for {
		if consts.SystemTuiPage == consts.TuiNil {
			break
		}
		xlog.Logger(xerr.XLevelDebug, xerr.XTagTui, xerr.XSUCCESS, fmt.Sprintf("当前 TUI 页面: %s", consts.SystemTuiPage))
		tuiPage := tea.NewProgram(
			tuiPage[consts.SystemTuiPage],
			tea.WithAltScreen(),
			tea.WithMouseCellMotion(),
		)
		if _, err := tuiPage.Run(); err != nil {
			panic(err)
		}
	}
}
