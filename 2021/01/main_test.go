package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "test",
			args: args{
				input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			wantResult: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solve(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("Solve() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
