package setup

import (
	tea "github.com/charmbracelet/bubbletea"
	"gorm.io/gorm"
	"os"
)

type Setup struct {
	SystemConfigDir string
	SystemDB        *gorm.DB
	ProjectDB       *gorm.DB
	WaitTui         *tea.Program
}

func New(tui *tea.Program) *Setup {
	return &Setup{
		WaitTui:         tui,
		SystemConfigDir: os.Getenv("HOME") + "/.xiaocode",
	}
}

func (s *Setup) Final() {
	s.WaitTui.Quit()
}
