package dto

type SystemConfigDTO struct {
	System struct {
		LogLevel string `toml:"LOG_LEVEL"` // 日志级别
	}
}
