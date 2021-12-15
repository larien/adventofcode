package main

import (
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
					"2199943210",
					"3987894921",
					"9856789892",
					"8767896789",
					"9899965678",
				},
			},
			wantResult: 15,
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
					"2199943210",
					"3987894921",
					"9856789892",
					"8767896789",
					"9899965678",
				},
			},
			wantResult: 1134,
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

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_addLowests(t *testing.T) {
	type args struct {
		inputs []int
		input  int
	}
	tests := []struct {
		name       string
		args       args
		wantLowest int
	}{
		{
			name: "test1",
			args: args{
				inputs: []int{4, 2, 6},
				input:  3,
			},
			wantLowest: 2,
		},
		{
			name: "test2",
			args: args{
				inputs: []int{2, 4, 6},
				input:  1,
			},
			wantLowest: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLowest := addLowests(tt.args.inputs, tt.args.input); gotLowest != tt.wantLowest {
				t.Errorf("addLowests() = %v, want %v", gotLowest, tt.wantLowest)
			}
		})
	}
}

func Test_findBasin(t *testing.T) {
	type args struct {
		input   [][]int
		lowests [3]int
	}
	tests := []struct {
		name      string
		args      args
		wantBasin int
	}{
		{
			name: "test1",
			args: args{
				input: [][]int{
					{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
				lowests: [3]int{2, 0, 0},
			},
			wantBasin: 3,
		},
		{
			name: "test2",
			args: args{
				input: [][]int{
					{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
				lowests: [3]int{0, 9, 0},
			},
			wantBasin: 9,
		},
		{
			name: "test3",
			args: args{
				input: [][]int{
					{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
				lowests: [3]int{5, 2, 2},
			},
			wantBasin: 14,
		},
		{
			name: "test4",
			args: args{
				input: [][]int{
					{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
					{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
					{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
					{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
					{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
				},
				lowests: [3]int{5, 6, 4},
			},
			wantBasin: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBasin := findBasin(tt.args.input, tt.args.lowests); gotBasin != tt.wantBasin {
				t.Errorf("findBasin() = %v, want %v", gotBasin, tt.wantBasin)
			}
		})
	}
}
