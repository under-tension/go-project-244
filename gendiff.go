package code

import (
	"encoding/json"
	"os"
)

func ParseFile(path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(path)

	var res map[string]interface{}

	if err != nil {
		return res, err
	}

	json.Unmarshal(content, &res)

	return res, nil
}
