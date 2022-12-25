package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	assert.Equal(t, "one", Serialize("one"))
	assert.Equal(t, "a b", Serialize("a", "b"))
	assert.Equal(t, "1 b 89", Serialize(1, "b", 89))
}

func TestContains(t *testing.T) {
	assert.True(t, Contains([]string{"a", "b"}, "b"))
	assert.False(t, Contains([]string{"a", "b"}, "c"))
	assert.True(t, Contains([]int{2, 3}, 3))
	assert.False(t, Contains([]int{12, 13}, 1))
}

func TestRemove(t *testing.T) {
	assert.Equal(t, []string{"a"}, Remove([]string{"a", "b"}, "b"))
	assert.Equal(t, []string{"a", "c"}, Remove([]string{"a", "b", "c", "b"}, "b"))
	assert.Equal(t, []string{"a"}, Remove([]string{"a", "b", "c", "b"}, "b", "c"))
	assert.Equal(t, []int{1, 2, 5}, Remove([]int{1, 3, 2, 5}, 3))
	assert.Equal(t, []int{1, 3, 2, 5}, Remove([]int{1, 3, 2, 5}, 4))
}
