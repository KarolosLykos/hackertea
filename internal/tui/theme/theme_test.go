package theme

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTheme(t *testing.T) {
	// Create new instance first.
	theme, err := GetTheme()
	require.NoError(t, err)
	require.NotNil(t, theme)

	// Get instance after the first time.
	theme1, err1 := GetTheme()
	require.NoError(t, err1)
	require.NotNil(t, theme1)

	assert.EqualValues(t, theme, theme1)
}

func TestNewTheme(t *testing.T) {
	theme, err := NewTheme()
	assert.NotNil(t, theme)
	assert.NoError(t, err)

	assert.Equal(t, lipgloss.NewStyle().Padding(1, 2, 1, 2), theme.Doc)
}
