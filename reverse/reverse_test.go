package reverse

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "reverse()",
			args: args{
				x: 123,
			},
			want: 321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRev := reverse(tt.args.x); gotRev != tt.want {
				t.Errorf("reverse() = %v, want %v", gotRev, tt.want)
			}
		})
	}
}
