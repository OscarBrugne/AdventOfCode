package utils

type CostQueueItem struct {
	Item interface{}
	Cost int
}

type CostQueue []CostQueueItem

func (q CostQueue) Len() int {
	return len(q)
}

func (q CostQueue) Less(i, j int) bool {
	return q[i].Cost < q[j].Cost
}

func (q CostQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *CostQueue) Push(x interface{}) {
	item := x.(CostQueueItem)
	*q = append(*q, item)
}

func (q *CostQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}
