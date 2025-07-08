package xerr

type XTag = string

const (
	XTagSetup XTag = "Setup" // XTagSetup 代表系统初始化相关的日志标签
	XTagTui   XTag = "Tui"   // XTagTui 代表 TUI 相关的日志标签
)
