package code

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
)

type diffAst struct {
	name   string
	val    interface{}
	prefix string
}

func ParseFile(path string) (map[string]interface{}, error) {
	emptyRes := make(map[string]interface{}, 0)
	var res map[string]interface{}

	content, err := os.ReadFile(path)

	if err != nil {
		return emptyRes, err
	}

	fileExt := filepath.Ext(path)
	fileExt = fileExt[1:]

	fabric := ParserFabric{}
	parser, err := fabric.getByFileExtension(fileExt)

	if err != nil {
		return emptyRes, err
	}

	res, err = parser.parse(string(content))

	if err != nil {
		return emptyRes, err
	}

	return res, nil
}

func GenDiff(first, second map[string]interface{}) string {
	diff := []diffAst{}

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

		if exist && !exist2 {
			node := diffAst{}
			node.name = key
			node.val = val
			node.prefix = "- "

			diff = append(diff, node)
			continue
		} else if !exist && exist2 {
			node := diffAst{}
			node.name = key
			node.val = val2
			node.prefix = "+ "

			diff = append(diff, node)
			continue
		} else if val != val2 {
			node := diffAst{}
			node.name = key
			node.val = val
			node.prefix = "- "

			node2 := diffAst{}
			node2.name = key
			node2.val = val2
			node2.prefix = "+ "

			diff = append(diff, node)
			diff = append(diff, node2)

			continue
		} else {
			node := diffAst{}
			node.name = key
			node.val = val
			node.prefix = "  "

			diff = append(diff, node)
			continue
		}
	}

	sort.Slice(diff, func(i int, j int) bool {
		return diff[i].name < diff[j].name
	})

	result := ""
	result += "{\n"

	for _, node := range diff {
		valEncoded, _ := json.Marshal(node.val)

		result += "	" + node.prefix + node.name + ": " + string(valEncoded) + "\n"

	}

	result += "}\n"

	return string(result)
}
