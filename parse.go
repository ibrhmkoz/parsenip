package parsenip

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrNoMatch = errors.New("no match")

func Parse(format, target string) ([]map[string]interface{}, error) {
	// Temporarily replace escaped braces with placeholder strings
	format = strings.ReplaceAll(format, `\\{`, `{{OPEN_BRACE}}`)
	format = strings.ReplaceAll(format, `\\}`, `{{CLOSE_BRACE}}`)

	// Escape regex special characters, excluding '{' and '}'
	var specialCharacters = []string{`.`, `^`, `$`, `*`, `+`, `?`, `(`, `)`, `[`, `]`, `|`}
	for _, char := range specialCharacters {
		format = strings.ReplaceAll(format, char, "\\"+char)
	}

	// Process custom tags as before
	format = regexp.MustCompile(`{:s:(\w+)}`).ReplaceAllString(format, `(?P<${1}_s>.+?)`)
	format = regexp.MustCompile(`{:d:(\w+)}`).ReplaceAllString(format, `(?P<${1}_d>\d+?)`)
	format = regexp.MustCompile(`{:f:(\w+)}`).ReplaceAllString(format, `(?P<${1}_f>\d+(\.\d+)?)`)
	format = regexp.MustCompile(`{:ad:(\w+)}`).ReplaceAllString(format, `(?P<${1}_ad>(\d+, )*\d+?)`)
	format = regexp.MustCompile(`{:af:(\w+)}`).ReplaceAllString(format, `(?P<${1}_af>(\d+(\.\d+), )*\d+(\.\d+))?`)
	format = regexp.MustCompile(`{:a:(\w+)}`).ReplaceAllString(format, `(?P<${1}_a>.+)`)
	format = regexp.MustCompile(`{:i}`).ReplaceAllString(format, `(?:.|\n)*?`)
	format = regexp.MustCompile(`{:e}`).ReplaceAllString(format, `\s*`)
	format = regexp.MustCompile(`{:we}`).ReplaceAllString(format, `(?:$|\s|\t|\n)+?`)

	// Revert the temporary placeholders back to their original form
	format = strings.ReplaceAll(format, `{{OPEN_BRACE}}`, `\{`)
	format = strings.ReplaceAll(format, `{{CLOSE_BRACE}}`, `\}`)

	re := regexp.MustCompile(format)
	matches := re.FindAllStringSubmatch(target, -1)

	if matches == nil {
		return nil, ErrNoMatch
	}

	var results []map[string]interface{}
	for _, match := range matches {
		result := make(map[string]interface{})
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				value := match[i]
				processParsedValue(value, name, result)
			}
		}
		results = append(results, result)
	}

	return results, nil
}

// Rest of the code remains the same...
func processParsedValue(value, name string, result map[string]interface{}) {
	if strings.HasSuffix(name, "_d") {
		name = name[:len(name)-2]
		result[name], _ = strconv.Atoi(value)
	} else if strings.HasSuffix(name, "_f") {
		name = name[:len(name)-2]
		result[name], _ = strconv.ParseFloat(value, 64)
	} else if strings.HasSuffix(name, "_ad") {
		name = name[:len(name)-3]
		split := strings.Split(value, ", ")
		intArray := make([]int, len(split))
		for i, s := range split {
			intArray[i], _ = strconv.Atoi(s)
		}
		result[name] = intArray
	} else if strings.HasSuffix(name, "_af") {
		name = name[:len(name)-3]
		split := strings.Split(value, ", ")
		floatArray := make([]float64, len(split))
		for i, s := range split {
			floatArray[i], _ = strconv.ParseFloat(s, 64)
		}
		result[name] = floatArray
	} else if strings.HasSuffix(name, "_a") {
		name = name[:len(name)-2]
		result[name] = strings.Split(value, ", ")
	} else {
		name = name[:len(name)-2]
		result[name] = value
	}
}
