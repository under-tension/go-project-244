package parsers

type DummyParser struct{}

func (p DummyParser) Parse(content string) (map[string]interface{}, error) {
	var res map[string]interface{}
	return res, nil
}
