package main

import "fmt"

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
