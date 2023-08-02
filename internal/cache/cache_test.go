package cache

import (
	"testing"

	"github.com/KarolosLykos/hackertea/internal/item"
)

func TestCache_Get(t *testing.T) {
	c := New()

	c.Set(1, &item.Item{ID: 1})

	v, ok := c.Get(1)
	if !ok {
		t.Errorf("should find the key")
	}

	if v.ID != 1 {
		t.Errorf("the value should be 1")
	}
}

func TestCache_Set(t *testing.T) {
	c := New()

	c.Set(1, &item.Item{ID: 1})

	_, ok := c.Get(1)
	if !ok {
		t.Errorf("should be present")
	}

	_, ok = c.Get(2)
	if ok {
		t.Errorf("should not be present")
	}
}
