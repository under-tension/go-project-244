package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	table := []struct {
		name     string
		filepath string
		expected map[string]any
	}{
		{name: "test one-level for json", filepath: "testdata/fixture/file-one-level-v1.json", expected: map[string]any{"host": "hexlet.io", "timeout": float64(50), "proxy": "123.234.53.22", "follow": false}},
		{name: "test one-level for yaml", filepath: "testdata/fixture/file-one-level-v1.yaml", expected: map[string]any{"host": "hexlet.io", "timeout": 50, "proxy": "123.234.53.22", "follow": false}},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			res, err := ParseFile(test.filepath)
			require.NoError(t, err)
			require.Equal(t, test.expected, res)
		})
	}
}

func TestGenDiff(t *testing.T) {
	table := []struct {
		name     string
		first    map[string]any
		second   map[string]any
		expected string
	}{
		{
			name:   "test one-level",
			first:  map[string]any{"host": "hexlet.io", "timeout": float64(50), "proxy": "123.234.53.22", "follow": false},
			second: map[string]any{"timeout": float64(20), "verbose": true, "host": "hexlet.io"},
			expected: `{
	- follow: false
	  host: "hexlet.io"
	- proxy: "123.234.53.22"
	- timeout: 50
	+ timeout: 20
	+ verbose: true
}
`,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			res, err := GenDiff(test.first, test.second, "stylish")
			require.NoError(t, err)
			require.Equal(t, test.expected, res)
		})
	}
}
