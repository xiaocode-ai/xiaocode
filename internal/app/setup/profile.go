package setup

import (
	"os"
	"time"

	"github.com/xiaocode-ai/xiaocode/pkg/xerr"
)

// CheckAndCreateSystemProfile
// 检查并创建系统配置文件夹
//
// 如果系统配置文件夹不存在，则创建系统配置文件夹
func (s *Setup) CheckAndCreateSystemProfile() {
	// 检查系统配置文件夹是否存在
	if _, err := os.Stat(s.systemConfigDir); os.IsNotExist(err) {
		osErr := os.MkdirAll(s.systemConfigDir, 0755)

		if osErr != nil {
			xerr.CutsomLog = append(xerr.CutsomLog, xerr.Log{
				Time:    time.Now(),
				Status:  xerr.XFAIL,
				Level:   xerr.XLevelError,
				Tag:     xerr.XTagSetup,
				Message: osErr.Error(),
			})
		}
	}

	// 检查配置文件是否存在
	configFile := s.systemConfigDir + "/config.toml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		os.Create(configFile)
	}
}
