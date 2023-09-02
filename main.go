package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Parse parses the target string using the given format and returns a map of captures.
func Parse(format, target string) map[string]interface{} {
	// Convert special syntax into regular expressions
	format = regexp.MustCompile(`{:s:(\w+)}`).ReplaceAllString(format, `(?P<${1}_s>\w+)`)
	format = regexp.MustCompile(`{:d:(\w+)}`).ReplaceAllString(format, `(?P<${1}_d>\d+)`)
	format = regexp.MustCompile(`{:f:(\w+)}`).ReplaceAllString(format, `(?P<${1}_f>\d+(\.\d+)?)`)
	format = regexp.MustCompile(`{:ad:(\w+)}`).ReplaceAllString(format, `(?P<${1}_ad>(\d+, )*\d+)`)
	format = regexp.MustCompile(`{:af:(\w+)}`).ReplaceAllString(format, `(?P<${1}_af>(\d+(\.\d+)?, )*\d+(\.\d+)?)`)
	format = regexp.MustCompile(`{:a:(\w+)}`).ReplaceAllString(format, `(?P<${1}_a>(\w+, )*\w+)`)

	re := regexp.MustCompile(format)
	match := re.FindStringSubmatch(target)

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

	return result
}

func main() {
	format := `Name: {:s:name}, Surname: {:s:surname}, Age: {:d:age}, Colors: {:a:colors}, Weight: {:f:weight}, Scores: {:ad:scores}, Grades: {:af:grades}`
	target := `Name: John, Surname: Wayne, Age: 30, Colors: red, blue, green, Weight: 70.5, Scores: 10, 20, 30, Grades: 3.5, 4.0, 3.8`

	captures := Parse(format, target)
	fmt.Println(captures)
}
