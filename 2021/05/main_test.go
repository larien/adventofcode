package main

import (
	"reflect"
	"testing"
)

func Test_solve1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "test1",
			args: args{
				input: []string{
					"0,9 -> 5,9",
					"8,0 -> 0,8", // diagonal
					"9,4 -> 3,4",
					"2,2 -> 2,1",
					"7,0 -> 7,4",
					"6,4 -> 2,0", // diagonal
					"0,9 -> 2,9",
					"3,4 -> 1,4",
					"0,0 -> 8,8", // diagonal
					"5,5 -> 8,2", // diagonal
				},
			},
			wantResult: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solve1(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("solve1() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "test1",
			args: args{
				input: []string{
					"0,9 -> 5,9",
					"8,0 -> 0,8", // diagonal
					"9,4 -> 3,4",
					"2,2 -> 2,1",
					"7,0 -> 7,4",
					"6,4 -> 2,0", // diagonal
					"0,9 -> 2,9",
					"3,4 -> 1,4",
					"0,0 -> 8,8", // diagonal
					"5,5 -> 8,2", // diagonal
				},
			},
			wantResult: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solve2(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("solve2() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name      string
		args      args
		wantVents []vent
	}{
		{
			name: "test1",
			args: args{
				input: []string{
					"0,9 -> 5,9",
					"8,0 -> 0,8",
				},
			},
			wantVents: []vent{
				{
					x1: 0,
					y1: 9,
					x2: 5,
					y2: 9,
				},
				{
					x1: 8,
					y1: 0,
					x2: 0,
					y2: 8,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVents := parseInput(tt.args.input)
			if !reflect.DeepEqual(gotVents, tt.wantVents) {
				t.Errorf("parseInput() = %v, want %v", gotVents, tt.wantVents)
			}
		})
	}
}
