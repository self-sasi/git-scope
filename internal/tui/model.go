package tui

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Bharath-code/git-scope/internal/config"
	"github.com/Bharath-code/git-scope/internal/model"
)

// State represents the current UI state
type State int

const (
	StateLoading State = iota
	StateReady
	StateError
)

// Model is the Bubbletea model for the TUI
type Model struct {
	cfg       *config.Config
	table     table.Model
	repos     []model.Repo
	state     State
	err       error
	statusMsg string
	width     int
	height    int
}

// NewModel creates a new TUI model
func NewModel(cfg *config.Config) Model {
	columns := []table.Column{
		{Title: "⬤", Width: 2},           // Status indicator
		{Title: "Repository", Width: 20},
		{Title: "Path", Width: 30},
		{Title: "Branch", Width: 15},
		{Title: "Staged", Width: 6},
		{Title: "Modified", Width: 8},
		{Title: "Untracked", Width: 9},
		{Title: "Last Commit", Width: 16},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(12),
	)

	// Apply modern table styles
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#7C3AED")).
		BorderBottom(true).
		Bold(true).
		Foreground(lipgloss.Color("#F9FAFB")).
		Background(lipgloss.Color("#374151"))
	
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#7C3AED")).
		Bold(true)
	
	s.Cell = s.Cell.
		Foreground(lipgloss.Color("#F9FAFB"))
		
	t.SetStyles(s)

	return Model{
		cfg:   cfg,
		table: t,
		state: StateLoading,
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return scanReposCmd(m.cfg)
}

// reposToRows converts repos to table rows with status indicators
func reposToRows(repos []model.Repo) []table.Row {
	// Sort by dirty first, then by name
	sorted := make([]model.Repo, len(repos))
	copy(sorted, repos)
	sort.Slice(sorted, func(i, j int) bool {
		// Dirty repos first
		if sorted[i].Status.IsDirty != sorted[j].Status.IsDirty {
			return sorted[i].Status.IsDirty
		}
		// Then by name
		return sorted[i].Name < sorted[j].Name
	})

	rows := make([]table.Row, 0, len(sorted))
	for _, r := range sorted {
		lastCommit := "N/A"
		if !r.Status.LastCommit.IsZero() {
			lastCommit = r.Status.LastCommit.Format("Jan 02 15:04")
		}

		// Status indicator
		indicator := "○" // Clean
		if r.Status.IsDirty {
			indicator = "●" // Dirty
		}

		rows = append(rows, table.Row{
			indicator,
			truncateString(r.Name, 20),
			truncatePath(r.Path, 30),
			truncateString(r.Status.Branch, 15),
			colorNumber(r.Status.Staged, "#10B981"),      // Green
			colorNumber(r.Status.Unstaged, "#F59E0B"),    // Amber
			colorNumber(r.Status.Untracked, "#9CA3AF"),   // Gray
			lastCommit,
		})
	}
	return rows
}

// truncatePath shortens a path to fit in the given width
func truncatePath(path string, maxLen int) string {
	if len(path) <= maxLen {
		return path
	}
	return "…" + path[len(path)-maxLen+1:]
}

// truncateString shortens a string with ellipsis
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-1] + "…"
}

// colorNumber returns a string representation of a number
func colorNumber(n int, _ string) string {
	if n == 0 {
		return "—"
	}
	return fmt.Sprintf("%d", n)
}
