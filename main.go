package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Parse(format, target string) map[string]interface{} {
	format = regexp.MustCompile(`{:s:(\w+)}`).ReplaceAllString(format, `(?P<${1}_s>\w+)`)
	format = regexp.MustCompile(`{:d:(\w+)}`).ReplaceAllString(format, `(?P<${1}_d>\d+)`)
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
	format := `Name: {:s:foo}, Surname: {:s:surname}, Age: {:d:age}, Colors: {:a:colors}`
	target := `Name: John, Surname: Wayne, Age: 30, Colors: red, blue, green`

	captures := Parse(format, target)
	fmt.Println(captures)
}
