package utils

import (
	"fmt"
	"sort"
)

// [Start, End[
type IntInterval struct{ Start, End int }

type IntIntervalOrIntIntervals interface{}

var EmptyIntInterval = IntInterval{}

func (ab IntInterval) String() string {
	return fmt.Sprintf("[%d, %d[", ab.Start, ab.End)
}

func (ab IntInterval) Equal(cd IntInterval) bool {
	a, b := ab.Start, ab.End
	c, d := cd.Start, cd.End
	return c == a && b == d
}

func (ab IntInterval) IsIn(intervals []IntInterval) (ok bool, index int) {
	for index, interval := range intervals {
		if ab.Equal(interval) {
			return true, index
		}
	}
	return false, -1
}

func (ab IntInterval) Contain(cd IntInterval) bool {
	a, b := ab.Start, ab.End
	c, d := cd.Start, cd.End
	return a <= c && d <= b
}

func (ab IntInterval) Intersection(cd IntInterval) IntInterval {
	a, b := ab.Start, ab.End
	c, d := cd.Start, cd.End
	intersectionStart := max(a, c)
	intersectionEnd := min(b, d)
	isOverlapping := intersectionStart < intersectionEnd
	if isOverlapping {
		return IntInterval{intersectionStart, intersectionEnd}
	}
	return EmptyIntInterval
}

func (ab IntInterval) Union(cd IntInterval) []IntInterval {
	a, b := ab.Start, ab.End
	c, d := cd.Start, cd.End
	isDisjoint := b < c || d < a
	if isDisjoint {
		if a < c {
			return []IntInterval{ab, cd}
		}
		return []IntInterval{cd, ab}
	}
	unionStart := min(a, c)
	unionEnd := max(b, d)
	return []IntInterval{{unionStart, unionEnd}}
}

func (ab IntInterval) Difference(cd IntInterval) []IntInterval {
	a, b := ab.Start, ab.End
	c, d := cd.Start, cd.End
	if cd.Contain(ab) {
		return []IntInterval{}
	}
	var difference []IntInterval
	if a < c {
		End := min(b, c)
		difference = append(difference, IntInterval{a, End})
	}
	if b > d {
		Start := max(a, d)
		difference = append(difference, IntInterval{Start, b})
	}
	return difference
}

func (ab IntInterval) DisjointUnion(cd IntInterval) []IntInterval {
	a, b := ab.Start, ab.End
	c, d := cd.Start, cd.End
	isDisjoint := b < c || d < a
	if isDisjoint {
		if a < c {
			return []IntInterval{ab, cd}
		}
		return []IntInterval{cd, ab}
	}
	var disjointUnion []IntInterval
	if a < c {
		disjointUnion = append(disjointUnion, IntInterval{a, c})
	} else {
		disjointUnion = append(disjointUnion, IntInterval{c, a})
	}
	if b < d {
		disjointUnion = append(disjointUnion, IntInterval{b, d})
	} else {
		disjointUnion = append(disjointUnion, IntInterval{d, b})
	}
	return disjointUnion
}

func (ab IntInterval) splitOn(cd IntInterval) (inter IntInterval, exter []IntInterval) {
	a, b := ab.Start, ab.End
	c, d := cd.Start, cd.End

	isDisjoint := b <= c || d <= a
	if isDisjoint {
		exter = append(exter, ab)
		return
	}

	inter = IntInterval{max(a, c), min(b, d)}

	if a < c {
		before := IntInterval{a, c}
		exter = append(exter, before)
	}
	if b > d {
		after := IntInterval{d, b}
		exter = append(exter, after)
	}

	return
}

func SortIntIntervals(intervals []IntInterval) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start != intervals[j].Start {
			return intervals[i].Start < intervals[j].Start
		}
		return intervals[i].End < intervals[j].End
	})
}

func MergeIntIntervals(intervals []IntInterval) []IntInterval {
	if len(intervals) <= 0 {
		return nil
	}
	SortIntIntervals(intervals)
	merged := []IntInterval{}
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := merged[len(merged)-1]
		if current.Start <= lastMerged.End {
			if current.End > lastMerged.End {
				lastMerged.End = current.End
				merged[len(merged)-1] = lastMerged
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}
