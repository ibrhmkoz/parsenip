package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	format := `Name: {:s:Name}
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

`

	expected := map[string]interface{}{
		"Name":    "John",
		"Surname": "Wayne",
		"Age":     30,
		"Colors":  []string{"red", "blue", "green"},
		"Weight":  75.5,
		"Scores":  []int{90, 80, 85},
		"Grades":  []float64{3.6, 3.7, 4.0},
	}

	result, err := Parse(format, target)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
