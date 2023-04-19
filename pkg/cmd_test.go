package cmd

import (
	"net/http"
	"reflect"
	"testing"
)

// Empty tests - just form demo , run in gitlab ci
func TestIndexHandler(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		mydata map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IndexHandler(tt.args.w, tt.args.r, tt.args.mydata)
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Login(tt.args.w, tt.args.r)
		})
	}
}

func TestLoadCfg(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadCfg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadCfg() = %v, want %v", got, tt.want)
			}
		})
	}
}
