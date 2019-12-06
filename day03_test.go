package main

import (
	"reflect"
	"testing"
)

func Test_measure(t *testing.T) {
	type args struct {
		origin coords
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want []coords
	}{
		name: "simple",
		args: {origin: coords{x: 2, y: 2}, width: 4, height: 3},
		want: []coords{
			{x: 2, y: 2}, {x: 3, y: 2}, {x: 4, y: 2}, {x: 5, y: 2}, {x: 6, y: 2},
			{x: 2, y: 3}, {x: 3, y: 3}, {x: 4, y: 3}, {x: 5, y: 3}, {x: 6, y: 3},
			{x: 2, y: 4}, {x: 3, y: 4}, {x: 4, y: 4}, {x: 5, y: 4}, {x: 6, y: 4},
			{x: 2, y: 5}, {x: 3, y: 5}, {x: 4, y: 5}, {x: 5, y: 5}, {x: 6, y: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := measure(tt.args.origin, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("measure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
