package gildedrose_test

import (
	"fmt"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	testCases := []struct {
		tc, itemName                    string
		sellIn, quality                 int
		expectedSellIn, expectedQuality int
	}{
		{"NormalItemQualityDecrease", "item X", 10, 10, 9, 9},
		{"NormalItemReachesZeroQuality", "item Z", 10, 1, 9, 0},
		{"NormalItemReachesZeroQuality", gildedrose.SULFURAS, 0, 30, 0, 30},
		{"AgedBrieGetsBetterWithAge", gildedrose.AGED_BRIE, 10, 40, 9, 41},
		{"QualityNeverDropsBelowZero", "item A", 10, 0, 9, 0},
		{"QualityNeverExceedsFifty", gildedrose.AGED_BRIE, 10, 50, 9, 50},
		{"BackstageQualityWithNormalSellIn", gildedrose.BACKSTAGE_PASS, 15, 20, 14, 21},
		{"BackstagePassQualityWithSellInLessThan10Days", gildedrose.BACKSTAGE_PASS, 9, 3, 8, 5},
		{"BackstagePassQualityWithSellInEqual10Days", gildedrose.BACKSTAGE_PASS, 10, 7, 9, 9},
		{"BackstagePassQualityWithSellInLessThan5Days", gildedrose.BACKSTAGE_PASS, 4, 3, 3, 6},
		{"BackstagePassQualityWithSellInEqual5Days", gildedrose.BACKSTAGE_PASS, 5, 3, 4, 6},
		{"BackstagePassQualityDropsToZeroAfterConcert", gildedrose.BACKSTAGE_PASS, -1, 3, -2, 0},
		{"ExpiredItemReducesQualityTwiceAsFast", "item B", -1, 5, -2, 3},
		{"DecreaseQualityTwiceAsFastForConjuredItem", gildedrose.CONJURED, 10, 12, 9, 10},
		{"QualityReductionDoublesForConjuredItemAfterSellIn", gildedrose.CONJURED, -2, 12, -3, 8},
	}

	for _, test := range testCases {
		t.Run(test.tc, func(t *testing.T) {
			items := []*gildedrose.Item{
				{Name: test.itemName, SellIn: test.sellIn, Quality: test.quality},
			}
			gildedrose.UpdateQuality(items)
			assert.Equal(t, test.expectedQuality, items[0].Quality, "Wrong value for quality")
			assert.Equal(t, test.expectedSellIn, items[0].SellIn, "Wrong value for sellIn")
		})
	}
	fmt.Printf("Executed '%d' testcases in total", len(testCases))
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
