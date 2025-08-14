package task1

import "sort"

/*
*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，
如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func MergeArray(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 1. 按区间左端点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 扫描合并
	var merged [][]int
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		// 当前区间与下一个区间有重叠：更新右端点
		if intervals[i][0] <= cur[1] {
			if intervals[i][1] > cur[1] {
				cur[1] = intervals[i][1]
			}
		} else {
			// 无重叠，把 cur 加入结果
			merged = append(merged, cur)
			cur = intervals[i]
		}
	}
	// 把最后一个区间加入结果
	merged = append(merged, cur)
	return merged
}
