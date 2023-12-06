package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func (item *Item) decreaseQualityBy(decreaseValue int) {
	if item.Quality > 0 {
		item.Quality = item.Quality - decreaseValue
	}
}

func (item *Item) increseQualityBy(increaseValue int) {
	if item.Quality < MAX_QUALITY {
		item.Quality = item.Quality + increaseValue
	}
}

var MAX_QUALITY = 50

var SULFURAS = "Sulfuras, Hand of Ragnaros"
var AGED_BRIE = "Aged Brie"
var BACKSTAGE_PASS = "Backstage passes to a TAFKAL80ETC concert"
var CONJURED = "Conjured Mana Cake"
var EXCEPTIONAL_ITEMS = []string{BACKSTAGE_PASS, AGED_BRIE, SULFURAS}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		if item.Name != SULFURAS {
			item.SellIn = item.SellIn - 1
		}
		if !containsString(EXCEPTIONAL_ITEMS, item.Name) {
			if item.Name == CONJURED {
				item.decreaseQualityBy(2)
			} else {
				item.decreaseQualityBy(1)
			}
		} else {
			if item.Name == BACKSTAGE_PASS {
				handleQualityForBackStagePass(item)
			} else {
				if item.Name != SULFURAS {
					item.increseQualityBy(1)
				}
			}
		}
		if item.SellIn < 0 {
			if !containsString(EXCEPTIONAL_ITEMS, item.Name) {
				if item.Name == CONJURED {
					item.decreaseQualityBy(2)
				} else {
					item.decreaseQualityBy(1)
				}
			}
		}
	}
}

func containsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func decreaseQualityForItem(item *Item) {
	if !containsString(EXCEPTIONAL_ITEMS, item.Name) {
		if item.Name == CONJURED {
			item.decreaseQualityBy(2)
		} else {
			item.decreaseQualityBy(1)
		}
	}
}

func handleQualityForBackStagePass(item *Item) {
	switch {
	case item.SellIn < 0:
		item.Quality = 0
	case item.SellIn <= 5:
		item.increseQualityBy(3)
	case item.SellIn <= 10:
		item.increseQualityBy(2)

	}
}
