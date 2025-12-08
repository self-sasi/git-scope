package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View renders the TUI
func (m Model) View() string {
	// Apply app container style
	content := m.renderContent()
	return appStyle.Render(content)
}

func (m Model) renderContent() string {
	var b strings.Builder

	switch m.state {
	case StateLoading:
		b.WriteString(m.renderLoading())
	case StateError:
		b.WriteString(m.renderError())
	case StateReady:
		b.WriteString(m.renderDashboard())
	}

	return b.String()
}

func (m Model) renderLoading() string {
	var b strings.Builder

	// Header
	b.WriteString(compactLogo())
	b.WriteString("  ")
	b.WriteString(loadingStyle.Render("Scanning repositories..."))
	b.WriteString("\n\n")

	// Scanning paths
	b.WriteString(subtitleStyle.Render("Searching for git repos in:"))
	b.WriteString("\n")
	for _, root := range m.cfg.Roots {
		b.WriteString(pathBulletStyle.Render("  → "))
		b.WriteString(pathStyle.Render(root))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	// Help
	b.WriteString(helpStyle.Render("Press " + helpKeyStyle.Render("q") + " to quit"))

	return b.String()
}

func (m Model) renderError() string {
	var b strings.Builder

	// Header with error indicator
	b.WriteString(compactLogo())
	b.WriteString("  ")
	b.WriteString(errorTitleStyle.Render("✗ Error"))
	b.WriteString("\n")

	// Error box
	errContent := ""
	if m.err != nil {
		errContent = m.err.Error()
	} else {
		errContent = "Unknown error occurred"
	}
	b.WriteString(errorBoxStyle.Render(errContent))
	b.WriteString("\n\n")

	// Help
	b.WriteString(helpItem("q", "quit"))
	b.WriteString("  •  ")
	b.WriteString(helpItem("r", "retry"))

	return b.String()
}

func (m Model) renderDashboard() string {
	var b strings.Builder

	// Header with logo
	b.WriteString(compactLogo())
	b.WriteString("\n\n")

	// Stats bar
	b.WriteString(m.renderStats())
	b.WriteString("\n\n")

	// Table
	b.WriteString(m.table.View())
	b.WriteString("\n")

	// Status message if any
	if m.statusMsg != "" {
		b.WriteString(statusStyle.Render("→ " + m.statusMsg))
		b.WriteString("\n")
	}

	// Help footer
	b.WriteString(m.renderHelp())

	return b.String()
}

func (m Model) renderStats() string {
	total := len(m.repos)
	dirty := 0
	clean := 0
	for _, r := range m.repos {
		if r.Status.IsDirty {
			dirty++
		} else {
			clean++
		}
	}

	stats := []string{
		statsBadgeStyle.Render(fmt.Sprintf("📁 %d repos", total)),
	}
	
	if dirty > 0 {
		stats = append(stats, dirtyBadgeStyle.Render(fmt.Sprintf("⚠ %d dirty", dirty)))
	}
	if clean > 0 {
		stats = append(stats, cleanBadgeStyle.Render(fmt.Sprintf("✓ %d clean", clean)))
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, stats...)
}

func (m Model) renderHelp() string {
	var items []string

	items = append(items, helpItem("↑↓", "navigate"))
	items = append(items, helpItem("enter", "open"))
	items = append(items, helpItem("r", "rescan"))
	items = append(items, helpItem("q", "quit"))

	return helpStyle.Render(strings.Join(items, "  •  "))
}
