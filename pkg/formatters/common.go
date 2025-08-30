package formatters

type DiffTree struct {
	Name   string
	Val    interface{}
	Prefix string
}

type FormatterInterface interface {
	Format(diffTree []DiffTree) (string, error)
}

type Settings struct {
	StartIdents int
}

type BaseFormatter struct {
	settings Settings
}

func IsDiffTreeSlice(v interface{}) bool {
	if v == nil {
		return false
	}

	_, ok := v.([]DiffTree)

	return ok
}
