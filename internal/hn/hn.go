package hn

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/KarolosLykos/hackertea/internal/cache"
	"github.com/KarolosLykos/hackertea/internal/client"
	"github.com/KarolosLykos/hackertea/internal/constants"
	"github.com/KarolosLykos/hackertea/internal/item"
)

var ErrInvalidItemType = errors.New("invalid item type")

type Service interface {
	GetItems(ctx context.Context, item constants.ItemType) ([]int, error)
	GetItem(ctx context.Context, id int) (*item.Item, error)
}

type HN struct {
	c     client.HttpClient
	cache cache.Cache
}

func New(c client.HttpClient, cache cache.Cache) *HN {
	return &HN{c: c, cache: cache}
}

func (h *HN) GetItems(ctx context.Context, item constants.ItemType) ([]int, error) {
	suffix, err := getSuffix(item)
	if err != nil {
		return nil, err
	}

	resp, err := h.c.Get(ctx, suffix)
	if err != nil {
		return nil, err
	}

	items := make([]int, 0)
	if err = json.Unmarshal(resp, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (h *HN) GetItem(ctx context.Context, id int) (*item.Item, error) {
	v, ok := h.cache.Get(id)
	if ok {
		return v, nil
	}

	suffix, _ := getSuffix(constants.Items.SingleItem)

	uri := fmt.Sprintf(suffix, strconv.Itoa(id))

	resp, err := h.c.Get(ctx, uri)
	if err != nil {
		return nil, err
	}

	i := &item.Item{}
	if err = json.Unmarshal(resp, i); err != nil {
		return nil, err
	}

	h.cache.Set(id, i)

	return i, nil
}

func getSuffix(item constants.ItemType) (string, error) {
	switch item {
	case constants.Items.NewItems:
		return constants.NewSuffix, nil
	case constants.Items.TopItems:
		return constants.TopSuffix, nil
	case constants.Items.BestItems:
		return constants.BestSuffix, nil
	case constants.Items.SingleItem:
		return constants.SingleSuffix, nil
	default:
		return "", ErrInvalidItemType
	}
}
