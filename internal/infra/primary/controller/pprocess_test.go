package controller

import (
	"github.com/rendis/lightms"
	"reflect"
	"testing"
)

func TestGetControllerInstance(t *testing.T) {
	type args struct {
		c []ControllerRunnable
	}
	tests := []struct {
		name string
		args args
		want lightms.PrimaryProcess
	}{
		{
			name: "TestGetControllerInstance",
			args: args{
				c: []ControllerRunnable{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetControllerInstance(tt.args.c); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetControllerInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetControllerInstance1(t *testing.T) {
	type args struct {
		c []ControllerRunnable
	}
	tests := []struct {
		name string
		args args
		want lightms.PrimaryProcess
	}{
		{
			name: "TestGetControllerInstance1",
			args: args{
				c: []ControllerRunnable{},
			},
			want: &GinController{[]ControllerRunnable{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetControllerInstance(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetControllerInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
