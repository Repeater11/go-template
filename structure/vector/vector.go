// Package vector 提供了泛型动态数组的实现。
package vector

import (
	"cmp"
	"slices"
)

// Vector 是一个通用的动态数组实现。
// T 可以是任何类型。
// Vector 提供了多种方法来操作和管理动态数组。
type Vector[T any] struct {
	data []T
}

// NewVector 创建一个空的 Vector。
// 可选地，可以传入初始元素来填充 Vector。
func NewVector[T any](elements ...T) *Vector[T] {
	return &Vector[T]{
		data: append([]T{}, elements...),
	}
}

// NewVectorFill 创建一个指定大小的 Vector，并用给定的值填充。
// 如果未提供值，则使用类型的零值进行填充。
func NewVectorFill[T any](size int, value ...T) *Vector[T] {
	data := make([]T, size)
	if len(value) > 0 {
		for i := range data {
			data[i] = value[0]
		}
	}
	return &Vector[T]{
		data: data,
	}
}

// NewVectorFromSlice 从给定的切片创建一个 Vector。
func NewVectorFromSlice[T any](slice []T) *Vector[T] {
	return &Vector[T]{
		data: append([]T{}, slice...),
	}
}

// Len 返回 Vector 中元素的数量。
func (v *Vector[T]) Len() int {
	return len(v.data)
}

// IsEmpty 检查 Vector 是否为空。
func (v *Vector[T]) IsEmpty() bool {
	return len(v.data) == 0
}

// PushBack 在 Vector 的末尾添加一个或多个元素。
func (v *Vector[T]) PushBack(elements ...T) {
	v.data = append(v.data, elements...)
}

// PopBack 从 Vector 的末尾移除并返回最后一个元素。
// 如果 Vector 为空，返回零值和 false。
func (v *Vector[T]) PopBack() (T, bool) {
	var zero T
	if v.IsEmpty() {
		return zero, false
	}
	element := v.data[len(v.data)-1]
	v.data = v.data[:len(v.data)-1]
	return element, true
}

// At 返回指定索引处的元素，不检查边界。
func (v *Vector[T]) At(index int) T {
	return v.data[index]
}

// Get 安全地返回指定索引处的元素。
// 如果索引无效，返回零值和 false。
func (v *Vector[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(v.data) {
		var zero T
		return zero, false
	}
	return v.data[index], true
}

// Set 设置指定索引处的元素的值。
// 如果索引无效返回 false。
func (v *Vector[T]) Set(index int, value T) bool {
	if index < 0 || index >= len(v.data) {
		return false
	}
	v.data[index] = value
	return true
}

// Insert 在指定索引处插入一个或多个元素。
// 如果索引无效返回 false。
func (v *Vector[T]) Insert(index int, elements ...T) bool {
	if index < 0 || index > len(v.data) {
		return false
	}
	v.data = slices.Insert(v.data, index, elements...)
	return true
}

// Erase 移除指定范围内的元素 [begin, end)。
// 如果范围无效返回 false。
func (v *Vector[T]) Erase(begin, end int) bool {
	if begin < 0 || end > len(v.data) || begin >= end {
		return false
	}
	v.data = slices.Delete(v.data, begin, end)
	return true
}

// Clear 移除 Vector 中的所有元素。
func (v *Vector[T]) Clear() {
	v.data = v.data[:0]
}

// Front 返回 Vector 的第一个元素。
// 如果 Vector 为空，返回零值和 false。
func (v *Vector[T]) Front() (T, bool) {
	if v.IsEmpty() {
		var zero T
		return zero, false
	}
	return v.data[0], true
}

// Back 返回 Vector 的最后一个元素。
// 如果 Vector 为空，返回零值和 false。
func (v *Vector[T]) Back() (T, bool) {
	if v.IsEmpty() {
		var zero T
		return zero, false
	}
	return v.data[len(v.data)-1], true
}

// Capacity 返回 Vector 的当前容量。
func (v *Vector[T]) Capacity() int {
	return cap(v.data)
}

// Reserve 调整 Vector 的容量以至少容纳指定数量的元素。
func (v *Vector[T]) Reserve(newCap int) {
	if newCap > cap(v.data) {
		newData := make([]T, len(v.data), newCap)
		copy(newData, v.data)
		v.data = newData
	}
}

// Resize 调整 Vector 的大小。
// 如果新大小大于当前大小，使用提供的值或零值填充新元素。
func (v *Vector[T]) Resize(newSize int, value ...T) {
	if newSize < 0 {
		return
	}

	currSize := len(v.data)

	if newSize < currSize {
		v.data = v.data[:newSize]
	} else if newSize > currSize {
		v.data = slices.Grow(v.data, newSize-currSize)
		v.data = v.data[:newSize]
		if len(value) > 0 {
			for i := currSize; i < newSize; i++ {
				v.data[i] = value[0]
			}
		} else {
			var zero T
			for i := currSize; i < newSize; i++ {
				v.data[i] = zero
			}
		}
	}
}

// Clone 创建并返回 Vector 的一个副本（深拷贝）。
func (v *Vector[T]) Clone() *Vector[T] {
	return NewVectorFromSlice(v.data)
}

// ToSlice 将 Vector 转换为一个切片并返回。
func (v *Vector[T]) ToSlice() []T {
	return append([]T{}, v.data...)
}

// Reverse 反转 Vector 中的元素顺序。
func (v *Vector[T]) Reverse() {
	slices.Reverse(v.data)
}

// Contains 检查 Vector 是否包含指定的元素。
func Contains[T comparable](v *Vector[T], element T) bool {
	return slices.Contains(v.data, element)
}

// IndexOf 返回指定元素在 Vector 中的索引，如果不存在则返回 -1。
func IndexOf[T comparable](v *Vector[T], element T) int {
	return slices.Index(v.data, element)
}

// Sort 使用自定义比较函数对 Vector 中的元素进行排序。
// cmp 函数应返回负数、零或正数，分别表示 a < b、a == b 或 a > b。
func (v *Vector[T]) Sort(cmp func(a, b T) int) {
	slices.SortFunc(v.data, cmp)
}

// Sort 对 Vector 中可排序类型的元素进行升序排序。
// 仅适用于实现了 cmp.Ordered 接口的类型（如 int, float64, string 等）。
func Sort[T cmp.Ordered](v *Vector[T]) {
	slices.Sort(v.data)
}

// Equal 检查两个 Vector 是否相等（元素和顺序均相同）。
func Equal[T comparable](v1, v2 *Vector[T]) bool {
	return slices.Equal(v1.data, v2.data)
}
