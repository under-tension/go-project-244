package code

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v3"
)

type ParserInterface interface {
	parse(content string) (map[string]interface{}, error)
}

type JsonParser struct{}

func (p JsonParser) parse(content string) (map[string]interface{}, error) {
	var res map[string]interface{}

	err := json.Unmarshal([]byte(content), &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

type YmlParser struct{}

func (p YmlParser) parse(content string) (map[string]interface{}, error) {
	var res map[string]interface{}

	err := yaml.Unmarshal([]byte(content), &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

type DummyParser struct{}

func (p DummyParser) parse(content string) (map[string]interface{}, error) {
	var res map[string]interface{}
	return res, nil
}
