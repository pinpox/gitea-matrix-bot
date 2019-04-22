package main

import (
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNewGiteaDB(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *GiteaDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGiteaDB(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGiteaDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiteaDB_Init(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbg := &GiteaDB{
				path: tt.fields.path,
			}
			dbg.Init()
		})
	}
}

func TestGiteaDB_GetToken(t *testing.T) {
	type fields struct {
		path string
	}
	type args struct {
		room string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbg := &GiteaDB{
				path: tt.fields.path,
			}
			if got := dbg.GetToken(tt.args.room); got != tt.want {
				t.Errorf("GiteaDB.GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiteaDB_GetAll(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbg := &GiteaDB{
				path: tt.fields.path,
			}
			if got := dbg.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GiteaDB.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
