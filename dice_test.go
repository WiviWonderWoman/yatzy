package yatzy

import (
	"testing"

	"github.com/matryer/is"
)

func TestGetKey(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{
			name:     "dice value 1",
			value:    1,
			expected: "  .  ",
		},
		{
			name:     "dice value 2",
			value:    2,
			expected: "    .\n\n.    ",
		},
		{
			name:     "dice value 3",
			value:    3,
			expected: ".    \n  .  \n    .",
		},
		{
			name:     "dice value 4",
			value:    4,
			expected: ".    .\n\n.    .",
		},
		{
			name:     "dice value 5",
			value:    5,
			expected: ".    .\n . \n.    .",
		},
		{
			name:     "dice value 6",
			value:    6,
			expected: ".    .\n.    .\n.    .",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := GetKey(tt.value)
			is.Equal(result, tt.expected)
		})
	}
}

func TestGetRandomValue(t *testing.T) {
	is := is.New(t)

	// Test that the function returns values within the expected range
	for i := 0; i < 100; i++ { // Run multiple times to ensure randomness
		result := GetRandomValue()
		is.True(result >= 1 && result <= 6) // Value should be between 1 and 6
	}
}
