package code

import (
	"code/pkg/fabrics"
	"code/pkg/formatters"
	"os"
	"path/filepath"
	"reflect"
	"sort"
)

func ParseFile(path string) (map[string]interface{}, error) {
	emptyRes := make(map[string]interface{}, 0)
	var res map[string]interface{}

	content, err := os.ReadFile(path)

	if err != nil {
		return emptyRes, err
	}

	fileExt := filepath.Ext(path)
	fileExt = fileExt[1:]

	fabric := fabrics.ParserFabric{}
	parser, err := fabric.GetByFileExtension(fileExt)

	if err != nil {
		return emptyRes, err
	}

	res, err = parser.Parse(string(content))

	if err != nil {
		return emptyRes, err
	}

	return res, nil
}

func GenDiff(first, second map[string]interface{}, format string) (string, error) {
	diff := DiffCalc(first, second)

	formatterFabric := fabrics.FormatterFabric{}
	formatter, err := formatterFabric.GetFormatterByStr(format)

	if err != nil {
		return "", err
	}

	result, err := formatter.Format(diff)

	if err != nil {
		return "", err
	}

	return result, nil
}

func DiffCalc(first, second map[string]interface{}) []formatters.DiffTree {
	diff := []formatters.DiffTree{}

	keys := map[string]string{}

	for key := range first {
		keys[key] = key
	}

	for key := range second {
		keys[key] = key
	}

	for key := range keys {
		val, exist := first[key]
		val2, exist2 := second[key]

		if isMap(val) && isMap(val2) {
			node := formatters.DiffTree{}
			node.Name = key
			node.Prefix = "  "

			subDiff := []formatters.DiffTree{}
			subDiff = append(subDiff, DiffCalc(val.(map[string]interface{}), val2.(map[string]interface{}))...)

			node.Val = subDiff

			diff = append(diff, node)

			continue
		}

		if exist && !exist2 {
			node := formatters.DiffTree{}
			node.Name = key
			node.Val = val
			node.Prefix = "- "

			diff = append(diff, node)
			continue
		} else if !exist && exist2 {
			node := formatters.DiffTree{}
			node.Name = key
			node.Val = val2
			node.Prefix = "+ "

			diff = append(diff, node)
			continue
		} else if val != val2 {
			node := formatters.DiffTree{}
			node.Name = key
			node.Val = val
			node.Prefix = "- "

			node2 := formatters.DiffTree{}
			node2.Name = key
			node2.Val = val2
			node2.Prefix = "+ "

			diff = append(diff, node)
			diff = append(diff, node2)

			continue
		} else {
			node := formatters.DiffTree{}
			node.Name = key
			node.Val = val
			node.Prefix = "  "

			diff = append(diff, node)
			continue
		}
	}

	sort.Slice(diff, func(i int, j int) bool {
		return diff[i].Name < diff[j].Name
	})

	return diff
}

func isMap(v interface{}) bool {
	if v == nil {
		return false
	}
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Map
}
