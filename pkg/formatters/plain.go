package formatters

import (
	"fmt"
)

type PlainFormatter struct {
	BaseFormatter
}

func (f PlainFormatter) Format(diffTree []DiffTree) (string, error) {
	return f.format(diffTree, "")
}

func (f PlainFormatter) format(diff []DiffTree, path string) (string, error) {
	result := ""

	for _, node := range diff {

		if node.Type == TYPE_ROOT && IsDiffTreeSlice(node.Val) {
			currentPath := fmt.Sprintf("%s%s.", path, node.Name)
			subTreeResult, err := f.format(node.Val.([]DiffTree), currentPath)

			if err != nil {
				return "", err
			}

			result += subTreeResult
		} else {
			oldVal, _ := f.prepareValue(node.OldVal)
			val, _ := f.prepareValue(node.Val)

			currentPath := path + node.Name

			switch node.Status {
			case STATUS_ADDED:
				result += fmt.Sprintf("Property '%s' was added with value: %s\n", currentPath, val)
			case STATUS_DELETED:
				result += fmt.Sprintf("Property '%s' was removed\n", currentPath)
			case STATUS_UPDATED:
				result += fmt.Sprintf("Property '%s' was updated. From %s to %s\n", currentPath, oldVal, val)
			}
		}
	}

	return result, nil
}

func (f PlainFormatter) prepareValue(val any) (string, error) {
	result := ""

	switch val.(type) {
	case string:
		result = fmt.Sprintf("'%s'", val)
	case nil:
		result = "null"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
		float32, float64, complex64, complex128:
		result = fmt.Sprintf("%v", val)
	default:
		result = "[complex value]"
	}

	return result, nil
}
