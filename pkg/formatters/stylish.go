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
	result := ""
	result += "{\n"

	for _, node := range diff {

		if node.Type == TYPE_ROOT && IsDiffTreeSlice(node.Val) {
			subTreeResult, err := f.format(node.Val.([]DiffTree), depth+1)

			if err != nil {
				return "", err
			}

			result += strings.Repeat("    ", depth+1) + node.Name + ": " + subTreeResult
		} else {
			oldValEncoded := renderValue(node.OldVal, depth)
			valEncoded := renderValue(node.Val, depth)

			space := strings.Repeat(" ", ((depth+1)*4)-2)
			result += space

			switch node.Status {
			case STATUS_ADDED:
				result += "+ " + node.Name + ": " + valEncoded
			case STATUS_DELETED:
				result += "- " + node.Name + ": " + oldValEncoded
			case STATUS_NON_CHANGE:
				result += "  " + node.Name + ": " + valEncoded
			case STATUS_UPDATED:
				result += "- " + node.Name + ": " + oldValEncoded
				result += space
				result += "+ " + node.Name + ": " + valEncoded
			}
		}

	}

	result += strings.Repeat("    ", depth) + "}"

	if depth > 0 {
		result += "\n"
	}

	return result, nil
}

func renderValue(val any, depth int) string {
	switch val := val.(type) {
	case map[string]any:
		return renderMap(val, depth+1)
	case nil:
		return "null\n"
	default:
		return fmt.Sprintf("%v\n", val)
	}
}

func renderMap(val map[string]any, depth int) string {
	keys := make([]string, 0)

	for key := range val {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool {
		return i < j
	})

	result := ""
	result += "{\n"

	for _, key := range keys {
		value := val[key]
		result += strings.Repeat("    ", depth+1) + fmt.Sprintf("%s: %s", key, renderValue(value, depth))
	}

	result += strings.Repeat("    ", depth) + "}\n"

	return result
}
