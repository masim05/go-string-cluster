package scluster

import (
	"slices"
	"testing"
)

func TestClusterStringSet(t *testing.T) {
	testCases := []struct {
		name      string
		strings   []string
		treashold float32
		expected  [][]string
	}{
		{
			name:      "3 empty strings",
			strings:   []string{"", "", ""},
			treashold: float32(0.2),
			expected:  [][]string{{"", "", ""}},
		},
		{
			name:      "empty string and non empty string",
			strings:   []string{"", "aaaaa"},
			treashold: float32(0.2),
			expected:  [][]string{{""}, {"aaaaa"}},
		},
		{
			name:      "two similar strings",
			strings:   []string{"aaaaab", "aaaaa"},
			treashold: float32(0.2),
			expected:  [][]string{{"aaaaab", "aaaaa"}},
		},
	}

	for _, tc := range testCases {
		actual := ClusterStringSet(tc.strings, tc.treashold)

		if len(actual) != len(tc.expected) {
			t.Errorf(tc.name, "length mismatch", tc.expected, actual)
		}

		for i := 0; i < len(tc.expected); i++ {
			if !slices.Equal(actual[i], tc.expected[i]) {
				t.Errorf(tc.name, "value mismatch", tc.expected, actual)
			}
		}
	}

}
