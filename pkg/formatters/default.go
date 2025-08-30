package formatters

import (
	"encoding/json"
	"strings"
)

type DefaultFormatter struct {
	BaseFormatter
}

func (f DefaultFormatter) Format(diffTree []DiffTree) (string, error) {
	f.settings = Settings{StartIdents: 0}
	return f.format(diffTree, f.settings.StartIdents)
}

func (f DefaultFormatter) format(diff []DiffTree, depth int) (string, error) {
	result := ""
	result += "{\n"

	for _, node := range diff {

		if IsDiffTreeSlice(node.Val) {
			subTreeResult, err := f.format(node.Val.([]DiffTree), depth+1)

			if err != nil {
				return "", err
			}

			result += strings.Repeat("\t", depth+1) + node.Prefix + node.Name + ": " + subTreeResult
		} else {
			valEncoded, _ := json.Marshal(node.Val)
			result += strings.Repeat("\t", depth+1) + node.Prefix + node.Name + ": " + string(valEncoded) + "\n"
		}

	}

	result += strings.Repeat("\t", depth) + "}\n"

	return result, nil
}
