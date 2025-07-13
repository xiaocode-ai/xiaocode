package xerr

type XTag = string

const (
	XTagSetup    XTag = "Setup"    // XTagSetup 代表系统初始化相关的日志标签
	XTagTui      XTag = "Tui"      // XTagTui 代表 TUI 相关的日志标签
	XTagVerify   XTag = "Verify"   // XTagAi 代表 AI 相关的日志标签
	XTagDatabase XTag = "Database" // XTagDatabase 代表数据库相关的日志标签
)
