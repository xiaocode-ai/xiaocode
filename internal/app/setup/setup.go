package setup

import "os"

type Setup struct {
	systemConfigDir string
}

func New() *Setup {
	return &Setup{
		systemConfigDir: os.Getenv("HOME") + "/.xiaocode",
	}
}