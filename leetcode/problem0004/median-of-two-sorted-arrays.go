package problem0004

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := MergeSlice(nums1, nums2)
	sort.Sort(sort.IntSlice(nums))
	return medianOf(nums)
}
func MergeSlice(s1 []int, s2 []int) []int {
	slice := make([]int, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

func medianOf(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}

	return float64(nums[l/2])
}
