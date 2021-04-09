package util

import (
	"testing"
)

func TestIsNumber(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "765は3桁以内の数値か",
			args: args{
				str: "765",
			},
			want: true,
		},
		{
			name: "23は3桁以内の数値か",
			args: args{
				str: "23",
			},
			want: true,
		},
		{
			name: "346283は3桁以内の数値か",
			args: args{
				str: "346283",
			},
			want: false,
		},
		{
			name: "hotaruは3桁以内の数値か",
			args: args{
				str: "hotaru",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.str); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	type args struct {
		array []string
		str   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[hotaru,nono,hiromi]にnonoはあるか",
			args: args{
				array: []string{
					"hotaru",
					"nono",
					"hiromi",
				},
				str: "nono",
			},
			want: 1,
		},
		{
			name: "[hotaru,hiromi,chiduru,yasuha]にtomoはあるか",
			args: args{
				array: []string{
					"hotaru",
					"hiromi",
					"chiduru",
					"yasuha",
				},
				str: "tomo",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.array, tt.args.str); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsStr(t *testing.T) {
	type args struct {
		reg string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "テスト: [0-9]+Hotaru",
			args: args{
				reg: "[0-9]+Hotaru",
				str: "ShiragikuHotaru",
			},
			want: false,
		},
		{
			name: "テスト: [A-Za-z]+Asahi",
			args: args{
				reg: "[A-Za-z]+Asahi",
				str: "SerizawaAsahi",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.reg, tt.args.str); got != tt.want {
				t.Errorf("ContainsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
