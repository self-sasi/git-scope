package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Color palette - Modern, vibrant theme
var (
	// Primary colors
	primaryColor   = lipgloss.Color("#7C3AED") // Purple
	secondaryColor = lipgloss.Color("#10B981") // Green
	accentColor    = lipgloss.Color("#F59E0B") // Amber
	dangerColor    = lipgloss.Color("#EF4444") // Red
	
	// Neutral colors
	bgColor      = lipgloss.Color("#1F2937")  // Dark gray
	surfaceColor = lipgloss.Color("#374151")  // Medium gray
	textColor    = lipgloss.Color("#F9FAFB")  // White
	mutedColor   = lipgloss.Color("#9CA3AF")  // Gray
	
	// Status colors
	cleanColor = lipgloss.Color("#10B981") // Green
	dirtyColor = lipgloss.Color("#F59E0B") // Amber
	errorColor = lipgloss.Color("#EF4444") // Red
)

// Application styles
var (
	// App container
	appStyle = lipgloss.NewStyle().
		Padding(1, 2)

	// Header / Title
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(primaryColor).
		Padding(0, 2).
		MarginBottom(1)
	
	// Logo ASCII art style
	logoStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true)

	// Subtitle with stats
	subtitleStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		MarginBottom(1)
	
	// Stats badges
	statsBadgeStyle = lipgloss.NewStyle().
		Foreground(textColor).
		Background(surfaceColor).
		Padding(0, 1).
		MarginRight(1)
	
	dirtyBadgeStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(dirtyColor).
		Padding(0, 1).
		Bold(true)
	
	cleanBadgeStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(cleanColor).
		Padding(0, 1).
		Bold(true)

	// Table styles
	tableContainerStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(surfaceColor).
		Padding(0, 1)

	// Help footer
	helpStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		MarginTop(1)
	
	helpKeyStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true)
	
	helpDescStyle = lipgloss.NewStyle().
		Foreground(mutedColor)

	// Status message
	statusStyle = lipgloss.NewStyle().
		Foreground(accentColor).
		MarginTop(1)

	// Error styling
	errorTitleStyle = lipgloss.NewStyle().
		Foreground(errorColor).
		Bold(true)
	
	errorBoxStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(errorColor).
		Padding(1, 2).
		MarginTop(1)

	// Loading styling
	loadingStyle = lipgloss.NewStyle().
		Foreground(secondaryColor).
		Bold(true)
	
	loadingSpinnerStyle = lipgloss.NewStyle().
		Foreground(primaryColor)
	
	// Scanning paths list
	pathStyle = lipgloss.NewStyle().
		Foreground(textColor).
		PaddingLeft(2)
	
	pathBulletStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true)

	// Repo row indicators
	dirtyIndicator = lipgloss.NewStyle().
		Foreground(dirtyColor).
		Bold(true).
		Render("●")
	
	cleanIndicator = lipgloss.NewStyle().
		Foreground(cleanColor).
		Render("○")
)

// Help item creates a styled help key-description pair
func helpItem(key, desc string) string {
	return helpKeyStyle.Render(key) + helpDescStyle.Render(" "+desc)
}

// Logo returns the ASCII art logo
func logo() string {
	return logoStyle.Render(`
  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
  ┃   ╔═╗╦╔╦╗  ╔═╗╔═╗╔═╗╔═╗╔═╗   ┃
  ┃   ║ ╦║ ║───╚═╗║  ║ ║╠═╝║╣    ┃
  ┃   ╚═╝╩ ╩   ╚═╝╚═╝╚═╝╩  ╚═╝   ┃
  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛`)
}

// Simpler logo for compact mode
func compactLogo() string {
	return titleStyle.Render(" 🔍 git-scope ")
}
