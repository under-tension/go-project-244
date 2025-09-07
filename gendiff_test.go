package code

import (
	"code/pkg/formatters"
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

func TestIsMap(t *testing.T) {
	table := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{
			name:     "map[string]interface{}",
			input:    map[string]interface{}{"a": 1},
			expected: true,
		},
		{
			name:     "map[string]string",
			input:    map[string]string{"a": "b"},
			expected: true,
		},
		{
			name:     "nil",
			input:    nil,
			expected: false,
		},
		{
			name:     "string",
			input:    "not a map",
			expected: false,
		},
		{
			name:     "int",
			input:    42,
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]interface{}{},
			expected: true,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			result := isMap(test.input)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestDiffCalc(t *testing.T) {
	t.Run("both maps are empty", func(t *testing.T) {
		first := map[string]interface{}{}
		second := map[string]interface{}{}
		diff := diffCalc(first, second)
		require.Empty(t, diff)
	})

	t.Run("first map is empty", func(t *testing.T) {
		first := map[string]interface{}{}
		second := map[string]interface{}{"a": 1}
		diff := diffCalc(first, second)

		require.Len(t, diff, 1)
		require.Equal(t, "a", diff[0].Name)
		require.Equal(t, formatters.STATUS_ADDED, diff[0].Status)
		require.Equal(t, 1, diff[0].Val)
	})

	t.Run("second map is empty", func(t *testing.T) {
		first := map[string]interface{}{"a": 1}
		second := map[string]interface{}{}
		diff := diffCalc(first, second)

		require.Len(t, diff, 1)
		require.Equal(t, "a", diff[0].Name)
		require.Equal(t, formatters.STATUS_DELETED, diff[0].Status)
		require.Equal(t, 1, diff[0].OldVal)
	})

	t.Run("same keys, same values", func(t *testing.T) {
		first := map[string]interface{}{"a": 1, "b": "str"}
		second := map[string]interface{}{"a": 1, "b": "str"}
		diff := diffCalc(first, second)

		require.Len(t, diff, 2)
		require.Equal(t, "a", diff[0].Name)
		require.Equal(t, formatters.STATUS_NON_CHANGE, diff[0].Status)
		require.Equal(t, 1, diff[0].Val)

		require.Equal(t, "b", diff[1].Name)
		require.Equal(t, formatters.STATUS_NON_CHANGE, diff[1].Status)
		require.Equal(t, "str", diff[1].Val)

		// Проверяем сортировку по имени
		require.Equal(t, []string{"a", "b"}, []string{diff[0].Name, diff[1].Name})
	})

	t.Run("same keys, different values", func(t *testing.T) {
		first := map[string]interface{}{"a": 1}
		second := map[string]interface{}{"a": 2}
		diff := diffCalc(first, second)

		require.Len(t, diff, 1)
		require.Equal(t, "a", diff[0].Name)
		require.Equal(t, formatters.STATUS_UPDATED, diff[0].Status)
		require.Equal(t, 1, diff[0].OldVal)
		require.Equal(t, 2, diff[0].Val)
	})

	t.Run("nested maps", func(t *testing.T) {
		first := map[string]interface{}{
			"common": map[string]interface{}{
				"key": "value",
			},
		}
		second := map[string]interface{}{
			"common": map[string]interface{}{
				"key": "updated",
			},
		}

		diff := diffCalc(first, second)
		require.Len(t, diff, 1)
		require.Equal(t, "common", diff[0].Name)
		require.Equal(t, formatters.TYPE_ROOT, diff[0].Type)
		require.Equal(t, formatters.STATUS_NON_CHANGE, diff[0].Status)

		subDiff, ok := diff[0].Val.([]formatters.DiffTree)
		require.True(t, ok)
		require.Len(t, subDiff, 1)
		require.Equal(t, "key", subDiff[0].Name)
		require.Equal(t, formatters.STATUS_UPDATED, subDiff[0].Status)
		require.Equal(t, "value", subDiff[0].OldVal)
		require.Equal(t, "updated", subDiff[0].Val)
	})

	t.Run("nested maps with added key", func(t *testing.T) {
		first := map[string]interface{}{
			"nested": map[string]interface{}{"a": 1},
		}
		second := map[string]interface{}{
			"nested": map[string]interface{}{"a": 1, "b": 2},
		}

		diff := diffCalc(first, second)
		require.Len(t, diff, 1)
		require.Equal(t, "nested", diff[0].Name)

		subDiff, ok := diff[0].Val.([]formatters.DiffTree)
		require.True(t, ok)
		require.Len(t, subDiff, 2)
		require.Equal(t, "a", subDiff[0].Name)
		require.Equal(t, formatters.STATUS_NON_CHANGE, subDiff[0].Status)
		require.Equal(t, 1, subDiff[0].Val)

		require.Equal(t, "b", subDiff[1].Name)
		require.Equal(t, formatters.STATUS_ADDED, subDiff[1].Status)
		require.Equal(t, 2, subDiff[1].Val)
	})

	t.Run("mixed: added, deleted, updated, unchanged", func(t *testing.T) {
		first := map[string]interface{}{
			"deleted":   "removed",
			"unchanged": "same",
			"updated":   "old",
			"nested":    map[string]interface{}{"inner": "val1"},
		}
		second := map[string]interface{}{
			"added":     "new",
			"unchanged": "same",
			"updated":   "new",
			"nested":    map[string]interface{}{"inner": "val2"},
		}

		diff := diffCalc(first, second)
		require.Len(t, diff, 5)

		// Проверяем сортировку по имени
		expectedOrder := []string{"added", "deleted", "nested", "unchanged", "updated"}
		for i, name := range expectedOrder {
			require.Equal(t, name, diff[i].Name)
		}

		// added
		require.Equal(t, formatters.STATUS_ADDED, diff[0].Status)
		require.Equal(t, "new", diff[0].Val)

		// deleted
		require.Equal(t, formatters.STATUS_DELETED, diff[1].Status)
		require.Equal(t, "removed", diff[1].OldVal)

		// nested
		require.Equal(t, formatters.TYPE_ROOT, diff[2].Type)
		subDiff, ok := diff[2].Val.([]formatters.DiffTree)
		require.True(t, ok)
		require.Len(t, subDiff, 1)
		require.Equal(t, "inner", subDiff[0].Name)
		require.Equal(t, formatters.STATUS_UPDATED, subDiff[0].Status)
		require.Equal(t, "val1", subDiff[0].OldVal)
		require.Equal(t, "val2", subDiff[0].Val)

		// unchanged
		require.Equal(t, formatters.STATUS_NON_CHANGE, diff[3].Status)
		require.Equal(t, nil, diff[3].OldVal)
		require.Equal(t, "same", diff[3].Val)

		// updated
		require.Equal(t, formatters.STATUS_UPDATED, diff[4].Status)
		require.Equal(t, "old", diff[4].OldVal)
		require.Equal(t, "new", diff[4].Val)
	})

	t.Run("deeply nested maps", func(t *testing.T) {
		first := map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": "old",
				},
			},
		}
		second := map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": "new",
				},
			},
		}

		diff := diffCalc(first, second)
		require.Len(t, diff, 1)
		require.Equal(t, "a", diff[0].Name)

		sub1, ok := diff[0].Val.([]formatters.DiffTree)
		require.True(t, ok)
		require.Len(t, sub1, 1)

		sub2, ok := sub1[0].Val.([]formatters.DiffTree)
		require.True(t, ok)
		require.Len(t, sub2, 1)

		require.Equal(t, formatters.STATUS_UPDATED, sub2[0].Status)
		require.Equal(t, "old", sub2[0].OldVal)
		require.Equal(t, "new", sub2[0].Val)
	})
}
