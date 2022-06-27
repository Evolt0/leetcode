package select_sort

import (
	"reflect"
	"testing"
)

func Test_selectSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: t.Name(),
			args: args{
				arr: []int{4, 5, 1, 2, 3, 6, 1},
			},
			want: []int{1, 1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
