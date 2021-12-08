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
					"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
					"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |  fcgedb cgb dgebacf gc",
					"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
					"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
					"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
					"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
					"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
					"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
					"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
					"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
				},
			},
			wantResult: 26,
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

func Test_parseSegments(t *testing.T) {
	type args struct {
		segments map[int]string
	}
	tests := []struct {
		name string
		args args
		want map[int][]string
	}{
		{
			name: "test1",
			args: args{
				segments: map[int]string{
					0: `
						 aaaa 
						b    c
						b    c
						 .... 
						e    f
						e    f
						 gggg `,
				},
			},
			want: map[int][]string{
				0: {
					"d", "e", "a", "g", "b", "c",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSegments(tt.args.segments); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []entry
	}{
		{
			name: "test1",
			args: args{
				input: []string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"},
			},
			want: []entry{
				{
					inputs:  []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
					outputs: []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
