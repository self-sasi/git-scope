package tui

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Adjust table height based on window size (leave room for header, stats, help, legend)
		tableHeight := m.height - 10
		if tableHeight < 5 {
			tableHeight = 5
		}
		m.table.SetHeight(tableHeight)

	case scanCompleteMsg:
		m.repos = msg.repos
		m.state = StateReady
		m.updateTable()
		m.statusMsg = ""
		return m, nil

	case scanErrorMsg:
		m.state = StateError
		m.err = msg.err
		return m, nil

	case openEditorMsg:
		// Open the editor and wait for it to close
		c := exec.Command(m.cfg.Editor, msg.path)
		return m, tea.ExecProcess(c, func(err error) tea.Msg {
			if err != nil {
				return editorClosedMsg{err: err}
			}
			return editorClosedMsg{}
		})

	case editorClosedMsg:
		if msg.err != nil {
			m.statusMsg = "Error: " + msg.err.Error()
		} else {
			m.statusMsg = ""
		}
		// Rescan after editor closes to update status
		return m, scanReposCmd(m.cfg)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			if m.state == StateReady {
				repo := m.GetSelectedRepo()
				if repo != nil {
					m.statusMsg = "Opening " + repo.Name + " in " + m.cfg.Editor + "..."
					return m, func() tea.Msg {
						return openEditorMsg{path: repo.Path}
					}
				}
			}

		case "r":
			// Rescan
			m.state = StateLoading
			m.statusMsg = "Rescanning..."
			return m, scanReposCmd(m.cfg)

		case "s":
			// Cycle through sort modes
			if m.state == StateReady {
				m.sortMode = (m.sortMode + 1) % 4
				m.updateTable()
				m.statusMsg = "Sorted by: " + m.GetSortModeName()
				return m, nil
			}

		case "1":
			if m.state == StateReady {
				m.sortMode = SortByDirty
				m.updateTable()
				m.statusMsg = "Sorted by: Dirty First"
				return m, nil
			}

		case "2":
			if m.state == StateReady {
				m.sortMode = SortByName
				m.updateTable()
				m.statusMsg = "Sorted by: Name"
				return m, nil
			}

		case "3":
			if m.state == StateReady {
				m.sortMode = SortByBranch
				m.updateTable()
				m.statusMsg = "Sorted by: Branch"
				return m, nil
			}

		case "4":
			if m.state == StateReady {
				m.sortMode = SortByLastCommit
				m.updateTable()
				m.statusMsg = "Sorted by: Recent"
				return m, nil
			}

		case "e":
			// Show editor config hint
			if m.state == StateReady {
				m.statusMsg = "Editor: " + m.cfg.Editor + " (change in ~/.config/git-scope/config.yml)"
			}
		}
	}

	// Update the table
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// editorClosedMsg is sent when the editor process closes
type editorClosedMsg struct {
	err error
}
