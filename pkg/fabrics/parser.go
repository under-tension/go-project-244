package fabrics

import (
	"code/pkg/parsers"
	"fmt"
)

type ParserFabricInterface interface {
	GetByFileExtension(ext string) (parsers.ParserInterface, error)
}

type ParserFabric struct{}

func (f ParserFabric) GetByFileExtension(ext string) (parsers.ParserInterface, error) {
	switch ext {
	case "json":
		return parsers.JsonParser{}, nil
	case "yml":
		return parsers.YmlParser{}, nil
	case "yaml":
		return parsers.YmlParser{}, nil
	default:
		return parsers.DummyParser{}, fmt.Errorf("unknow extension %s", ext)
	}
}
