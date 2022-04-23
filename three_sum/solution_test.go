package three_sum

import (
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[-1,0,1,2,-1,-4]",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "[]",
			args: args{nums: []int{}},
			want: [][]int{},
		},
		{
			name: "[0]",
			args: args{nums: []int{0}},
			want: [][]int{},
		},
		{
			name: "[-1,0,1,2,-1,-4,-2,-3,3,0,4]",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}},
			want: [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !checkEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func checkEqual(target, from [][]int) bool {

	if len(target) != len(from) {
		return false
	}
	if len(target) == 0 && len(from) == 0 {
		return true
	}
	hashMap := make(map[[3]int]int)
	for _, value := range target {
		arr := [3]int{}
		for idx := range arr {
			arr[idx] = value[idx]
		}
		hashMap[arr] += 1
	}
	for _, value := range from {
		arr := [3]int{}
		for idx := range arr {
			arr[idx] = value[idx]
		}
		hashMap[arr] -= 1
	}
	for _, value := range hashMap {
		if value != 0 {
			return false
		}
	}
	return true
}
