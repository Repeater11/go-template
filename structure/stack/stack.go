// Package stack 提供了基于 Deque 的泛型栈实现，接口风格贴近 C++ std::stack。
package stack

import "github.com/Repeater11/go-template/structure/deque"

// Stack 对外只暴露 LIFO 语义。
type Stack[T any] struct {
	deque *deque.Deque[T]
}

// NewStack 创建并返回一个新的空栈。
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		deque: deque.NewDeque[T](),
	}
}

// Len 返回栈中元素数量。
func (s *Stack[T]) Len() int {
	if s == nil || s.deque == nil {
		return 0
	}
	return s.deque.Len()
}

// IsEmpty 判断栈是否为空。
func (s *Stack[T]) IsEmpty() bool {
	return s == nil || s.deque == nil || s.deque.IsEmpty()
}

// Top 返回栈顶元素但不移除。
func (s *Stack[T]) Top() (T, bool) {
	var zero T
	if s == nil || s.deque == nil {
		return zero, false
	}
	return s.deque.Back()
}

// Push 压入一个元素到栈顶。
func (s *Stack[T]) Push(elem T) {
	if s == nil {
		return
	}
	s.ensureDeque()
	s.deque.PushBack(elem)
}

// Pop 弹出并返回栈顶元素，若栈为空返回零值和 false。
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s == nil || s.deque == nil {
		return zero, false
	}
	return s.deque.PopBack()
}

// Clear 清空栈中的所有元素。
func (s *Stack[T]) Clear() {
	if s == nil || s.deque == nil {
		return
	}
	s.deque.Clear()
}

// ToSlice 以自底向顶顺序返回所有元素。
func (s *Stack[T]) ToSlice() []T {
	if s == nil || s.deque == nil {
		return []T{}
	}
	return s.deque.ToSlice()
}

// Swap 交换两个栈的内容。
func (s *Stack[T]) Swap(other *Stack[T]) {
	if s == nil || other == nil || s == other {
		return
	}
	s.ensureDeque()
	other.ensureDeque()
	s.deque, other.deque = other.deque, s.deque
}

// Clone 创建并返回栈的深拷贝。
func (s *Stack[T]) Clone() *Stack[T] {
	clone := NewStack[T]()
	if s == nil || s.deque == nil {
		return clone
	}
	clone.deque = s.deque.Clone()
	return clone
}

// Equal 判断两个栈是否拥有相同内容（自底向顶顺序）。
func Equal[T comparable](a, b *Stack[T]) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Len() != b.Len() {
		return false
	}
	as := a.ToSlice()
	bs := b.ToSlice()
	for i := range as {
		if as[i] != bs[i] {
			return false
		}
	}
	return true
}

// NotEqual 是 Equal 的反义函数。
func NotEqual[T comparable](a, b *Stack[T]) bool {
	return !Equal(a, b)
}

// ensureDeque 确保底层 deque 已初始化。
func (s *Stack[T]) ensureDeque() {
	if s.deque == nil {
		s.deque = deque.NewDeque[T]()
	}
}
