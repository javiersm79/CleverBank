package controller

import (
	"github.com/rendis/lightms"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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

func Test_configuration(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	tests := []struct {
		name string
		args args
		want *http.Server
	}{
		{
			name: "Test_configuration",
			args: args{
				router: nil,
			},
			want: &http.Server{
				Addr:         ":0",
				Handler:      nil,
				ReadTimeout:  0 * time.Second,
				WriteTimeout: 0 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configuration(tt.args.router); !reflect.DeepEqual(got.Addr, tt.want.Addr) {
				t.Errorf("configuration() = %v, want %v", got, tt.want)
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
