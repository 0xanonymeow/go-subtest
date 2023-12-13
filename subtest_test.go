package subtest

import (
	"testing"
)

type SampleStruct struct {
	Name string
	Age  int
}

func TestSubtest(t *testing.T) {
	sampleStruct := SampleStruct{
		Name: "John",
		Age:  30,
	}

	subtests := []Subtest{
		{
			Name:         "struct_compare",
			ExpectedData: sampleStruct,
			ExpectedErr:  nil,
			Test: func() (interface{}, error) {
				return sampleStruct, nil
			},
		},
	}

	RunSubtests(t, subtests)
}
