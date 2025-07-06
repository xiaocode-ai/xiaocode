package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/xiaocode-ai/xiaocode/internal/tui/index"
)

var currentPage = "index"

func New(keyboard *index.CustomKeyboard) tea.Model {
	switch currentPage {
	case "index":
		return index.NewTui(keyboard)
	default:
		return index.NewTui(keyboard)
	}
}
