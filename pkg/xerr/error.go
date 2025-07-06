package xerr

import "time"

// Log
// 日志结构体
type Log struct {
	Time    time.Time
	Status  XStatus
	Level   XLevel
	Tag     XTag
	Message string
}

type XStatus = string

// XStatus
// 状态
const (
	XSUCCESS XStatus = "SUCCESS" // 成功
	XFAIL    XStatus = "FAIL"    // 失败
)

type XLevel = string

// XLevel
// 日志级别
const (
	XLevelDebug  XLevel = "DEBUG"  // 调试
	XLevelInfo   XLevel = "INFO"   // 信息
	XLevelNotice XLevel = "NOTICE" // 注意
	XLevelWarn   XLevel = "WARN"   // 警告
	XLevelError  XLevel = "ERROR"  // 错误
	XLevelPanic  XLevel = "PANIC"  // 恐慌
)
