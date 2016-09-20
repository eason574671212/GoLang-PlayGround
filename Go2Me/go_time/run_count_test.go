package go_time

import (
	"reflect"
	"testing"
)

func TestFastTimeTest(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name:"base",
			args:args{
				[]string{"1", "2", "3", "4"},
			},
			want: string("size 4 string"),
		},
		{
			name:"more",
			args:args{
				[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
			},
			want: string("size 10 string"),
		},
		{
			name:"more",
			args:args{
				[]string{
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
				},
			},
			want: string("size 50 string"),

		},
		{
			name:"more",
			args:args{
				[]string{
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
				},
			},
			want: string("size 200 string"),
		},
	}
	for _, tt := range tests {
		if got := FastTimeTest(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
			t.Logf("%q. FastTimeTest() = %v ns, test %v", tt.name, got, tt.want)
		}
	}
}
