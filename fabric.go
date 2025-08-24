package code

import (
	"fmt"
)

type ParserFabricInterface interface {
	getByFileExtension(ext string) (ParserInterface, error)
}

type ParserFabric struct{}

func (f ParserFabric) getByFileExtension(ext string) (ParserInterface, error) {
	switch ext {
	case "json":
		return JsonParser{}, nil
	case "yml":
		return YmlParser{}, nil
	default:
		return DummyParser{}, fmt.Errorf("unknow extension %s", ext)
	}
}
