package formatters

import (
	"encoding/json"
	"strings"
)

type DefaultFormatter struct {
	BaseFormatter
}

func (f DefaultFormatter) Format(diffTree []DiffTree) (string, error) {
	return f.format(diffTree, 0)
}

func (f DefaultFormatter) format(diff []DiffTree, depth int) (string, error) {
	result := ""
	result += strings.Repeat(" ", depth) + "{\n"

	for _, node := range diff {

		if IsDiffTreeSlice(node.Val) {
			subTreeResult, err := f.format(node.Val.([]DiffTree), depth+1)

			if err != nil {
				return "", err
			}

			result += strings.Repeat("\t", depth) + node.Prefix + node.Name + ":" + subTreeResult
		} else {
			valEncoded, _ := json.Marshal(node.Val)
			result += strings.Repeat("\t", depth) + node.Prefix + node.Name + ": " + string(valEncoded) + "\n"
		}

	}

	result += strings.Repeat("\t", depth) + "}\n"

	return result, nil
}
