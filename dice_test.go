package yatzy

import (
	"testing"

	"github.com/matryer/is"
)

func TestGetDiceMap(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected map[string]int
	}{
		{
			name:     "dice value 1",
			value:    1,
			expected: map[string]int{"\n . \n": 1},
		},
		{
			name:     "dice value 2",
			value:    2,
			expected: map[string]int{"  .\n\n.  ": 2},
		},
		{
			name:     "dice value 3",
			value:    3,
			expected: map[string]int{".  \n . \n  .": 3},
		},
		{
			name:     "dice value 4",
			value:    4,
			expected: map[string]int{". .\n\n. .": 4},
		},
		{
			name:     "dice value 5",
			value:    5,
			expected: map[string]int{". .\n . \n. .": 5},
		},
		{
			name:     "dice value 6",
			value:    6,
			expected: map[string]int{". .\n. .\n. .": 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			result := getDiceMap(tt.value)
			is.Equal(result, tt.expected)
		})
	}
}
