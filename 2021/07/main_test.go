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
				input: []string{"16,1,2,0,4,2,7,1,2,14"},
			},
			wantResult: 37,
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
				input: []string{"16,1,2,0,4,2,7,1,2,14"},
			},
			wantResult: 168,
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

func Test_calculateCosts(t *testing.T) {
	type args struct {
		positions []int
		position  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				positions: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
				position:  5,
			},
			want: 168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCost(tt.args.positions, tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateCosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
