package templater

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func parse(template string) []string {
	ast, err := parseTemplate(template)

	if err != nil {
		panic(err)
	}

	return ast
}

func parseTemplate(template string) ([]string, error) {
	var arr []string
	var startPos int = 0

	for {
		startMarker := strings.Index(template, "{{")

		if startMarker == -1 {
			arr = append(arr, template)
			break
		}

		endMarker := strings.Index(template, "}}")

		if endMarker == -1 {
			return nil, errors.New("Start marker not matched")
		}

		if startMarker != startPos {
			arr = append(arr, template[startPos:startMarker])
		}

		arr = append(arr, template[startMarker:endMarker+2])
		template = template[endMarker+2:]
		startPos = endMarker + 2
	}

	return arr, nil
}

func compileToString(template string, data map[string]string) string {
	ast := parse(template)
	var resultArr []string

	for i, item := range ast {

		if ok, _ := regexp.MatchString(`{{\s*([^}]+)\s*}}`, item); ok {

			key := strings.ReplaceAll(ast[i], "{{", "")
			key = strings.ReplaceAll(key, "}}", "")
			key = strings.TrimSpace(key)

			replacement, ok := data[key]

			if ok {
				resultArr = append(resultArr, replacement)
			} else {
				resultArr = append(resultArr, item)
			}

		} else {
			fmt.Println(item)

			resultArr = append(resultArr, item)
		}
	}

	joinedArr := strings.Join(resultArr, "")

	return joinedArr
}

func Compile(template string) string {
	data := make(map[string]string)
	data["test"] = "success"

	return compileToString(template, data)
}
