package ui

import (
	"fmt"
	"os"

	"github.com/LintaoAmons/undercontrol/src/domain/account"
	tea "github.com/charmbracelet/bubbletea"
)

type tuiModel struct {
	Accounts []account.Account
}

func (m tuiModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m tuiModel) View() string {
	s := "Init successed"

	return s
}

func InitTui() {
    p := tea.NewProgram(tuiModel{}, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

