package search

import "testing"

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。

// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

// 你可以假设数组中不存在重复的元素。

// 你的算法时间复杂度必须是 O(log n) 级别。

// 示例 1:

// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4
// 示例 2:

// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1

// 链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array

func TestSearchInRotatedSortedArray(t *testing.T) {
	t.Log(searchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

// TODO
func searchInRotatedSortedArray(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + (high - low)) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] { // 左边有序
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // 右边有序
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
