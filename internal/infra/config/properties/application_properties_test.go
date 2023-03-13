package properties

import (
	"reflect"
	"testing"
)

func TestGetApplicationProperties(t *testing.T) {
	tests := []struct {
		name string
		want *ApplicationProperties
	}{
		{
			name: "Test_GetApplicationProperties",
			want: &ApplicationProperties{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetApplicationProperties(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}
