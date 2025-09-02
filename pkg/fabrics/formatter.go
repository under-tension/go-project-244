package fabrics

import (
	"code/pkg/formatters"
	"fmt"
)

type FormatterFabricInterface interface {
	GetFormatterByStr(key string) (formatters.FormatterInterface, error)
}

type FormatterFabric struct{}

func (f FormatterFabric) GetFormatterByStr(key string) (formatters.FormatterInterface, error) {
	switch key {
	case "stylish":
		return formatters.DefaultFormatter{}, nil
	case "plain":
		return formatters.PlainFormatter{}, nil
	default:
		return formatters.DummyFormatter{}, fmt.Errorf("unknow %s formatter", key)
	}
}
