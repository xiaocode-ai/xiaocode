package xlog

import (
	"github.com/xiaocode-ai/xiaocode/pkg/xerr"
	"time"
)

// Logger
// 日志记录器
//
// 用于记录日志信息，将日志信息添加到 CustomLogs 中
func Logger(level xerr.XLevel, tag xerr.XTag, status xerr.XStatus, message string) {
	CustomLogs = append(CustomLogs, &xerr.Log{
		Time:    time.Now(),
		Status:  status,
		Level:   level,
		Tag:     tag,
		Message: message,
	})
}
