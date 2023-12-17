package utils

type IntPriorityQueueItem struct {
	Item     interface{}
	Priority int
}

type IntPriorityQueue []IntPriorityQueueItem

func (q IntPriorityQueue) Len() int {
	return len(q)
}

func (q IntPriorityQueue) Less(i, j int) bool {
	return q[i].Priority < q[j].Priority
}

func (q IntPriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *IntPriorityQueue) Push(x interface{}) {
	item := x.(IntPriorityQueueItem)
	*q = append(*q, item)
}

func (q *IntPriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}
