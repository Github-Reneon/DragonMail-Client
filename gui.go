package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choice   []string
	cursor   int
	selected map[int]struct{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel() model {
	return model{
		choice:   []string{"First Option", "Second Option", "Third Option"},
		selected: make(map[int]struct{}),
	}
}

func (m model) View() string {
	s := "Options man\n"

	for i, choice := range m.choice {
		cursor := " "
		if m.cursor == i {
			cursor = ">" // cursor
		}

		checked := " "
		// Check to see if this choice is selected
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		// Render
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a keypress
	case tea.KeyMsg:
		// What's the keypress
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choice)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}
