// Package queue 提供了基于 Deque 泛型队列的实现。
package queue

import "github.com/Repeater11/go-template/structure/deque"

// Queue 是一个泛型队列，基于 Deque 实现。
type Queue[T any] struct {
	deque *deque.Deque[T]
}

// NewQueue 创建并返回一个新的空队列。
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		deque: deque.NewDeque[T](),
	}
}

// Len 返回队列中元素的数量。
func (q *Queue[T]) Len() int {
	if q == nil || q.deque == nil {
		return 0
	}
	return q.deque.Len()
}

// IsEmpty 检查队列是否为空。
func (q *Queue[T]) IsEmpty() bool {
	if q == nil || q.deque == nil {
		return true
	}
	return q.deque.IsEmpty()
}

// Front 返回队列前端的元素但不移除它。
func (q *Queue[T]) Front() (T, bool) {
	if q.deque.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.deque.Front()
}

// Back 返回队列后端的元素但不移除它。
func (q *Queue[T]) Back() (T, bool) {
	if q.deque.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.deque.Back()
}

// Push 在队列后端添加一个元素。
func (q *Queue[T]) Push(elem T) {
	q.ensureDeque()
	q.deque.PushBack(elem)
}

// Pop 移除并返回队列前端的元素。
// 如果队列为空，返回零值和 false。
func (q *Queue[T]) Pop() (T, bool) {
	if q.deque.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.deque.PopFront()
}

// Clear 清空队列中的所有元素。
func (q *Queue[T]) Clear() {
	if q.IsEmpty() {
		return
	}
	q.deque.Clear()
}

// ensureDeque 确保内部的 deque 已初始化。
func (q *Queue[T]) ensureDeque() {
	if q.deque == nil {
		q.deque = deque.NewDeque[T]()
	}
}

// To Slice 返回队列中所有元素的切片表示。
func (q *Queue[T]) ToSlice() []T {
	if q.IsEmpty() {
		return []T{}
	}
	return q.deque.ToSlice()
}

// Swap 交换两个队列的内容。
func (q *Queue[T]) Swap(other *Queue[T]) {
	q.ensureDeque()
	other.ensureDeque()
	q.deque.Swap(other.deque)
}

// Equal 检查两个队列是否相等。
func Equal[T comparable](q1, q2 *Queue[T]) bool {
	if q1.Len() != q2.Len() {
		return false
	}
	slice1 := q1.ToSlice()
	slice2 := q2.ToSlice()
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

// Clone 创建并返回队列的一个深拷贝。
func (q *Queue[T]) Clone() *Queue[T] {
	if q.IsEmpty() {
		return NewQueue[T]()
	}
	cloneDeque := q.deque.Clone()
	return &Queue[T]{deque: cloneDeque}
}
