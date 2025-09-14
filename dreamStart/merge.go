package dreamstart

import "sort"

func merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	resIntervals := [][]int{}

	//区间按左边排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	resIntervals = append(resIntervals, intervals[0])

	for i := 1; i < len(intervals); i++ {
		resIntervalsLast := resIntervals[len(resIntervals)-1]

		if resIntervalsLast[1] >= intervals[i][0] {

			if resIntervalsLast[1] < intervals[i][1] {
				resIntervalsLast[1] = intervals[i][1]
			}
		} else {
			resIntervals = append(resIntervals, intervals[i])
		}

	}

	return resIntervals
}
