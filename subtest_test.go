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
	sampleMap := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	subtests := []Subtest{
		{
			Name:         "struct_compare",
			ExpectedData: sampleStruct,
			ExpectedErr:  nil,
			Test: func() (interface{}, error) {
				compareStruct := SampleStruct{
					Name: "John",
					Age:  30,
				}

				return compareStruct, nil
			},
		},
		{
			Name:         "map_compare",
			ExpectedData: sampleMap,
			ExpectedErr:  nil,
			Test: func() (interface{}, error) {
				compareMap := map[string]interface{}{
					"name": "John",
					"age":  30,
				}

				return compareMap, nil
			},
		},
	}

	RunSubtests(t, subtests)
}
