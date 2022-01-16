package main

import "testing"

func Test_filterToAnagram(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "multiple anagram",
			args: args{
				data: []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterToAnagram(tt.args.data)
		})
	}
}
