package main

import (
	"net/http"
	"testing"
)

func Test_setupListener(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupListener()
		})
	}
}

func TestPostHandler(t *testing.T) {
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
			PostHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_generateMessage(t *testing.T) {
	type args struct {
		data        GiteaPostData
		eventHeader string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMessage(tt.args.data, tt.args.eventHeader); got != tt.want {
				t.Errorf("generateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
