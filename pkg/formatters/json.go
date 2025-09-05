package formatters

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JsonFormatter struct {
	BaseFormatter
}

func (f JsonFormatter) Format(diffTree []DiffTree) (string, error) {
	f.settings = Settings{StartIdents: 0}
	return f.format(diffTree, f.settings.StartIdents)
}

func (f JsonFormatter) format(diff []DiffTree, depth int) (string, error) {
	var lines []string

	for _, node := range diff {

		if node.Type == TYPE_ROOT && IsDiffTreeSlice(node.Val) {
			subTreeResult, err := f.format(node.Val.([]DiffTree), depth+1)

			if err != nil {
				return "", err
			}

			line := fmt.Sprintf("\"%s\": %s", node.Name, subTreeResult)
			lines = append(lines, line)
		} else {
			oldValEncoded, _ := json.Marshal(node.OldVal)
			valEncoded, _ := json.Marshal(node.Val)

			switch node.Status {
			case STATUS_ADDED:
				lines = append(lines, fmt.Sprintf("\"+ %s\": %s", node.Name, string(valEncoded)))
			case STATUS_DELETED:
				lines = append(lines, fmt.Sprintf("\"- %s\": %s", node.Name, string(oldValEncoded)))
			case STATUS_NON_CHANGE:
				lines = append(lines, fmt.Sprintf("\"  %s\": %s", node.Name, string(valEncoded)))
			case STATUS_UPDATED:
				lines = append(lines, fmt.Sprintf("\"- %s\": %s", node.Name, string(oldValEncoded)))
				lines = append(lines, fmt.Sprintf("\"+ %s\": %s", node.Name, string(valEncoded)))
			}
		}
	}

	result := fmt.Sprintf("{\n%s\n}", strings.Join(lines, ",\n"))

	return result, nil
}
