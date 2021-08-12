package util

import (
	"testing"
)

func TestTruncateString(t *testing.T) {
	type args struct {
		str   string
		width int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "'fuyuko'を5文字に丸める",
			args: args{
				str:   "fuyuko",
				width: 5,
			},
			want: "fuyu…",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncateString(tt.args.str, tt.args.width); got != tt.want {
				t.Errorf("TruncateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchesRegexp(t *testing.T) {
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
			name: "'ShiragikuHotaru' に '[0-9]+Hotaru' がマッチするか",
			args: args{
				reg: "[0-9]+Hotaru",
				str: "ShiragikuHotaru",
			},
			want: false,
		},
		{
			name: "'SerizawaAsahi' に '[A-Za-z]+Asahi' がマッチするか",
			args: args{
				reg: "[A-Za-z]+Asahi",
				str: "SerizawaAsahi",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchesRegexp(tt.args.reg, tt.args.str); got != tt.want {
				t.Errorf("MatchesRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsThreeDigitsNumber(t *testing.T) {
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
			if got := IsThreeDigitsNumber(tt.args.str); got != tt.want {
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

func TestIsEndLFCode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "末尾が改行コードかどうか 1",
			args: args{
				text: "rinze\n",
			},
			want: true,
		},
		{
			name: "末尾が改行コードかどうか 2",
			args: args{
				text: "morino",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEndLFCode(tt.args.text); got != tt.want {
				t.Errorf("IsEndLFCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
