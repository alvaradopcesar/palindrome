package duplicate

import (
	"reflect"
	"testing"
)

func Test_repetition(t *testing.T) {
	type args struct {
		st string
	}
	tests := []struct {
		name           string
		args           args
		wantStringList []string
	}{
		{
			name : "Test_repetition",
			args: args{ st : "abcabcdbb" },
			wantStringList: []string{"abc","abcd","b","b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStringList := repetition(tt.args.st); !reflect.DeepEqual(gotStringList, tt.wantStringList) {
				t.Errorf("repetition() = %v, want %v", gotStringList, tt.wantStringList)
			}
		})
	}
}

func Test_longestStringNoDuplicate(t *testing.T) {
	type args struct {
		stringList []string
	}
	tests := []struct {
		name        string
		args        args
		wantLongest string
	}{
		{
			name : "Test_longestStringNoDuplicate",
			args : args { stringList: []string{"abc","abcd","b","b"} },
			wantLongest: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLongest := longestStringNoDuplicate(tt.args.stringList); gotLongest != tt.wantLongest {
				t.Errorf("longestStringNoDuplicate() = %v, want %v", gotLongest, tt.wantLongest)
			}
		})
	}
}