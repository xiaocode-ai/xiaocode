package index

import (
	"context"
	"time"
)

var (
	displayEnterContentContext context.Context    // 显示请输入内容「上下文」
	displayEnterContentCancel  context.CancelFunc // 显示请输入内容「取消函数」
)

// DisplayEnterContent
// 显示请输入内容
//
// 用作在聊天框中显示请输入内容；用户在没有输入内容时，按下回车键，会显示请输入内容
func (m *Tui) DisplayEnterContent() {
	// 如果存在定时器上下文，清除现有定时器
	if displayEnterContentContext != nil && displayEnterContentCancel != nil {
		displayEnterContentCancel()
	}

	// 重新创建新的上下文和取消函数
	displayEnterContentContext, displayEnterContentCancel = context.WithCancel(context.Background())

	go func(ctx context.Context) {
		m.custom.plzInputContent = true

		// 使用 select 监听取消信号和定时器
		select {
		case <-ctx.Done():
			// 收到取消信号，立即退出
			m.custom.plzInputContent = false
			return
		case <-time.After(3 * time.Second):
			// 正常超时，执行完成逻辑
			m.custom.plzInputContent = false
		}

		// 定时器完成后清理上下文
		displayEnterContentContext = nil
		displayEnterContentCancel = nil
	}(displayEnterContentContext)
}
