package utils

import (
	"context"
	"errors"
	"runtime"
	"testing"

	"github.com/charmbracelet/bubbles/list"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/KarolosLykos/hackertea/internal/item"
	mock_hn "github.com/KarolosLykos/hackertea/internal/mock/hn"
)

func TestUtils_Max(t *testing.T) {
	assert.Equal(t, 5, Max(2, 5))
	assert.Equal(t, 5, Max(5, 2))
	assert.Equal(t, 0, Max(0, 0))
}

func TestUtils_Min(t *testing.T) {
	assert.Equal(t, 2, Min(2, 5))
	assert.Equal(t, 2, Min(5, 2))
	assert.Equal(t, 0, Min(0, 0))
}

func TestUtils_Open(t *testing.T) {
	testCases := []struct {
		name          string
		url           string
		runtimeOS     string
		expectedError error
	}{
		{
			name:          "Empty URL",
			url:           "",
			runtimeOS:     runtime.GOOS,
			expectedError: nil,
		},
		{
			name:          "Unsupported platform",
			url:           "https://example.com",
			runtimeOS:     "unknown",
			expectedError: errors.New("unsupported platform"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Open(tc.url, tc.runtimeOS)

			if (err == nil && tc.expectedError != nil) ||
				(err != nil && tc.expectedError == nil) ||
				(err != nil && err.Error() != tc.expectedError.Error()) {
				t.Errorf("Expected error '%v', but got '%v'", tc.expectedError, err)
			}
		})
	}
}

func TestUtils_FetchStories(t *testing.T) {
	tt := []struct {
		name          string
		workers       int
		ids           [][]int
		tabID         int
		start, end    int
		hnStub        func(hn *mock_hn.MockService)
		expectedItems []list.Item
		expectedLen   int
	}{
		{
			name:    "fetch stories",
			workers: 3,
			ids:     [][]int{{1, 2, 3}},
			tabID:   0, start: 0, end: 3,
			hnStub: func(hn *mock_hn.MockService) {
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 1}, nil)
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 2}, nil)
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 3}, nil)
			},
			expectedItems: []list.Item{&item.Item{ID: 1}, &item.Item{ID: 2}, &item.Item{ID: 3}},
			expectedLen:   3,
		},
		{
			name:    "fetch stories with less workers",
			workers: 1,
			ids:     [][]int{{1, 2, 3}},
			tabID:   0, start: 0, end: 3,
			hnStub: func(hn *mock_hn.MockService) {
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 1}, nil)
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 2}, nil)
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 3}, nil)
			},
			expectedItems: []list.Item{&item.Item{ID: 1}, &item.Item{ID: 2}, &item.Item{ID: 3}},
			expectedLen:   3,
		},
		{
			name:    "fetch stories with unnecessary workers",
			workers: 3,
			ids:     [][]int{{1}},
			tabID:   0, start: 0, end: 1,
			hnStub: func(hn *mock_hn.MockService) {
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 1}, nil)
			},
			expectedItems: []list.Item{&item.Item{ID: 1}},
			expectedLen:   1,
		},
		{
			name:    "fetch stories with error while getting item",
			workers: 3,
			ids:     [][]int{{1, 2, 3}},
			tabID:   0, start: 0, end: 3,
			hnStub: func(hn *mock_hn.MockService) {
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 1}, nil)
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(nil, errors.New("error getting item"))
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&item.Item{ID: 3}, nil)
			},
			expectedItems: []list.Item{
				&item.Item{ID: 1},
				&item.Item{Titl: "Could not get item (error getting item)"},
				&item.Item{ID: 3},
			},
			expectedLen: 3,
		},
		{
			name:    "wrong tabID",
			workers: 3,
			ids:     [][]int{{1}},
			tabID:   1, start: 0, end: 1,
			hnStub: func(hn *mock_hn.MockService) {
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedItems: []list.Item{},
			expectedLen:   0,
		},
		{
			name:    "wrong start - end",
			workers: 3,
			ids:     [][]int{{1}},
			tabID:   0, start: 1, end: 2,
			hnStub: func(hn *mock_hn.MockService) {
				hn.EXPECT().GetItem(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedItems: []list.Item{},
			expectedLen:   0,
		},
	}

	for _, tc := range tt {
		// call the GetItems method and check the result
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Set up mock client response
			mockHN := mock_hn.NewMockService(ctrl)
			tc.hnStub(mockHN)

			items := FetchStories(
				context.Background(),
				mockHN,
				tc.ids,
				tc.workers,
				tc.tabID,
				tc.start, tc.end,
			)

			assert.ElementsMatch(t, items, tc.expectedItems)
			assert.Len(t, items, tc.expectedLen)
		})
	}
}
