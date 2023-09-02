package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Parse(format, target string) (map[string]interface{}, error) {
	// Convert special syntax into regular expressions
	format = regexp.MustCompile(`{:s:(\w+)}`).ReplaceAllString(format, `(?P<${1}_s>\w+)`)
	format = regexp.MustCompile(`{:d:(\w+)}`).ReplaceAllString(format, `(?P<${1}_d>\d+)`)
	format = regexp.MustCompile(`{:f:(\w+)}`).ReplaceAllString(format, `(?P<${1}_f>\d+(\.\d+))`)
	format = regexp.MustCompile(`{:ad:(\w+)}`).ReplaceAllString(format, `(?P<${1}_ad>(\d+, )*\d+)`)
	format = regexp.MustCompile(`{:af:(\w+)}`).ReplaceAllString(format, `(?P<${1}_af>(\d+(\.\d+), )*\d+(\.\d+))`)
	format = regexp.MustCompile(`{:a:(\w+)}`).ReplaceAllString(format, `(?P<${1}_a>(\w+, )*\w+)`)
	//format = regexp.MustCompile(`{:X}`).ReplaceAllString(format, `(?:.*\n)*.*`)
	//format = regexp.MustCompile(`{:X}`).ReplaceAllString(format, `(?:.*\n?)*`)
	format = regexp.MustCompile(`{:i}`).ReplaceAllString(format, `(?:.|\n)*`)

	re := regexp.MustCompile(format)
	match := re.FindStringSubmatch(target)

	if match == nil {
		return nil, fmt.Errorf("no match")
	}

	result := make(map[string]interface{})
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			value := match[i]
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
	}

	return result, nil
}
