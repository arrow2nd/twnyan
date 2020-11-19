package util

import (
	"reflect"
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
			name: "123は3桁以内の数値か",
			args: args{
				str: "123",
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
			name: "1111は3桁以内の数値か",
			args: args{
				str: "1111",
			},
			want: false,
		},
		{
			name: "abcは3桁以内の数値か",
			args: args{
				str: "abc",
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
			name: "[a,b,c]にbはあるか",
			args: args{
				array: []string{
					"a",
					"b",
					"c",
				},
				str: "b",
			},
			want: 1,
		},
		{
			name: "[bb,aa,abc,ef,a]にefgはあるか",
			args: args{
				array: []string{
					"bb",
					"aa",
					"abc",
					"ef",
					"a",
				},
				str: "efg",
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

func TestFetchStringSpecifiedType(t *testing.T) {
	type args struct {
		args  []string
		aType []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "[12]から[12,'']を取り出す",
			args: args{
				args:  []string{"12"},
				aType: []string{"num", "str"},
			},
			want: []string{"12", ""},
		},
		{
			name: "[123,abcd]から[123]を取り出す",
			args: args{
				args:  []string{"123", "abcd"},
				aType: []string{"num", "num"},
			},
			want: []string{"123", ""},
		},
		{
			name: "引数が多すぎる",
			args: args{
				args:  []string{"123", "abcd"},
				aType: []string{"num"},
			},
			wantErr: true,
		},
		{
			name: "引数が無い",
			args: args{
				args:  []string{},
				aType: []string{"num"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchStringSpecifiedType(tt.args.args, tt.args.aType...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchStringSpecifiedType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchStringSpecifiedType() = %v, want %v", got, tt.want)
			}
		})
	}
}