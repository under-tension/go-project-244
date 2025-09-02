package parsers

import yaml "gopkg.in/yaml.v3"

type YmlParser struct{}

func (p YmlParser) Parse(content string) (map[string]interface{}, error) {
	var res map[string]interface{}

	err := yaml.Unmarshal([]byte(content), &res)

	if err != nil {
		return res, err
	}

	return res, nil
}
