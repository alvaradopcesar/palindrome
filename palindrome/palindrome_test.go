package palindrome

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Pal
	}{
		{
			name : "test 01",
			want : &Pal{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPal_IsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name : "cesar",
			args: args{ s : "cesar" },
			want : false,
		},
		{
			name: "civic",
			args: args{s: "civic"},
			want: true,
		},
		{
			name : "madam",
			args: args{ s : "madam" },
			want : true,
		},
		{
			name : "radar",
			args: args{ s : "radar" },
			want : true,
		},
		{
			name : "123456789",
			args: args{ s : "123456789" },
			want : false,
		},
		{
			name : "1234321",
			args: args{ s : "1234321" },
			want : true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pal := &Pal{}
			if got := pal.IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
