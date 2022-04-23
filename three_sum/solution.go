package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// sort the nums
	sort.Ints(nums) // 避免重複的值
	result := make([][]int, 0)
	length := len(nums)
	start, end, index := 0, length-1, 1
	for ; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] { // 當於到左邊的值重複的值,因為已經找過,直接往右移
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] { // 當於到右邊的值重複的值,因為已經找過,直接往左移
				end--
				continue
			}
			addNum := nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
