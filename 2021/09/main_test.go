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
				input: []string{},
			},
			wantResult: 0,
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
