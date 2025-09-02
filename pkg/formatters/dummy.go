package formatters

type DummyFormatter struct {
	BaseFormatter
}

func (f DummyFormatter) Format(diffTree []DiffTree) (string, error) {
	return "", nil
}
