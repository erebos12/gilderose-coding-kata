package utils

import (
	"runtime"
	"strings"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func CheckTtem(t *testing.T, item *gildedrose.Item, expectedQuality int, expectedSellIn int) (success, error bool) {
	if item.Quality != expectedQuality {
		pc, _, _, _ := runtime.Caller(1)
		function := runtime.FuncForPC(pc)
		t.Errorf("Test '%s' failed. Quality must be '%d' but got '%d'.", ExtractTestName(function.Name()), expectedQuality, item.Quality)
	}
	if item.SellIn != expectedSellIn {
		pc, _, _, _ := runtime.Caller(1)
		function := runtime.FuncForPC(pc)
		t.Errorf("Test '%s' failed. SellIn must be '%d' but got '%d'.", ExtractTestName(function.Name()), expectedSellIn, item.SellIn)
	}
	return true, false
}

func ExtractTestName(input string) string {
	parts := strings.Split(input, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return "NO_FUNCTION_NAME_FOUND"
}
