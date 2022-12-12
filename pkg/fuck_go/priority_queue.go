package fuck_go

import "container/heap"

// An PriorityQueueEntry is something we manage in a priority queue.
type PriorityQueueEntry[T interface{}] struct {
	value    T   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T interface{}] []*PriorityQueueEntry[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*PriorityQueueEntry[T])
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) PushT(priority int, x T) {
	item := &PriorityQueueEntry[T]{
		value:    x,
		priority: priority,
	}
	heap.Push(pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) PopT() T {
	x, _ := heap.Pop(pq).(*PriorityQueueEntry[T])
	return x.value
}
