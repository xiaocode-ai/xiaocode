package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/xiaocode-ai/xiaocode/internal/tui/index"
)

var currentPage = "index"

func New(keyboard *index.Keyboard) tea.Model {
	switch currentPage {
	case "index":
		return index.NewTui(keyboard)
	default:
		return index.NewTui(keyboard)
	}
}
