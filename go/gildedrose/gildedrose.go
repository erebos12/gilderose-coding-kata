package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func (item *Item) DecreaseQualityBy(decreaseValue int) {
	if item.Quality > 0 {
		item.Quality = item.Quality - decreaseValue
	}
}

func (item *Item) IncreaseQualityBy(increaseValue int) {
	if (item.Quality + increaseValue) <= MAX_QUALITY {
		item.Quality = item.Quality + increaseValue
	} else {
		increaseValue = MAX_QUALITY - item.Quality
		item.Quality += increaseValue
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
				item.DecreaseQualityBy(2)
			} else {
				item.DecreaseQualityBy(1)
			}
		} else {
			if item.Name == BACKSTAGE_PASS {
				handleQualityForBackStagePass(item)
			} else {
				if item.Name != SULFURAS {
					item.IncreaseQualityBy(1)
				}
			}
		}
		if item.SellIn < 0 {
			if !containsString(EXCEPTIONAL_ITEMS, item.Name) {
				if item.Name == CONJURED {
					item.DecreaseQualityBy(2)
				} else {
					item.DecreaseQualityBy(1)
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

func handleQualityForBackStagePass(item *Item) {
	switch {
	case item.SellIn < 0:
		item.Quality = 0
	case item.SellIn <= 5:
		item.IncreaseQualityBy(3)
	case item.SellIn <= 10:
		item.IncreaseQualityBy(2)
	default:
		item.IncreaseQualityBy(1)
	}
}
