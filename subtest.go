package subtest

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type Subtest struct {
	Name         string
	ExpectedData interface{}
	ExpectedErr  error
	Test         func() (interface{}, error)
	Setup        func()
	Teardown     func()
}

func RunSubtests(t *testing.T, subtests []Subtest) {
	for _, subtest := range subtests {
		t.Run(subtest.Name, func(t *testing.T) {
			if subtest.Setup != nil {
				subtest.Setup()
			}

			result, err := subtest.Test()
			valueType := reflect.TypeOf(result)

			if valueType != nil && (valueType.Kind() == reflect.Ptr || valueType.Kind() == reflect.Slice) {
				if valueType.Kind() == reflect.Slice &&
					reflect.ValueOf(result).Len() != 0 &&
					reflect.ValueOf(subtest.ExpectedData).Len() != 0 {
					if !reflect.DeepEqual(result, subtest.ExpectedData) {
						t.Errorf("expected (%v), got (%v)", subtest.ExpectedData, result)
					}
				}
			} else if valueType != nil && (valueType.Kind() == reflect.Struct) {
				if !reflect.DeepEqual(result, subtest.ExpectedData) {
					t.Errorf("expected (%v), got (%v)", subtest.ExpectedData, result)
				}
			} else if valueType != nil && valueType.Kind() == reflect.Map {
				if fmt.Sprint(result) != fmt.Sprint(subtest.ExpectedData) {
					t.Errorf("expected (%v), got (%v)", subtest.ExpectedData, result)
				}
			} else {
				if result != subtest.ExpectedData {
					t.Errorf("expected (%v), got (%v)", subtest.ExpectedData, result)
				}
			}

			if err != nil && subtest.ExpectedErr != nil && err.Error() != subtest.ExpectedErr.Error() {
				t.Errorf("expected error (%v), got error (%v)", subtest.ExpectedErr, err)
			}

			if (err != nil && subtest.ExpectedErr == nil) || (err == nil && subtest.ExpectedErr != nil) {
				t.Errorf("expected error (%v), got error (%v)", subtest.ExpectedErr, err)
			}

			t.Cleanup(func() {
				os.Clearenv()

				if subtest.Teardown != nil {
					subtest.Teardown()
				}
			})
		})
	}
}
