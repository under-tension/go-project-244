package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	table := []struct {
		name     string
		filepath string
		expected map[string]any
	}{
		{name: "test one-level for json", filepath: "testdata/fixture/gendiff/input/one-level-v1.json", expected: map[string]any{"host": "hexlet.io", "timeout": float64(50), "proxy": "123.234.53.22", "follow": false}},
		{name: "test one-level for yaml", filepath: "testdata/fixture/gendiff/input/one-level-v1.yaml", expected: map[string]any{"host": "hexlet.io", "timeout": 50, "proxy": "123.234.53.22", "follow": false}},
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
	stylishExpected, _ := os.ReadFile("testdata/fixture/gendiff/expected/multi-level-stylish.txt")
	plainExpected, _ := os.ReadFile("testdata/fixture/gendiff/expected/multi-level-plain.txt")
	jsonExpected, _ := os.ReadFile("testdata/fixture/gendiff/expected/multi-level-json.json")

	table := []struct {
		name     string
		format   string
		first    string
		second   string
		expected string
	}{
		{
			name:     "json with format stylish",
			format:   "stylish",
			first:    "testdata/fixture/gendiff/input/multi-levelv1.json",
			second:   "testdata/fixture/gendiff/input/multi-levelv2.json",
			expected: string(stylishExpected),
		},
		{
			name:     "yaml with format stylish",
			format:   "stylish",
			first:    "testdata/fixture/gendiff/input/multi-levelv1.yaml",
			second:   "testdata/fixture/gendiff/input/multi-levelv2.yaml",
			expected: string(stylishExpected),
		},
		{
			name:     "json with format plain",
			format:   "plain",
			first:    "testdata/fixture/gendiff/input/multi-levelv1.json",
			second:   "testdata/fixture/gendiff/input/multi-levelv2.json",
			expected: string(plainExpected),
		},
		{
			name:     "yaml with format plain",
			format:   "plain",
			first:    "testdata/fixture/gendiff/input/multi-levelv1.yaml",
			second:   "testdata/fixture/gendiff/input/multi-levelv2.yaml",
			expected: string(plainExpected),
		},
		{
			name:     "json with format json",
			format:   "json",
			first:    "testdata/fixture/gendiff/input/multi-levelv1.json",
			second:   "testdata/fixture/gendiff/input/multi-levelv2.json",
			expected: string(jsonExpected),
		},
		{
			name:     "yaml with format json",
			format:   "json",
			first:    "testdata/fixture/gendiff/input/multi-levelv1.yaml",
			second:   "testdata/fixture/gendiff/input/multi-levelv2.yaml",
			expected: string(jsonExpected),
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			res, err := GenDiff(test.first, test.second, test.format)
			require.NoError(t, err)
			require.Equal(t, test.expected, res)
		})
	}
}
