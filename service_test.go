package main

import (
	"reflect"
	"testing"

	"github.com/savageking-io/noerrorcode/database"
	"github.com/urfave/cli"
)

func TestServe(t *testing.T) {
	ConfigFilepath = "testdata/config/unit-test.yaml"
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//{"Initial test", args{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Serve(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Serve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetLogLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
	}{
		{"No log level", args{level: ""}},
		{"Trace level", args{level: "trace"}},
		{"Debug level", args{level: "debug"}},
		{"Info level", args{level: "info"}},
		{"Warn level", args{level: "warn"}},
		{"Error level", args{level: "error"}},
		{"Fatal level", args{level: "fatal"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLogLevel(tt.args.level)
		})
	}
}

func TestInitMySQL(t *testing.T) {
	type args struct {
		config *database.MySQLConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *database.MySQL
		wantErr bool
	}{
		{"Nil config", args{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitMySQL(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitMySQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitMySQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
