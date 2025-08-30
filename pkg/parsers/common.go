package parsers

type ParserInterface interface {
	Parse(content string) (map[string]interface{}, error)
}
