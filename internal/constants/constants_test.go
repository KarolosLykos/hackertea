package constants

import (
	"testing"
)

func TestItemType_Title(t *testing.T) {
	items := Items

	if items.TopItems.Title() != "TOP" {
		t.Errorf("wanted Top got %v", items.TopItems.Title())
	}

	if items.BestItems.Title() != "BEST" {
		t.Errorf("wanted Best got %v", items.BestItems.Title())
	}

	if items.NewItems.Title() != "NEW" {
		t.Errorf("wanted New got %v", items.NewItems.Title())
	}

	if items.SingleItem.Title() != "ITEM" {
		t.Errorf("wanted Item got %v", items.SingleItem.Title())
	}
}
