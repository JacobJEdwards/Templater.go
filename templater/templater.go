package templater

import (
	"errors"
	"regexp"
	"strings"
)

var matchRegex = regexp.MustCompile(`{{\s*([^}]+)\s*}}`)

func parse(template string) ([]string, error) {
	ast, err := parseTemplate(template)

	if err != nil {
		return nil, err
	}

	return ast, nil
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

func compileToString(template string, data map[string]string) (string, error) {
	ast, err := parse(template)
    if err != nil {
        return "", err
    }
	var resultStr strings.Builder
	resultStr.Grow(len(template))

	for _, item := range ast {
        if ok := matchRegex.MatchString(item); ok {
			key := strings.ReplaceAll(item, "{{", "")
			key = strings.ReplaceAll(key, "}}", "")
			key = strings.TrimSpace(key)

			replacement, ok := data[key]

			if ok {
				resultStr.WriteString(replacement)
			} else {
				resultStr.WriteString(item)
			}

		} else {
			resultStr.WriteString(item)
		}
	}

	return resultStr.String(), nil
}

func Compile(template string, data map[string]string) (string, error) {
	return compileToString(template, data)
}
