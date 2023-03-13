package properties

import (
	"reflect"
	"testing"
)

func TestGetDatabaseProp(t *testing.T) {
	result := GetDatabaseProp()

	if reflect.TypeOf(result) != reflect.TypeOf(&DatabaseProp{}) {
		t.Errorf("GetDatabaseProp Expected Type: [%T], Result: [%T]", &DatabaseProp{}, result)
	}
}
