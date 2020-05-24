package go_compare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func compareInt(x, y int) Linear {
	if x == y {
		return Equal
	}
	if x < y {
		return LessThan
	}
	return GreaterThan
}

func TestIsGreaterThanOrEqual(t *testing.T) {
	cases := map[string]struct {
		left, right int
		expected    bool
	}{
		"less than": {
			left:     5,
			right:    10,
			expected: false,
		},
		"equal": {
			left:     10,
			right:    10,
			expected: true,
		},
		"greater than": {
			left:     15,
			right:    10,
			expected: true,
		},
	}

	for caseName, c := range cases {
		actual := IsGreaterThanOrEqual(compareInt(c.left, c.right))
		assert.Equal(t, c.expected, actual, caseName)
	}
}

func TestIsLessThanOrEqual(t *testing.T) {
	cases := map[string]struct {
		left, right int
		expected    bool
	}{
		"less than": {
			left:     5,
			right:    10,
			expected: true,
		},
		"equal": {
			left:     10,
			right:    10,
			expected: true,
		},
		"greater than": {
			left:     15,
			right:    10,
			expected: false,
		},
	}

	for caseName, c := range cases {
		actual := IsLessThanOrEqual(compareInt(c.left, c.right))
		assert.Equal(t, c.expected, actual, caseName)
	}
}
