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

	content, err := os.ReadFile(filepath.Clean(path))

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

func GenDiff(firstFile, secondFile string, format string) (string, error) {
	firstFileMap, err := ParseFile(firstFile)

	if err != nil {
		return "", err
	}

	secondFileMap, err := ParseFile(secondFile)

	if err != nil {
		return "", err
	}

	diff := diffCalc(firstFileMap, secondFileMap)

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

func diffCalc(first, second map[string]interface{}) []formatters.DiffTree {
	diff := []formatters.DiffTree{}

	keys := map[string]string{}

	for key := range first {
		keys[key] = key
	}

	for key := range second {
		keys[key] = key
	}

	for key := range keys {
		node := processDiff(key, first, second)
		diff = append(diff, node)
	}

	sort.Slice(diff, func(i int, j int) bool {
		return diff[i].Name < diff[j].Name
	})

	return diff
}

func processDiff(key string, first, second map[string]interface{}) formatters.DiffTree {
	val, exist := first[key]
	val2, exist2 := second[key]

	var node formatters.DiffTree
	node.Name = key

	if isMap(val) && isMap(val2) {
		node.Type = formatters.TYPE_ROOT
		node.Status = formatters.STATUS_NON_CHANGE
		subDiff := diffCalc(val.(map[string]interface{}), val2.(map[string]interface{}))
		node.Val = subDiff
		return node
	}

	if !exist {
		node.Type = formatters.TYPE_FINAL
		node.Val = val2
		node.Status = formatters.STATUS_ADDED
	} else if !exist2 {
		node.Type = formatters.TYPE_FINAL
		node.OldVal = val
		node.Status = formatters.STATUS_DELETED
	} else if val != val2 {
		node.Type = formatters.TYPE_FINAL
		node.OldVal = val
		node.Val = val2
		node.Status = formatters.STATUS_UPDATED
	} else {
		node.Type = formatters.TYPE_FINAL
		node.Val = val
		node.Status = formatters.STATUS_NON_CHANGE
	}

	return node
}

func isMap(v interface{}) bool {
	if v == nil {
		return false
	}
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Map
}
