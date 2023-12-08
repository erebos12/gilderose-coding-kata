package gildedrose

import (
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/utils"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	MAX_QUALITY    = 50
	SULFURAS       = "Sulfuras, Hand of Ragnaros"
	AGED_BRIE      = "Aged Brie"
	BACKSTAGE_PASS = "Backstage passes to a TAFKAL80ETC concert"
	CONJURED       = "Conjured Mana Cake"
)

var EXCEPTIONAL_ITEMS = []string{BACKSTAGE_PASS, AGED_BRIE, SULFURAS}

// UpdateQuality iterates through the given slice of Items, decrementing
// SellIn and updating Quality according to item type rules. It handles
// expired items by calling handleExpiredItem.
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
	if utils.ContainsString(EXCEPTIONAL_ITEMS, item.Name) {
		switch item.Name {
		case BACKSTAGE_PASS:
			handleQualityForBackStagePass(item)
		default:
			if item.Name != SULFURAS {
				item.IncreaseQualityBy(1)
			}
		}
	} else {
		decreaseValue := 1
		if item.Name == CONJURED {
			decreaseValue = 2
		}
		item.DecreaseQualityBy(decreaseValue)
	}
}

// handleExpiredItem handles updating Quality for items with expired SellIn.
// It decreases Quality for normal items and Conjured items.
// It sets Quality to 0 for Backstage Passes once SellIn is negative.
func handleExpiredItem(item *Item) {
	if !utils.ContainsString(EXCEPTIONAL_ITEMS, item.Name) {
		decreaseValue := 1
		if item.Name == CONJURED {
			decreaseValue = 2
		}
		item.DecreaseQualityBy(decreaseValue)
	} else if item.Name == BACKSTAGE_PASS {
		item.Quality = 0
	}
}

// handleQualityForBackStagePass increases the Quality of a Backstage Pass
// based on its SellIn value. The Quality is increased by 3 if SellIn is
// less than or equal to 5, by 2 if SellIn is less than or equal to 10,
// and by 1 otherwise.
func handleQualityForBackStagePass(item *Item) {
	switch {
	case item.SellIn <= 5:
		item.IncreaseQualityBy(3)
	case item.SellIn <= 10:
		item.IncreaseQualityBy(2)
	default:
		item.IncreaseQualityBy(1)
	}
}

// DecreaseQualityBy decreases the Quality of the given Item by the given
// decreaseValue, to a minimum of 0.
func (item *Item) DecreaseQualityBy(decreaseValue int) {
	if decreaseValue > item.Quality {
		item.Quality = 0
	} else {
		item.Quality -= decreaseValue
	}
}

// IncreaseQualityBy increases the Quality of the given Item by the given
// increaseValue, up to a maximum of MAX_QUALITY.
func (item *Item) IncreaseQualityBy(increaseValue int) {
	if (item.Quality + increaseValue) <= MAX_QUALITY {
		item.Quality += increaseValue
	} else {
		increaseValue = MAX_QUALITY - item.Quality
		item.Quality += increaseValue
	}
}
