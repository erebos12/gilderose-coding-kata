package gildedrose

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

func updateItemQuality(item *Item) {
	if containsString(EXCEPTIONAL_ITEMS, item.Name) {
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

func handleExpiredItem(item *Item) {
	if !containsString(EXCEPTIONAL_ITEMS, item.Name) {
		decreaseValue := 1
		if item.Name == CONJURED {
			decreaseValue = 2
		}
		item.DecreaseQualityBy(decreaseValue)
	} else if item.Name == BACKSTAGE_PASS {
		item.Quality = 0
	}
}

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

func (item *Item) DecreaseQualityBy(decreaseValue int) {
	if decreaseValue > item.Quality {
		item.Quality = 0
	} else {
		item.Quality -= decreaseValue
	}
}

func (item *Item) IncreaseQualityBy(increaseValue int) {
	if (item.Quality + increaseValue) <= MAX_QUALITY {
		item.Quality += increaseValue
	} else {
		increaseValue = MAX_QUALITY - item.Quality
		item.Quality += increaseValue
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
