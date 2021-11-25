package util

import (
	"testing"
	"time"
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
			name: "半角文字",
			args: args{
				str:   "fuyuko",
				width: 5,
			},
			want: "fuyu…",
		},
		{
			name: "全角文字",
			args: args{
				str:   "芹沢あさひ",
				width: 5,
			},
			want: "芹沢…",
		},
		{
			name: "絵文字",
			args: args{
				str:   "🐶🐈🍺",
				width: 5,
			},
			want: "🐶🐈…",
		},
		{
			name: "半角文字 + 絵文字",
			args: args{
				str:   "suki💓💓",
				width: 7,
			},
			want: "suki💓…",
		},
		{
			name: "全角文字 + 絵文字",
			args: args{
				str:   "ビール🍺🍺🍺",
				width: 9,
			},
			want: "ビール🍺…",
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
			name: "半角文字（マッチする）",
			args: args{
				reg: "[1-9]{3,4}Pro",
				str: "765Pro",
			},
			want: true,
		},
		{
			name: "半角文字（マッチしない）",
			args: args{
				reg: "[0-9]+Hotaru",
				str: "ShiragikuHotaru",
			},
			want: false,
		},
		{
			name: "全角文字（マッチする）",
			args: args{
				reg: "七草(にちか|はづき)",
				str: "七草はづき",
			},
			want: true,
		},
		{
			name: "全角文字（マッチしない）",
			args: args{
				reg: "たなかまみ{3}",
				str: "たなかまみみ",
			},
			want: false,
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
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "3桁の数値",
			arg:  "765",
			want: true,
		},
		{
			name: "2桁の数値",
			arg:  "77",
			want: true,
		},
		{
			name: "6桁の数値",
			arg:  "346283",
			want: false,
		},
		{
			name: "数値以外（半角）",
			arg:  "hotaru",
			want: false,
		},
		{
			name: "数値以外（全角）",
			arg:  "凛世",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsThreeDigitsNumber(tt.arg); got != tt.want {
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
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "半角文字列配列（存在する）",
			args: args{
				array: []string{
					"hotaru",
					"nono",
					"hiromi",
				},
				str: "nono",
			},
			want:  1,
			want1: true,
		},
		{
			name: "半角文字列配列（存在しない）",
			args: args{
				array: []string{
					"hotaru",
					"hiromi",
					"chiduru",
					"yasuha",
				},
				str: "tomo",
			},
			want:  0,
			want1: false,
		},
		{
			name: "全角文字列配列（存在する）",
			args: args{
				array: []string{
					"白菊ほたる",
					"関裕美",
					"森久保乃々",
				},
				str: "森久保乃々",
			},
			want:  2,
			want1: true,
		},
		{
			name: "全角文字列配列（存在しない）",
			args: args{
				array: []string{
					"白菊ほたる",
					"関裕美",
					"松尾千鶴",
					"岡崎泰葉",
				},
				str: "藤居朋",
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IndexOf(tt.args.array, tt.args.str)
			if got != tt.want {
				t.Errorf("IndexOf() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IndexOf() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsSameDate(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want bool
	}{
		{
			name: "現在の日時",
			arg:  time.Now(),
			want: true,
		},
		{
			name: "今日の日付",
			arg:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
			want: true,
		},
		{
			name: "過去の日付",
			arg:  time.Date(2018, 4, 24, 0, 0, 0, 0, time.Local),
			want: false,
		},
		{
			name: "未来の日付",
			arg:  time.Now().Add(time.Hour * 24),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameDate(tt.arg); got != tt.want {
				t.Errorf("IsSameDate() = %v, want %v", got, tt.want)
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
			name: "LF",
			args: args{
				text: "rinze\n",
			},
			want: true,
		},
		{
			name: "CRLF",
			args: args{
				text: "morino\r\n",
			},
			want: true,
		},
		{
			name: "改行なし",
			args: args{
				text: "rinze",
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
