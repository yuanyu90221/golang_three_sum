# three_sum

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

## Examples

```
Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:

Input: nums = []
Output: []
Example 3:

Input: nums = [0]
Output: []
 

Constraints:

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

```

## 觀察

首先是 要找出 3個不同 element 在 array 中, 並且 3個 element 相加為 0

且找出的值不能重複

因此 可以透過 排序來讓避免讓重複的值被搜尋

以 array index 當作 其中一個值 剩下兩個直就可以透過 two sum的作法來做

然後每次檢查 當下值與前一個 index 是否相同 如果是代表已經查過了 

## 實作

```golang
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

```