package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/utils"
)

func Test_Decrease_Quality_For_Normal_Item(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "foo", SellIn: 10, Quality: 10},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 9, 9)
}

func Test_Decrease_Quality_For_Normal_Item_Quality_to_0(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "foo", SellIn: 10, Quality: 1},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 0, 9)
}

func Test_Sulfuras_Never_Changes(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 30},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 30, 0)
}

func Test_Increase_Quality_For_Aged_Brie_The_Older_It_Gets(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Aged Brie", SellIn: 10, Quality: 40},
		{Name: "Aged Brie", SellIn: 9, Quality: 41},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 41, 9)
	utils.CheckTtem(t, items[1], 42, 8)
}

func Test_Quality_Can_Never_Be_Negative_for_some_item(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Some item", SellIn: 10, Quality: 0},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 0, 9)
}

func Test_Backstage_With_Normal_SellIn(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 21, 14)
}

func Test_Backstage_With_SellIn_Smaller_Equal_Than_10(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 9, Quality: 3},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 5, 8)
}

func Test_Backstage_With_SellIn_Smaller_Equal_Than_5(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 4, Quality: 3},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 6, 3)
}

func Test_Backstage_With_SellIn_After_Concert_Date_Drops_to_Zero(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 3},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 0, -2)
}

func Test_Decrease_Quality_Twice_As_Fast_When_SellIn_Passed(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Some item", SellIn: -1, Quality: 5},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 3, -2)
}

func Test_Decrease_Quality_Twice_As_Fast_For_Conjured_Item(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Conjured Mana Cake", SellIn: 10, Quality: 12},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 10, 9)
}

func Test_Decrease_Quality_Twice_As_Fast_For_Conjured_Item_After_SellIn_Passed(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Conjured Mana Cake", SellIn: -2, Quality: 12},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 8, -3)
}

func Test_Quality_Can_Never_Be_More_Then_50(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Aged Brie", SellIn: 10, Quality: 50},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 50, 9)
}

func Test_Quality_Can_Never_Be_Less_Then_0(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Any item", SellIn: 4, Quality: 0},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 0, 3)
}
