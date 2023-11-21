package example

import (
	"errors"
	"testing"

	"github.com/0xanonymeow/go-subtest"
)

func TestExample(t *testing.T) {
	subtests := []subtest.Subtest{
		{
			Name:         "example",
			ExpectedData: "go-subtest",
			ExpectedErr:  nil,
			Test: func() (interface{}, error) {
				s, err := Print("go-subtest")

				return s, err
			},
		},
		{
			Name:         "example_error",
			ExpectedData: "",
			ExpectedErr:  errors.New("error expected"),
			Test: func() (interface{}, error) {
				s, err := Print("error")

				return s, err
			},
		},
	}

	subtest.RunSubtests(t, subtests)
}
