package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

func TestNormalItemQualityDecrease(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "foo", SellIn: 10, Quality: 10},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 9
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 9
	assert.Equal(t, expectedSellIn, items[0].SellIn)

}

func TestNormalItemReachesZeroQuality(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "foo", SellIn: 10, Quality: 1},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 0
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 9
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestSulfurasStaysLegendary(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.SULFURAS, SellIn: 0, Quality: 30},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 30
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 0
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestAgedBrieGetsBetterWithAge(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.AGED_BRIE, SellIn: 10, Quality: 40},
		{Name: gildedrose.AGED_BRIE, SellIn: 9, Quality: 41},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 41
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 9
	assert.Equal(t, expectedSellIn, items[0].SellIn)

	expectedQuality = 42
	assert.Equal(t, expectedQuality, items[1].Quality)
	expectedSellIn = 8
	assert.Equal(t, expectedSellIn, items[1].SellIn)
}

func TestQualityNeverDropsBelowZero(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Some item", SellIn: 10, Quality: 0},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 0
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 9
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestBackstageQualityWithNormalSellIn(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.BACKSTAGE_PASS, SellIn: 15, Quality: 20},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 21
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 14
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestBackstagePassQualityWithSellInLessThan10Days(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.BACKSTAGE_PASS, SellIn: 9, Quality: 3},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 5
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 8
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestBackstagePassQualityWithSellInLessThan5Days(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.BACKSTAGE_PASS, SellIn: 4, Quality: 3},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 6
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 3
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestBackstagePassQualityDropsToZeroAfterConcert(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.BACKSTAGE_PASS, SellIn: -1, Quality: 3},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 0
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := -2
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestExpiredItemReducesQualityTwiceAsFast(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Some item", SellIn: -1, Quality: 5},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 3
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := -2
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestDecreaseQualityTwiceAsFastForConjuredItem(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.CONJURED, SellIn: 10, Quality: 12},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 10
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 9
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestQualityReductionDoublesForConjuredItemAfterSellIn(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.CONJURED, SellIn: -2, Quality: 12},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 8
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := -3
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func TestQualityNeverExceedsFifty(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: gildedrose.AGED_BRIE, SellIn: 10, Quality: 50},
	}
	// when
	gildedrose.UpdateQuality(items)
	// then
	expectedQuality := 50
	assert.Equal(t, expectedQuality, items[0].Quality)
	expectedSellIn := 9
	assert.Equal(t, expectedSellIn, items[0].SellIn)
}

func Test_Fixture(t *testing.T) {

	var items = []*gildedrose.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
		{Name: gildedrose.AGED_BRIE, SellIn: 2, Quality: 0},
		{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
		{Name: gildedrose.SULFURAS, SellIn: 0, Quality: 80},
		{Name: gildedrose.SULFURAS, SellIn: -1, Quality: 80},
		{Name: gildedrose.BACKSTAGE_PASS, SellIn: 15, Quality: 20},
		{Name: gildedrose.BACKSTAGE_PASS, SellIn: 10, Quality: 49},
		{Name: gildedrose.BACKSTAGE_PASS, SellIn: 5, Quality: 49},
		{Name: gildedrose.CONJURED, SellIn: 3, Quality: 6}, // <-- :O
	}
	days := 4
	for day := 0; day < days; day++ {
		if day == 1 {
			expectedQuality := 19
			assert.Equal(t, expectedQuality, items[0].Quality)
			expectedSellIn := 9
			assert.Equal(t, expectedSellIn, items[0].SellIn)

			expectedQuality = 1
			assert.Equal(t, expectedQuality, items[1].Quality)
			expectedSellIn = 1
			assert.Equal(t, expectedSellIn, items[1].SellIn)

			expectedQuality = 6
			assert.Equal(t, expectedQuality, items[2].Quality)
			expectedSellIn = 4
			assert.Equal(t, expectedSellIn, items[2].SellIn)

			expectedQuality = 80
			assert.Equal(t, expectedQuality, items[3].Quality)
			expectedSellIn = 0
			assert.Equal(t, expectedSellIn, items[3].SellIn)

			expectedQuality = 80
			assert.Equal(t, expectedQuality, items[4].Quality)
			expectedSellIn = -1
			assert.Equal(t, expectedSellIn, items[4].SellIn)

			expectedQuality = 21
			assert.Equal(t, expectedQuality, items[5].Quality)
			expectedSellIn = 14
			assert.Equal(t, expectedSellIn, items[5].SellIn)

			expectedQuality = 50
			assert.Equal(t, expectedQuality, items[6].Quality)
			expectedSellIn = 9
			assert.Equal(t, expectedSellIn, items[6].SellIn)

			expectedQuality = 50
			assert.Equal(t, expectedQuality, items[7].Quality)
			expectedSellIn = 4
			assert.Equal(t, expectedSellIn, items[7].SellIn)

			expectedQuality = 4
			assert.Equal(t, expectedQuality, items[8].Quality)
			expectedSellIn = 2
			assert.Equal(t, expectedSellIn, items[8].SellIn)
		}
		if day == 2 {
			expectedQuality := 18
			assert.Equal(t, expectedQuality, items[0].Quality)
			expectedSellIn := 8
			assert.Equal(t, expectedSellIn, items[0].SellIn)

			expectedQuality = 2
			assert.Equal(t, expectedQuality, items[1].Quality)
			expectedSellIn = 0
			assert.Equal(t, expectedSellIn, items[1].SellIn)

			expectedQuality = 5
			assert.Equal(t, expectedQuality, items[2].Quality)
			expectedSellIn = 3
			assert.Equal(t, expectedSellIn, items[2].SellIn)

			expectedQuality = 80
			assert.Equal(t, expectedQuality, items[3].Quality)
			expectedSellIn = 0
			assert.Equal(t, expectedSellIn, items[3].SellIn)

			expectedQuality = 80
			assert.Equal(t, expectedQuality, items[4].Quality)
			expectedSellIn = -1
			assert.Equal(t, expectedSellIn, items[4].SellIn)

			expectedQuality = 22
			assert.Equal(t, expectedQuality, items[5].Quality)
			expectedSellIn = 13
			assert.Equal(t, expectedSellIn, items[5].SellIn)

			expectedQuality = 50
			assert.Equal(t, expectedQuality, items[6].Quality)
			expectedSellIn = 8
			assert.Equal(t, expectedSellIn, items[6].SellIn)

			expectedQuality = 50
			assert.Equal(t, expectedQuality, items[7].Quality)
			expectedSellIn = 3
			assert.Equal(t, expectedSellIn, items[7].SellIn)

			expectedQuality = 2
			assert.Equal(t, expectedQuality, items[8].Quality)
			expectedSellIn = 1
			assert.Equal(t, expectedSellIn, items[8].SellIn)
		}
		if day == 3 {
			expectedQuality := 17
			assert.Equal(t, expectedQuality, items[0].Quality)
			expectedSellIn := 7
			assert.Equal(t, expectedSellIn, items[0].SellIn)

			expectedQuality = 3
			assert.Equal(t, expectedQuality, items[1].Quality)
			expectedSellIn = -1
			assert.Equal(t, expectedSellIn, items[1].SellIn)

			expectedQuality = 4
			assert.Equal(t, expectedQuality, items[2].Quality)
			expectedSellIn = 2
			assert.Equal(t, expectedSellIn, items[2].SellIn)

			expectedQuality = 80
			assert.Equal(t, expectedQuality, items[3].Quality)
			expectedSellIn = 0
			assert.Equal(t, expectedSellIn, items[3].SellIn)

			expectedQuality = 80
			assert.Equal(t, expectedQuality, items[4].Quality)
			expectedSellIn = -1
			assert.Equal(t, expectedSellIn, items[4].SellIn)

			expectedQuality = 23
			assert.Equal(t, expectedQuality, items[5].Quality)
			expectedSellIn = 12
			assert.Equal(t, expectedSellIn, items[5].SellIn)

			expectedQuality = 50
			assert.Equal(t, expectedQuality, items[6].Quality)
			expectedSellIn = 7
			assert.Equal(t, expectedSellIn, items[6].SellIn)

			expectedQuality = 50
			assert.Equal(t, expectedQuality, items[7].Quality)
			expectedSellIn = 2
			assert.Equal(t, expectedSellIn, items[7].SellIn)

			expectedQuality = 0
			assert.Equal(t, expectedQuality, items[8].Quality)
			expectedSellIn = 0
			assert.Equal(t, expectedSellIn, items[8].SellIn)
		}
		gildedrose.UpdateQuality(items)
	}
}

func Test_IncreaseQualityBy_Function_01(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Any item", SellIn: 4, Quality: 48},
	}
	// when
	items[0].IncreaseQualityBy(4)
	// then
	expectedQuality := 50
	assert.Equal(t, expectedQuality, items[0].Quality)
}

func Test_IncreaseQualityBy_Function_02(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Any item", SellIn: 6, Quality: 48},
	}
	// when
	items[0].IncreaseQualityBy(2)
	// then
	expectedQuality := 50
	assert.Equal(t, expectedQuality, items[0].Quality)
}

func Test_IncreaseQualityBy_Function_03(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Any item", SellIn: 17, Quality: 50},
	}
	// when
	items[0].IncreaseQualityBy(4)
	// then
	expectedQuality := 50
	assert.Equal(t, expectedQuality, items[0].Quality)
}

func Test_DecreaseQualityBy_4_With_Quality_2(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Any item", SellIn: 17, Quality: 2},
	}
	// when
	items[0].DecreaseQualityBy(4)
	// then
	expectedQuality := 0
	assert.Equal(t, expectedQuality, items[0].Quality)
}

func Test_DecreaseQualityBy_4_With_Quality_4(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Any item", SellIn: 17, Quality: 4},
	}
	// when
	items[0].DecreaseQualityBy(4)
	// then
	expectedQuality := 0
	assert.Equal(t, expectedQuality, items[0].Quality)
}

func Test_DecreaseQualityBy_2_With_Quality_3(t *testing.T) {
	// given
	var items = []*gildedrose.Item{
		{Name: "Any item", SellIn: 17, Quality: 3},
	}
	// when
	items[0].DecreaseQualityBy(2)
	// then
	expectedQuality := 1
	assert.Equal(t, expectedQuality, items[0].Quality)
}
