package utils_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/utils"
	"github.com/stretchr/testify/assert"
)

func TestContainsString(t *testing.T) {
	slice := []string{"foo", "bar"}

	contains := utils.ContainsString(slice, "foo")
	assert.True(t, contains)

	contains = utils.ContainsString(slice, "baz")
	assert.False(t, contains)

	contains = utils.ContainsString([]string{}, "foo")
	assert.False(t, contains)
}
