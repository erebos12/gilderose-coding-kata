package gildedrose

import (
	"slices"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	MAX_QUALITY    = 50
	MIN_QUALITY    = 0
	SULFURAS       = "Sulfuras, Hand of Ragnaros"
	AGED_BRIE      = "Aged Brie"
	BACKSTAGE_PASS = "Backstage passes to a TAFKAL80ETC concert"
	CONJURED       = "Conjured Mana Cake"
)

var EXCEPTIONAL_ITEMS = []string{BACKSTAGE_PASS, AGED_BRIE, SULFURAS}

// UpdateQuality iterates through the given list of items, decrementing
// SellIn and updating Quality for each item. It handles special cases
// for expired items and items like Sulfuras. This function encapsulates
// the core business logic for decrementing inventory.
func UpdateQuality(items []*Item) {
	for _, item := range items {
		if item.Name != SULFURAS {
			item.SellIn--
		}
		updateItemQuality(item)
		if item.SellIn < 0 {
			handleExpiredItem(item)
		}
	}
}

// updateItemQuality updates the quality of the given item based on its properties.
// It handles increasing or decreasing quality for exceptional items like Aged Brie and Backstage Passes,
// as well as decreasing quality for normal items and Conjured items.
func updateItemQuality(item *Item) {
	if slices.Contains(EXCEPTIONAL_ITEMS, item.Name) {
		switch item.Name {
		case BACKSTAGE_PASS:
			handleQualityForBackStagePass(item)
		default:
			if item.Name != SULFURAS {
				item.increaseQualityBy(1)
			}
		}
	} else {
		decreaseValue := 1
		if item.Name == CONJURED {
			decreaseValue = 2
		}
		item.decreaseQualityBy(decreaseValue)
	}
}

// handleExpiredItem handles updating the Quality of items whose SellIn has
// expired (gone below 0). It decreases Quality for normal items, and sets
// Quality to 0 for Backstage Passes. It does not modify Sulfuras items.
func handleExpiredItem(item *Item) {
	if slices.Contains(EXCEPTIONAL_ITEMS, item.Name) {
		if item.Name == BACKSTAGE_PASS {
			item.Quality = MIN_QUALITY
		}
	} else {
		decreaseValue := 1
		if item.Name == CONJURED {
			decreaseValue = 2
		}
		item.decreaseQualityBy(decreaseValue)
	}
}

// handleQualityForBackStagePass increases the Quality of a Backstage Pass
// based on its SellIn value. The Quality is increased by 3 if SellIn is
// less than or equal to 5, by 2 if SellIn is less than or equal to 10,
// and by 1 otherwise.
func handleQualityForBackStagePass(item *Item) {
	switch {
	case item.SellIn <= 5:
		item.increaseQualityBy(3)
	case item.SellIn <= 10:
		item.increaseQualityBy(2)
	default:
		item.increaseQualityBy(1)
	}
}

// DecreaseQualityBy decreases the Quality of the given Item by the given
// decreaseValue, to a minimum of 0.
func (item *Item) decreaseQualityBy(decreaseValue int) {
	if decreaseValue > item.Quality {
		item.Quality = MIN_QUALITY
	} else {
		item.Quality -= decreaseValue
	}
}

// IncreaseQualityBy increases the Quality of the given Item by the given
// increaseValue, up to a maximum of MAX_QUALITY.
func (item *Item) increaseQualityBy(increaseValue int) {
	if (item.Quality + increaseValue) <= MAX_QUALITY {
		item.Quality += increaseValue
	} else {
		increaseValue = MAX_QUALITY - item.Quality
		item.Quality += increaseValue
	}
}
