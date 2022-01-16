package main

import "testing"

func Test_findFirstStringInBracket(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "(bola basket)", args: args{str: "(bola basket)"}, want: "bola basket"},
		{name: "(Saya bermain)", args: args{str: "(Saya bermain)"}, want: "Saya bermain"},
		{name: "( Saya bermain )", args: args{str: "(Saya bermain)"}, want: "Saya bermain"},
		{name: "( Saya bermain)", args: args{str: "(Saya bermain)"}, want: "Saya bermain"},
		{name: "(Saya bermain )", args: args{str: "(Saya bermain)"}, want: "Saya bermain"},
		{name: "(Saya (Kamu) bermain)", args: args{str: "(Saya bermain)"}, want: "Saya bermain"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstStringInBracket(tt.args.str); got != tt.want {
				t.Errorf("findFirstStringInBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}
