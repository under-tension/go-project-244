package formatters

const (
	STATUS_ADDED = iota
	STATUS_DELETED
	STATUS_UPDATED
	STATUS_NON_CHANGE
)

const (
	TYPE_ROOT = iota
	TYPE_FINAL
)

type DiffTree struct {
	Name   string
	Type   int
	Status int
	OldVal any
	Val    any
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
