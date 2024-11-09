package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel() model {
	return model{
		choice:   []string{"First Option", "Second Option", "Third Option"},
		selected: make(map[int]struct{}),
	}
}
