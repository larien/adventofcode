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
				input: []string{"3,4,3,1,2"},
			},
			wantResult: 5934,
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
				input: []string{"3,4,3,1,2"},
			},
			wantResult: 26984457539,
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

func Test_simulateLanternFish(t *testing.T) {
	type args struct {
		input []string
		days  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				input: []string{"3,4,3,1,2"},
				days:  18,
			},
			want: 26,
		},
		{
			name: "test2",
			args: args{
				input: []string{"3,4,3,1,2"},
				days:  80,
			},
			want: 5934,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateLanternFish(tt.args.input, tt.args.days); got != tt.want {
				t.Errorf("simulateLanternFish() = %v, want %v", got, tt.want)
			}
		})
	}
}
