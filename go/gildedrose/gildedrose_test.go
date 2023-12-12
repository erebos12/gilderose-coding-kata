package gildedrose_test

import (
	"fmt"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	tc, itemName                    string
	sellIn, quality                 int
	expectedSellIn, expectedQuality int
}

var testCases = []*TestCase{
	{"NormalItemQualityDecrease", "item X", 10, 10, 9, 9},
	{"NormalItemReachesZeroQuality", "item Z", 10, 1, 9, 0},
	{"SulfurasStaysLegendary", gildedrose.SULFURAS, 0, 30, 0, 30},
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

func TestAllTestCases(t *testing.T) {
	for _, test := range testCases {
		test := test // local copy for t.Run which runs parallel (go-routines)
		t.Run(test.tc, func(t *testing.T) {
			testQuality(t, test)
		})
	}
	fmt.Printf("Executed '%d' testcases in total", len(testCases))
}

func testQuality(t *testing.T, test *TestCase) {
	items := []*gildedrose.Item{
		{Name: test.itemName, SellIn: test.sellIn, Quality: test.quality},
	}
	gildedrose.UpdateQuality(items)
	assert.Equal(t, test.expectedQuality, items[0].Quality, "Wrong value for quality")
	assert.Equal(t, test.expectedSellIn, items[0].SellIn, "Wrong value for sellIn")
}
