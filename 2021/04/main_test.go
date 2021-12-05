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
					"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
					"",
					"22 13 17 11  0",
					" 8  2 23  4 24",
					"21  9 14 16  7",
					" 6 10  3 18  5",
					" 1 12 20 15 19",
					"",
					" 3 15  0  2 22",
					" 9 18 13 17  5",
					"19  8  7 25 23",
					"20 11 10 24  4",
					"14 21 16 12  6",
					"",
					"14 21 17 24  4",
					"10 16 15  9 19",
					"18  8 23 26 20",
					"22 11 13  6  5",
					" 2  0 12  3  7",
				},
			},
			wantResult: 4512,
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
					"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
					"",
					"22 13 17 11  0",
					" 8  2 23  4 24",
					"21  9 14 16  7",
					" 6 10  3 18  5",
					" 1 12 20 15 19",
					"",
					" 3 15  0  2 22",
					" 9 18 13 17  5",
					"19  8  7 25 23",
					"20 11 10 24  4",
					"14 21 16 12  6",
					"",
					"14 21 17 24  4",
					"10 16 15  9 19",
					"18  8 23 26 20",
					"22 11 13  6  5",
					" 2  0 12  3  7",
				},
			},
			wantResult: 1924,
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

func Test_boards_checkWinners(t *testing.T) {
	tests := []struct {
		name    string
		b       boards
		winners map[int]int
	}{
		{
			name: "row winner",
			b: boards{
				{
					{1, 2, 3, 4, 5},
					{1, 2, 3, 4, 5},
					{1, 2, 3, 4, 5},
					{-1, -1, 3, -1, -1},
					{1, 2, 3, 4, 5},
				},
				{
					{1, 2, 3, 4, 5},
					{1, 2, 3, 4, 5},
					{1, 2, 3, 4, 5},
					{-1, -1, -1, -1, -1},
					{1, 2, 3, 4, 5},
				},
			},
			winners: map[int]int{
				1: 3,
			},
		},
		{
			name: "column winner",
			b: boards{
				{
					{1, 2, 3, 4, 5},
					{1, 2, 3, 4, 5},
					{1, 2, 3, 4, 5},
					{-1, -1, 3, -1, -1},
					{1, 2, 3, 4, 5},
				},
				{
					{1, 2, 3, 4, -1},
					{1, 2, 3, 4, -1},
					{1, 2, 3, 4, -1},
					{-1, -1, 3, -1, -1},
					{1, 2, 3, 4, -1},
				},
			},
			winners: map[int]int{
				1: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			winners := tt.b.checkWinners()
			if reflect.DeepEqual(winners, tt.winners) {
				t.Errorf("boards.checkWinners() got = %v, want %v", winners, tt.winners)
			}
		})
	}
}
