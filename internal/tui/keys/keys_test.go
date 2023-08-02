package keys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListKeyMap(t *testing.T) {
	listKeys := NewListKeyMap()

	// Test NewListKeyMap
	assert.NotNil(t, listKeys.nextPage)
	assert.NotNil(t, listKeys.previousPage)
	assert.NotNil(t, listKeys.nextTab)
	assert.NotNil(t, listKeys.previousTab)
	assert.NotNil(t, listKeys.fetchNextPage)

	// Test KeyBindings
	bindings := listKeys.KeyBindings()
	assert.Equal(t, 5, len(bindings()))
	assert.Contains(t, bindings(), listKeys.nextPage)
	assert.Contains(t, bindings(), listKeys.previousPage)
	assert.Contains(t, bindings(), listKeys.nextTab)
	assert.Contains(t, bindings(), listKeys.previousTab)
	assert.Contains(t, bindings(), listKeys.fetchNextPage)
}
