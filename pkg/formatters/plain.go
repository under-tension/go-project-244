package formatters

import (
	"fmt"
	"strings"
)

type PlainFormatter struct {
	BaseFormatter
}

func (f PlainFormatter) Format(diffTree []DiffTree) (string, error) {
	return f.format(diffTree, "")
}

func (f PlainFormatter) format(diff []DiffTree, path string) (string, error) {
	var lines []string

	for _, node := range diff {

		if node.Type == TYPE_ROOT && IsDiffTreeSlice(node.Val) {
			currentPath := fmt.Sprintf("%s%s.", path, node.Name)
			subTreeResult, err := f.format(node.Val.([]DiffTree), currentPath)

			if err != nil {
				return "", err
			}

			if subTreeResult != "" {
				lines = append(lines, subTreeResult)
			}
		} else {
			oldVal, _ := f.prepareValue(node.OldVal)
			val, _ := f.prepareValue(node.Val)

			currentPath := path + node.Name

			var line string
			switch node.Status {
			case STATUS_ADDED:
				line = fmt.Sprintf("Property '%s' was added with value: %s", currentPath, val)
			case STATUS_DELETED:
				line = fmt.Sprintf("Property '%s' was removed", currentPath)
			case STATUS_UPDATED:
				line = fmt.Sprintf("Property '%s' was updated. From %s to %s", currentPath, oldVal, val)
			}

			if line != "" {
				lines = append(lines, line)
			}
		}
	}

	return strings.Join(lines, "\n"), nil
}

func (f PlainFormatter) prepareValue(val any) (string, error) {
	result := ""

	switch val.(type) {
	case string:
		result = fmt.Sprintf("'%s'", val)
	case nil:
		result = "null"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
		float32, float64, complex64, complex128, bool:
		result = fmt.Sprintf("%v", val)
	default:
		result = "[complex value]"
	}

	return result, nil
}
