package hn

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KarolosLykos/hackertea/internal/constants"
	"github.com/KarolosLykos/hackertea/internal/item"
	"github.com/KarolosLykos/hackertea/internal/mock/cache"
	"github.com/KarolosLykos/hackertea/internal/mock/client"
)

func TestGetSuffix(t *testing.T) {

	tt := []struct {
		name   string
		value  constants.ItemType
		suffix string
		err    error
	}{
		{name: "New items", value: constants.Items.NewItems, suffix: constants.NewSuffix},
		{name: "Top items", value: constants.Items.TopItems, suffix: constants.TopSuffix},
		{name: "Best items", value: constants.Items.BestItems, suffix: constants.BestSuffix},
		{name: "Single item", value: constants.Items.SingleItem, suffix: constants.SingleSuffix},
		{name: "Error", value: "", err: ErrInvalidItemType},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			suffix, err := getSuffix(tc.value)
			if err != nil && tc.err != nil {
				assert.ErrorIs(t, err, ErrInvalidItemType)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.suffix, suffix)
			}
		})
	}
}

func TestGetItems(t *testing.T) {
	tt := []struct {
		name          string
		itemType      constants.ItemType
		clientStub    func(client *mock_client.MockHttpClient)
		expectedItems []int
		expectedError error
	}{
		{name: "wrong item type", itemType: "", clientStub: func(client *mock_client.MockHttpClient) {}, expectedItems: nil, expectedError: ErrInvalidItemType},
		{name: "success case", itemType: constants.Items.TopItems, clientStub: func(client *mock_client.MockHttpClient) {
			client.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return([]byte(`[123, 456, 789]`), nil)
		}, expectedItems: []int{123, 456, 789}, expectedError: nil},
		{name: "invalid json", itemType: constants.Items.TopItems, clientStub: func(client *mock_client.MockHttpClient) {
			client.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return([]byte(`{invalid json}`), nil)
		}, expectedItems: nil, expectedError: errors.New("invalid character")},
		{name: "client error response", itemType: constants.Items.TopItems, clientStub: func(client *mock_client.MockHttpClient) {
			client.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(nil, errors.New("client error"))
		}, expectedItems: nil, expectedError: errors.New("client error")},
	}

	for _, tc := range tt {
		// call the GetItems method and check the result
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Set up mock client response
			mockClient := mock_client.NewMockHttpClient(ctrl)
			tc.clientStub(mockClient)

			hn := New(mockClient, nil)
			items, err := hn.GetItems(context.Background(), tc.itemType)
			if err != nil && tc.expectedError != nil {
				assert.ErrorContains(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedItems, items)
			}
		})
	}
}

func TestHN_GetItem(t *testing.T) {
	testCases := []struct {
		name       string
		id         int
		clientStub func(client *mock_client.MockHttpClient)
		cacheStub  func(cache *mock_cache.MockCache)
		response   []byte
		respErr    error
		expected   *item.Item
		expectErr  bool
	}{
		{
			name: "cache hit",
			id:   123,
			clientStub: func(client *mock_client.MockHttpClient) {
				client.EXPECT().Get(gomock.Any(), gomock.Any()).Times(0)
			},
			cacheStub: func(cache *mock_cache.MockCache) {
				cache.EXPECT().Get(gomock.Any()).Times(1).Return(&item.Item{ID: 123}, true)
			},
			response:  nil,
			expected:  &item.Item{ID: 123},
			expectErr: false,
		},
		{
			name: "cache miss",
			id:   123,
			clientStub: func(client *mock_client.MockHttpClient) {
				client.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return([]byte(`{"id":123}`), nil)
			},
			cacheStub: func(cache *mock_cache.MockCache) {
				cache.EXPECT().Get(gomock.Any()).Times(1).Return(nil, false)
				cache.EXPECT().Set(gomock.Any(), gomock.Any()).Times(1)
			},
			expected:  &item.Item{ID: 123},
			expectErr: false,
		},
		{
			name: "invalid response",
			id:   123,
			clientStub: func(client *mock_client.MockHttpClient) {
				client.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return([]byte(`{"invalid"`), nil)
			},
			cacheStub: func(cache *mock_cache.MockCache) { cache.EXPECT().Get(gomock.Any()).Times(1).Return(nil, false) },
			expected:  nil,
			expectErr: true,
		},
		{
			name: "GET error response",
			id:   123,
			clientStub: func(client *mock_client.MockHttpClient) {
				client.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(nil, errors.New("get error"))
			},
			cacheStub: func(cache *mock_cache.MockCache) { cache.EXPECT().Get(gomock.Any()).Times(1).Return(nil, false) },
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Set up mock client response
			mockCache := mock_cache.NewMockCache(ctrl)
			mockClient := mock_client.NewMockHttpClient(ctrl)

			tc.cacheStub(mockCache)
			tc.clientStub(mockClient)

			h := &HN{c: mockClient, cache: mockCache}

			// Call GetItem
			i, err := h.GetItem(context.Background(), tc.id)

			// Check error
			if tc.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// Check item
			require.Equal(t, tc.expected, i)
		})
	}
}
