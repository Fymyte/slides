// Package styles implements the theming logic for slides
package styles

import (
	_ "embed"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
	"github.com/charmbracelet/glamour/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const (
	salmon = lipgloss.Color("#E8B4BC")
)

var (
	// Author is the style for the author text in the bottom-left corner of the
	// presentation.
	Author = lipgloss.NewStyle().Foreground(salmon).Align(lipgloss.Left).MarginLeft(2)
	// Date is the style for the date text in the bottom-left corner of the
	// presentation.
	Date = lipgloss.NewStyle().Faint(true).Align(lipgloss.Left).Margin(0, 1)
	// Page is the style for the pagination progress information text in the
	// bottom-right corner of the presentation.
	Page = lipgloss.NewStyle().Foreground(salmon).Align(lipgloss.Right).MarginRight(3)
	// Slide is the style for the slide.
	Slide = lipgloss.NewStyle().Padding(1)
	// Status is the style for the status bar at the bottom of the
	// presentation.
	Status = lipgloss.NewStyle().Padding(1)
	// Search is the style for the search input at the bottom-left corner of
	// the screen when searching is active.
	Search = lipgloss.NewStyle().Faint(true).Align(lipgloss.Left).MarginLeft(2)
)

var (
	// DefaultTheme is the default theme for the presentation.
	//go:embed theme.json
	DefaultTheme []byte
)

// JoinHorizontal joins two strings horizontally and fills the space in-between.
func JoinHorizontal(left, right string, width int) string {
	w := width - lipgloss.Width(right)
	return lipgloss.PlaceHorizontal(w, lipgloss.Left, left) + right
}

// JoinVertical joins two strings vertically and fills the space in-between.
func JoinVertical(top, bottom string, height int) string {
	h := height - lipgloss.Height(bottom)
	return lipgloss.PlaceVertical(h, lipgloss.Top, top) + bottom
}

// SelectTheme picks a glamour style config based
// on the theme provided in the markdown header
func SelectTheme(theme string) glamour.TermRendererOption {
	var styleConfig ansi.StyleConfig
	switch theme {
	case styles.AutoStyle:
		return glamour.WithAutoStyle()
	case styles.AsciiStyle:
		styleConfig = styles.ASCIIStyleConfig
	case styles.DarkStyle:
		styleConfig = styles.DarkStyleConfig
	case styles.LightStyle:
		styleConfig = styles.LightStyleConfig
	case styles.PinkStyle:
		styleConfig = styles.PinkStyleConfig
	case styles.NoTTYStyle:
		styleConfig = styles.NoTTYStyleConfig
	case styles.DraculaStyle:
		styleConfig = styles.DraculaStyleConfig
	case styles.TokyoNightStyle:
		styleConfig = styles.DraculaStyleConfig
	default:
		var themeReader io.Reader
		var err error
		if strings.HasPrefix(theme, "http") {
			var resp *http.Response
			resp, err = http.Get(theme)
			if err != nil {
				return getDefaultTheme()
			}
			defer resp.Body.Close()
			themeReader = resp.Body
		} else {
			file, err := os.Open(theme)
			if err != nil {
				return getDefaultTheme()
			}
			defer file.Close()
			themeReader = file
		}
		bytes, err := io.ReadAll(themeReader)
		if err == nil {
			return glamour.WithStylesFromJSONBytes(bytes)
		}
		// Should log a warning so the user knows we failed to read their theme file
		return getDefaultTheme()
	}

	return glamour.WithStyles(styleConfig)
}

func getDefaultTheme() glamour.TermRendererOption {
	if termenv.EnvNoColor() {
		return glamour.WithStandardStyle("notty")
	}

	if !termenv.HasDarkBackground() {
		return glamour.WithStandardStyle("light")
	}

	return glamour.WithStylesFromJSONBytes(DefaultTheme)
}
