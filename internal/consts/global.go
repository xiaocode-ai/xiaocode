package consts

import "github.com/xiaocode-ai/xiaocode/internal/models/dto"

const (
	GlobalName    = "XiaoCode"                     // 软件名称
	GlobalVersion = "v0.0.1-alpha"                 // 软件版本
	GlobalAuthor  = "XiaoLFeng"                    // 软件作者
	GlobalEmail   = "gm@x-lf.cn"                   // 软件作者邮箱
	GlobalGithub  = "https://github.com/XiaoLFeng" // 软件作者 Github
)

var (
	GlobalSystemProfile  *dto.SystemConfigDTO  // 系统配置文件
	GlobalProjectProfile *dto.ProjectConfigDTO // 项目配置文件
	GlobalWaitQuit       = false               // 全局等待退出标志
)
