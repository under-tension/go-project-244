package formatters

import (
	"fmt"
	"sort"
	"strings"
)

type StylishFormatter struct {
	BaseFormatter
}

func (f StylishFormatter) Format(diffTree []DiffTree) (string, error) {
	f.settings = Settings{StartIdents: 0}
	return f.format(diffTree, f.settings.StartIdents)
}

func (f StylishFormatter) format(diff []DiffTree, depth int) (string, error) {
	var lines []string

	for _, node := range diff {

		if node.Type == TYPE_ROOT && IsDiffTreeSlice(node.Val) {
			subTreeResult, err := f.format(node.Val.([]DiffTree), depth+1)

			if err != nil {
				return "", err
			}

			line := strings.Repeat("    ", depth+1) + node.Name + ": " + subTreeResult
			lines = append(lines, line)
		} else {
			oldValEncoded := renderValue(node.OldVal, depth)
			valEncoded := renderValue(node.Val, depth)

			space := strings.Repeat(" ", ((depth+1)*4)-2)

			switch node.Status {
			case STATUS_ADDED:
				lines = append(lines, space+fmt.Sprintf("+ %s: %s", node.Name, valEncoded))
			case STATUS_DELETED:
				lines = append(lines, space+fmt.Sprintf("- %s: %s", node.Name, oldValEncoded))
			case STATUS_NON_CHANGE:
				lines = append(lines, space+fmt.Sprintf("  %s: %s", node.Name, valEncoded))
			case STATUS_UPDATED:
				lines = append(lines, space+fmt.Sprintf("- %s: %s", node.Name, oldValEncoded))
				lines = append(lines, space+fmt.Sprintf("+ %s: %s", node.Name, valEncoded))
			}
		}

	}

	end := strings.Repeat("    ", depth) + "}"

	result := fmt.Sprintf("{\n%s\n%s", strings.Join(lines, "\n"), end)

	return result, nil
}

func renderValue(val any, depth int) string {
	switch val := val.(type) {
	case map[string]any:
		return renderMap(val, depth+1)
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", val)
	}
}

func renderMap(val map[string]any, depth int) string {
	keys := make([]string, 0)

	for key := range val {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool {
		return keys[i] < keys[j]
	})

	var lines []string

	for _, key := range keys {
		value := val[key]
		line := strings.Repeat("    ", depth+1) + fmt.Sprintf("%s: %s", key, renderValue(value, depth))
		lines = append(lines, line)
	}

	end := strings.Repeat("    ", depth) + "}"

	result := fmt.Sprintf("{\n%s\n%s", strings.Join(lines, "\n"), end)

	return result
}
