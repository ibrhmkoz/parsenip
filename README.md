## Parsenip

Parsenip is a Go package that provides a function to parse a target string into a slice of maps containing key-value pairs. The function `Parse` takes two strings: a format string and a target string. The format string contains placeholders with labels that the `Parse` function will use to extract values from the target string and store them in a map with the corresponding labels as keys.

## Motivation
The motivation behind Parsenip was to bring the type-level pattern matching capabilities found in TypeScript and the parsing functionality available in Python's Parse library to the Go programming language. By combining the strengths of both, we created a versatile and robust parsing tool that makes it easier to extract structured data from strings in Go applications.

## Installation

To install the parsenip package, run the following command:

`go get github.com/ibrhmkoz/parsenip`

## Usage

Here's an example of how to use the `Parse` function:

```go
func TestParse(t *testing.T) {
	format := `
Name: {:s:Name}
Surname: {:s:Surname}
Age: {:d:Age}
Colors: {:a:Colors}
{:i}
Weight: {:f:Weight}
Scores: {:ad:Scores}
Grades: {:af:Grades}`

	target := `


Name: John
Surname: Wayne
Age: 30
Colors: red, blue, green
        
Foo: Boo
Goo: Coo

Weight: 75.5
Scores: 90, 80, 85
Grades: 3.6, 3.7, 4.0

Name: James
Surname: Johnson
Age: 37
Colors: green, blue, green
        
Roo: Loo
Hoo: Poo

Weight: 80.5
Scores: 50, 80, 85
Grades: 3.6, 3.7, 4.0

`

	expected := []map[string]interface{}{
		{
			"Name":    "John",
			"Surname": "Wayne",
			"Age":     30,
			"Colors":  []string{"red", "blue", "green"},
			"Weight":  75.5,
			"Scores":  []int{90, 80, 85},
			"Grades":  []float64{3.6, 3.7, 4.0},
		},
		{
			"Name":    "James",
			"Surname": "Johnson",
			"Age":     37,
			"Colors":  []string{"green", "blue", "green"},
			"Weight":  80.5,
			"Scores":  []int{50, 80, 85},
			"Grades":  []float64{3.6, 3.7, 4.0},
		},
	}

	result, err := Parse(format, target)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
```

The `Parse` function will return a slice of maps containing the parsed values. Each map in the slice corresponds to a section of the target string that matches the format string.

## Placeholders

The format string can contain the following placeholders:

- `{:s:label}`: Matches any string and stores it in the map with the given label.
- `{:d:label}`: Matches any integer and stores it in the map with the given label.
- `{:f:label}`: Matches any floating-point number and stores it in the map with the given label.
- `{:ad:label}`: Matches a comma-separated list of integers and stores them in the map with the given label as a slice of integers.
- `{:af:label}`: Matches a comma-separated list of floating-point numbers and stores them in the map with the given label as a slice of floats.
- `{:a:label}`: Matches a comma-separated list of strings and stores them in the map with the given label as a slice of strings.
- `{:i}`: Ignores any text until the next placeholder or the end of the string.
- `{:e}`: Matches any amount of whitespace.
