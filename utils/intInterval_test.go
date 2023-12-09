package utils

import (
	"reflect"
	"testing"
)

func TestIntInterval_String(t *testing.T) {
	interval := IntInterval{Start: 2, End: 5}
	expected := "[2, 5["
	if result := interval.String(); result != expected {
		t.Errorf("String representation is incorrect, got: %s, expected: %s.", result, expected)
	}
}

func TestIntInterval_Equal(t *testing.T) {
	interval1 := IntInterval{Start: 2, End: 5}
	interval2 := IntInterval{Start: 2, End: 5}
	interval3 := IntInterval{Start: 3, End: 6}
	if !interval1.Equal(interval2) {
		t.Errorf("Expected intervals %s and %s to be equal, but they were not.", interval1, interval2)
	}
	if interval1.Equal(interval3) {
		t.Errorf("Expected intervals %s and %s to be unequal, but they were equal.", interval1, interval2)
	}
}

func TestIsIn(t *testing.T) {
	intervals := []IntInterval{
		{Start: 1, End: 5},
		{Start: 6, End: 10},
		{Start: 11, End: 15},
	}
	interval := IntInterval{Start: 6, End: 10}
	exists, index := interval.IsIn(intervals)
	if !exists || index != 1 {
		t.Errorf("Interval %v not found in the slice %v.", interval, intervals)
	}
	interval = IntInterval{Start: 16, End: 20}
	exists, index = interval.IsIn(intervals)
	if exists || index != -1 {
		t.Errorf("Interval %v found in the slice %v, but it shouldn't exist.", interval, intervals)
	}
}

func TestContain(t *testing.T) {
	interval1 := IntInterval{Start: 2, End: 6}
	interval2 := IntInterval{Start: 3, End: 5}
	interval3 := IntInterval{Start: 7, End: 9}
	if !interval1.Contain(interval2) {
		t.Errorf("Expected %v to contain %v, but it did not.", interval1, interval2)
	}
	if interval1.Contain(interval3) {
		t.Errorf("Expected %v to not contain %v, but it did.", interval1, interval3)
	}
}

func TestIntersection(t *testing.T) {
	interval1 := IntInterval{Start: 2, End: 5}
	interval2 := IntInterval{Start: 4, End: 7}
	expected := IntInterval{Start: 4, End: 5}
	result := interval1.Intersection(interval2)
	if !result.Equal(expected) {
		t.Errorf("Intersection was incorrect, got: %v, expected: %v.", result, expected)
	}
	interval3 := IntInterval{Start: 6, End: 8}
	result = interval1.Intersection(interval3)
	if !result.Equal(EmptyIntInterval) {
		t.Errorf("Expected empty interval for non-intersecting intervals, got: %v", result)
	}
}

func TestUnion(t *testing.T) {
	interval1 := IntInterval{Start: 2, End: 5}
	interval2 := IntInterval{Start: 4, End: 7}
	expected := []IntInterval{{Start: 2, End: 7}}
	result := interval1.Union(interval2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Union was incorrect, got: %v, expected: %v.", result, expected)
	}
	interval3 := IntInterval{Start: 8, End: 10}
	result = interval1.Union(interval3)
	if len(result) != 2 {
		t.Errorf("Expected disjoint intervals for non-overlapping intervals, got: %v.", result)
	}
}

func TestDifference(t *testing.T) {
	interval1 := IntInterval{Start: 2, End: 7}
	interval2 := IntInterval{Start: 4, End: 6}
	expected := []IntInterval{{Start: 2, End: 4}, {Start: 6, End: 7}}
	result := interval1.Difference(interval2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Difference was incorrect, got: %v, expected: %v.", result, expected)
	}
	interval3 := IntInterval{Start: 1, End: 9}
	result = interval1.Difference(interval3)
	if len(result) != 0 {
		t.Errorf("Expected no difference for fully contained interval, got: %v.", result)
	}
}

func TestDisjointUnion(t *testing.T) {
	interval1 := IntInterval{Start: 2, End: 5}
	interval2 := IntInterval{Start: 7, End: 9}
	expected := []IntInterval{{Start: 2, End: 5}, {Start: 7, End: 9}}
	result := interval1.DisjointUnion(interval2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DisjointUnion was incorrect, got: %v, expected: %v.", result, expected)
	}
	interval3 := IntInterval{Start: 4, End: 8}
	result = interval1.DisjointUnion(interval3)
	if len(result) != 2 {
		t.Errorf("Expected disjoint intervals for intersecting intervals, got: %v.", result)
	}
}

func TestSplitOn(t *testing.T) {
	ab := IntInterval{2, 5}
	cd := IntInterval{5, 10}
	inter, exter := ab.splitOn(cd)
	expectedInter := IntInterval{}
	expectedExter := []IntInterval{{2, 5}}

	if inter != expectedInter {
		t.Errorf("Intersection is incorrect, got: %v, want: %v.", inter, expectedInter)
	}

	if !reflect.DeepEqual(exter, expectedExter) {
		t.Errorf("External part is incorrect, got: %v, want: %v.", exter, expectedExter)
	}

	ab = IntInterval{2, 10}
	cd = IntInterval{5, 7}
	inter, exter = ab.splitOn(cd)
	expectedInter = IntInterval{5, 7}
	expectedExter = []IntInterval{{2, 5}, {7, 10}}

	if inter != expectedInter {
		t.Errorf("Intersection is incorrect, got: %v, want: %v.", inter, expectedInter)
	}

	if !reflect.DeepEqual(exter, expectedExter) {
		t.Errorf("External part is incorrect, got: %v, want: %v.", exter, expectedExter)
	}
}

func TestSortIntIntervals(t *testing.T) {
	intervals := []IntInterval{
		{Start: 5, End: 1},
		{Start: 1, End: 4},
		{Start: 2, End: 8},
		{Start: 2, End: 6},
	}
	expected := []IntInterval{
		{Start: 1, End: 4},
		{Start: 2, End: 6},
		{Start: 2, End: 8},
		{Start: 5, End: 1},
	}
	SortIntIntervals(intervals)
	if !reflect.DeepEqual(intervals, expected) {
		t.Errorf("Intervals were not sorted as expected, got: %v, want: %v.", intervals, expected)
	}
}

func TestMergeIntIntervals(t *testing.T) {
	intervals1 := []IntInterval{
		{Start: 1, End: 3},
		{Start: 2, End: 4},
		{Start: 5, End: 7},
	}
	expected1 := []IntInterval{
		{Start: 1, End: 4},
		{Start: 5, End: 7},
	}
	result1 := MergeIntIntervals(intervals1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("MergeIntIntervals failed, expected: %v, got: %v", expected1, result1)
	}
	intervals2 := []IntInterval{
		{Start: 2, End: 4},
		{Start: 6, End: 8},
		{Start: 10, End: 12},
	}
	expected2 := []IntInterval{
		{Start: 2, End: 4},
		{Start: 6, End: 8},
		{Start: 10, End: 12},
	}
	result2 := MergeIntIntervals(intervals2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("MergeIntIntervals failed, expected: %v, got: %v", expected2, result2)
	}
	intervals3 := []IntInterval{
		{Start: 1, End: 3},
		{Start: 2, End: 4},
		{Start: 4, End: 7},
	}
	expected3 := IntInterval{Start: 1, End: 7}
	result3 := MergeIntIntervals(intervals3)[0]
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("MergeIntIntervals failed, expected: %v, got: %v", expected3, result3)
	}
}
