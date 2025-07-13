package ai

import (
	"fmt"
	"github.com/xiaocode-ai/xiaocode/internal/models/entity"
	"github.com/xiaocode-ai/xiaocode/pkg/xerr"
	"github.com/xiaocode-ai/xiaocode/pkg/xlog"
)

// AddLinkAPI
// 添加链接 API
//
// 用于操作数据库添加基本的链接 API 操作
// 如果出现错误，则返回 true 并设置错误信息
func (m *Tui) AddLinkAPI() bool {
	if m.createModel.AiLink.Value() == "" {
		xlog.Logger(xerr.XLevelNotice, xerr.XTagVerify, xerr.XFAIL, "链接不能为空")
		m.showCreateErr = "链接不能为空"
		return true
	}
	if m.createModel.AiKey.Value() == "" {
		xlog.Logger(xerr.XLevelNotice, xerr.XTagVerify, xerr.XFAIL, "密钥不能为空")
		m.showCreateErr = "密钥不能为空"
		return true
	}
	if m.createModel.AiDescription.Value() == "" {
		xlog.Logger(xerr.XLevelNotice, xerr.XTagVerify, xerr.XFAIL, "描述不能为空")
		m.showCreateErr = "描述不能为空"
		return true
	}

	tx := m.db.Create(&entity.AiApiEntity{
		Link:        m.createModel.AiLink.Value(),
		ApiKey:      m.createModel.AiKey.Value(),
		Description: m.createModel.AiDescription.Value(),
		Active:      true,
	})
	if tx.Error != nil {
		xlog.Logger(xerr.XLevelError, xerr.XTagDatabase, xerr.XFAIL, "添加链接 API 失败: "+tx.Error.Error())
		m.showCreateErr = "添加链接 API 失败: " + tx.Error.Error()
		return true
	}

	xlog.Logger(
		xerr.XLevelNotice, xerr.XTagDatabase, xerr.XSUCCESS,
		fmt.Sprintf("添加链接 API 成功: \n\t地址：%s\n\t密钥：%s\n\t描述：%s",
			m.createModel.AiLink.Value(),
			m.createModel.AiKey.Value(),
			m.createModel.AiDescription.Value(),
		),
	)
	m.showCreateErr = ""
	return false
}
