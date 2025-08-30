package parsers

import "encoding/json"

type JsonParser struct{}

func (p JsonParser) Parse(content string) (map[string]interface{}, error) {
	var res map[string]interface{}

	err := json.Unmarshal([]byte(content), &res)

	if err != nil {
		return res, err
	}

	return res, nil
}
