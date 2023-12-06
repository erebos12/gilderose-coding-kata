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
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 8, Quality: 5},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 7, Quality: 9},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	utils.CheckTtem(t, items[0], 5, 8)
	utils.CheckTtem(t, items[1], 7, 7)
	utils.CheckTtem(t, items[2], 11, 6)
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

func Test_Fixture(t *testing.T) {

	var items = []*gildedrose.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
		{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 49},
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6}, // <-- :O
	}
	days := 4
	for day := 0; day < days; day++ {
		if day == 1 {
			utils.CheckTtem(t, items[0], 19, 9)
			utils.CheckTtem(t, items[1], 1, 1)
			utils.CheckTtem(t, items[2], 6, 4)
			utils.CheckTtem(t, items[3], 80, 0)
			utils.CheckTtem(t, items[4], 80, -1)
			utils.CheckTtem(t, items[5], 21, 14)
			utils.CheckTtem(t, items[6], 50, 9)
			utils.CheckTtem(t, items[7], 50, 4)
			utils.CheckTtem(t, items[8], 4, 2)
		}
		if day == 2 {
			utils.CheckTtem(t, items[0], 18, 8)
			utils.CheckTtem(t, items[1], 2, 0)
			utils.CheckTtem(t, items[2], 5, 3)
			utils.CheckTtem(t, items[3], 80, 0)
			utils.CheckTtem(t, items[4], 80, -1)
			utils.CheckTtem(t, items[5], 22, 13)
			utils.CheckTtem(t, items[6], 50, 8)
			utils.CheckTtem(t, items[7], 50, 3)
			utils.CheckTtem(t, items[8], 2, 1)
		}
		if day == 3 {
			utils.CheckTtem(t, items[0], 17, 7)
			utils.CheckTtem(t, items[1], 3, -1)
			utils.CheckTtem(t, items[2], 4, 2)
			utils.CheckTtem(t, items[3], 80, 0)
			utils.CheckTtem(t, items[4], 80, -1)
			utils.CheckTtem(t, items[5], 23, 12)
			utils.CheckTtem(t, items[6], 50, 7)
			utils.CheckTtem(t, items[7], 50, 2)
			utils.CheckTtem(t, items[8], 0, 0)
		}
		gildedrose.UpdateQuality(items)
	}
}
