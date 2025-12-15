// Package deque 提供了泛型双端队列的实现，采用分段存储方式。
package deque

import "fmt"

const (
	// blockSize 定义了每个存储块的大小。
	blockSize = 128
	// initialMapSize 中央 map 的初始大小。
	initialMapSize = 8
)

// Deque 是一个泛型双端队列，采用分段存储方式实现。
type Deque[T any] struct {
	mapData  [][]T // 中央 map：每个元素是一个数据块
	mapSize  int   // 中央 map 的当前大小
	mapStart int   // 第一个有效块在 map 中的索引
	mapEnd   int   // 最后一个有效块的下一个索引

	headBlock  int // 头部块的索引（相对于 mapStart）
	headOffset int // 头部元素在头部块中的偏移

	tailBlock  int // 尾部块的索引（相对于 mapStart）
	tailOffset int // 尾部元素在尾部块中的偏移

	size int // 双端队列中元素的总数
}

// NewDeque 创建并返回一个新的空 Deque。
func NewDeque[T any]() *Deque[T] {
	d := &Deque[T]{
		mapSize:  initialMapSize,
		mapData:  make([][]T, initialMapSize),
		mapStart: initialMapSize / 2, // 从中间开始，方便两端扩展
		mapEnd:   initialMapSize/2 + 1,
		size:     0,
	}

	// 分配第一个数据块
	d.mapData[d.mapStart] = make([]T, blockSize)

	// 头尾指向第一个块的中间位置
	d.headBlock = 0
	d.tailBlock = 0
	d.headOffset = blockSize / 2
	d.tailOffset = blockSize / 2

	return d
}

// Len 返回 Deque 中元素的数量。
func (d *Deque[T]) Len() int {
	return d.size
}

// IsEmpty 检查 Deque 是否为空。
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Clear 清空 Deque 中的所有元素。
func (d *Deque[T]) Clear() {
	// 重新初始化 Deque
	*d = *NewDeque[T]()
}

// PushBack 在 Deque 的尾部添加一个元素。
func (d *Deque[T]) PushBack(elem T) {
	// 情况 1: deque 为空
	if d.IsEmpty() {
		d.mapData[d.mapStart+d.tailBlock][d.tailOffset] = elem
		d.size++
		return
	}

	// 先移动 tail 指针到下一个位置
	if d.tailOffset < blockSize-1 {
		// 当前块还有空间
		d.tailOffset++
	} else {
		// 当前块已满，需要新块
		// 检查是否需要扩展中央 map
		if d.mapStart+d.tailBlock+1 >= len(d.mapData) {
			d.expandMap()
		}

		// 移到下一个块
		d.tailBlock++
		d.tailOffset = 0

		// 如果下一个块还没分配，分配一个新块
		blockIndex := d.mapStart + d.tailBlock
		if blockIndex >= len(d.mapData) {
			panic(fmt.Sprintf("blockIndex out of range: %d (mapSize=%d)", blockIndex, len(d.mapData)))
		}
		if d.mapData[blockIndex] == nil {
			d.mapData[blockIndex] = make([]T, blockSize)
		}

		// 更新 mapEnd（新块已使用）
		if blockIndex+1 > d.mapEnd {
			d.mapEnd = blockIndex + 1
		}
	}

	// 在新位置插入元素
	d.mapData[d.mapStart+d.tailBlock][d.tailOffset] = elem
	d.size++
}

// PushFront 在 Deque 的头部添加一个元素。
func (d *Deque[T]) PushFront(elem T) {
	// 情况 1: deque 为空
	if d.IsEmpty() {
		d.mapData[d.mapStart+d.headBlock][d.headOffset] = elem
		d.size++
		return
	}

	// 先移动 head 指针到前一个位置
	if d.headOffset > 0 {
		// 当前块还有空间（向前）
		d.headOffset--
	} else {
		// 当前块已满，需要前一个块
		// 检查是否需要扩展中央 map
		if d.mapStart+d.headBlock-1 < 0 {
			d.expandMap()
		}

		// 移到前一个块
		d.headBlock--
		d.headOffset = blockSize - 1

		// 如果前一个块还没分配，分配一个新块
		blockIndex := d.mapStart + d.headBlock
		if blockIndex < 0 || blockIndex >= len(d.mapData) {
			panic(fmt.Sprintf("blockIndex out of range: %d (mapStart=%d, headBlock=%d)", blockIndex, d.mapStart, d.headBlock))
		}
		if d.mapData[blockIndex] == nil {
			d.mapData[blockIndex] = make([]T, blockSize)
		}
	}

	// 在新位置插入元素
	d.mapData[d.mapStart+d.headBlock][d.headOffset] = elem
	d.size++
}

// expandMap 当 map 空间不足时，扩展中央 map。
func (d *Deque[T]) expandMap() {
	// 计算新的 map 大小（翻倍）
	newMapSize := d.mapSize * 2
	newMapData := make([][]T, newMapSize)

	// 计算新的起始位置（居中）
	newMapStart := (newMapSize - (d.mapEnd - d.mapStart)) / 2

	// 复制现有的块到新的 map
	copy(newMapData[newMapStart:], d.mapData[d.mapStart:d.mapEnd])

	// 更新字段
	d.mapData = newMapData
	d.mapSize = newMapSize
	d.mapEnd = newMapStart + (d.mapEnd - d.mapStart)
	d.mapStart = newMapStart
}

// PopBack 从 Deque 的尾部移除并返回一个元素。
// 如果 Deque 为空，返回零值和 false。
func (d *Deque[T]) PopBack() (T, bool) {
	var zero T
	if d.IsEmpty() {
		return zero, false
	}

	// 获取尾部元素
	elem := d.mapData[d.mapStart+d.tailBlock][d.tailOffset]

	// 清零防止内存泄漏
	d.mapData[d.mapStart+d.tailBlock][d.tailOffset] = zero

	// 如果删除后为空
	if d.size == 1 {
		d.size = 0
		d.headBlock = 0
		d.tailBlock = 0
		d.headOffset = blockSize / 2
		d.tailOffset = blockSize / 2
		return elem, true
	}

	// 减少大小
	d.size--

	// 向前移动 tail 指针
	if d.tailOffset > 0 {
		// 当前块还有元素，向前移动
		d.tailOffset--
	} else {
		// 当前块已空，移到前一个块的末尾
		d.tailBlock--
		d.tailOffset = blockSize - 1
	}

	return elem, true
}

// PopFront 从 Deque 的头部移除并返回一个元素。
// 如果 Deque 为空，返回零值和 false。
func (d *Deque[T]) PopFront() (T, bool) {
	var zero T
	if d.IsEmpty() {
		return zero, false
	}

	// 获取头部元素
	elem := d.mapData[d.mapStart+d.headBlock][d.headOffset]

	// 清零防止内存泄漏
	d.mapData[d.mapStart+d.headBlock][d.headOffset] = zero

	// 如果删除后为空
	if d.size == 1 {
		d.size = 0
		d.headBlock = 0
		d.tailBlock = 0
		d.headOffset = blockSize / 2
		d.tailOffset = blockSize / 2
		return elem, true
	}

	// 减少大小
	d.size--

	// 向后移动 head 指针
	if d.headOffset < blockSize-1 {
		// 当前块还有元素，向后移动
		d.headOffset++
	} else {
		// 当前块已空，移到下一个块的开头
		d.headBlock++
		d.headOffset = 0
	}

	return elem, true
}

// Front 返回 Deque 头部的元素但不移除它。
// 如果 Deque 为空，返回零值和 false。
func (d *Deque[T]) Front() (T, bool) {
	var zero T
	if d.IsEmpty() {
		return zero, false
	}
	return d.mapData[d.mapStart+d.headBlock][d.headOffset], true
}

// Back 返回 Deque 尾部的元素但不移除它。
// 如果 Deque 为空，返回零值和 false。
func (d *Deque[T]) Back() (T, bool) {
	var zero T
	if d.IsEmpty() {
		return zero, false
	}
	return d.mapData[d.mapStart+d.tailBlock][d.tailOffset], true
}

// At 返回指定索引处的元素。
// 索引从 0 开始，不检查边界，越界会引发 panic。
func (d *Deque[T]) At(index int) T {
	// 计算元素在哪个块和块内偏移
	absoluteIndex := d.headOffset + index
	blockIndex := d.headBlock + absoluteIndex/blockSize
	offset := absoluteIndex % blockSize

	return d.mapData[d.mapStart+blockIndex][offset]
}

// Get 安全地返回指定索引处的元素。
// 如果索引无效，返回零值和 false。
func (d *Deque[T]) Get(index int) (T, bool) {
	var zero T
	if index < 0 || index >= d.size {
		return zero, false
	}
	return d.At(index), true
}

// Set 设置指定索引处的元素的值。
// 如果索引无效返回 false。
func (d *Deque[T]) Set(index int, value T) bool {
	if index < 0 || index >= d.size {
		return false
	}

	// 计算元素在哪个块和块内偏移
	absoluteIndex := d.headOffset + index
	blockIndex := d.headBlock + absoluteIndex/blockSize
	offset := absoluteIndex % blockSize

	d.mapData[d.mapStart+blockIndex][offset] = value
	return true
}

// Clone 创建并返回 Deque 的一个深拷贝。
func (d *Deque[T]) Clone() *Deque[T] {
	clone := NewDeque[T]()
	for i := 0; i < d.size; i++ {
		clone.PushBack(d.At(i))
	}
	return clone
}

// Toslice 将 Deque 转换为一个切片并返回。
func (d *Deque[T]) ToSlice() []T {
	if d.IsEmpty() {
		return []T{}
	}

	result := make([]T, d.size)
	for i := 0; i < d.size; i++ {
		result[i] = d.At(i)
	}
	return result
}

// Insert 在指定索引处插入一个元素。
// 如果索引无效返回 false。
func (d *Deque[T]) Insert(index int, elem T) bool {
	// 边界检查：允许在末尾插入
	if index < 0 || index > d.size {
		return false
	}

	// 特殊情况优化
	if index == 0 {
		d.PushFront(elem)
		return true
	}
	if index == d.size {
		d.PushBack(elem)
		return true
	}

	// 决定移动方向：前半部分还是后半部分
	if index < d.size/2 {
		// 向前移动元素
		// 先在头部添加一个占位符
		d.PushFront(d.At(0))
		for i := 0; i < index; i++ {
			d.Set(i, d.At(i+1)) // 将元素向前移动一位
		}
		d.Set(index, elem) // 插入新元素
	} else {
		// 向后移动元素
		// 先在尾部添加一个占位符
		d.PushBack(d.At(d.size - 1))
		for i := d.size - 2; i > index; i-- {
			d.Set(i, d.At(i-1)) // 将元素向后移动一位
		}
		d.Set(index, elem) // 插入新元素
	}

	return true
}

// Erase 删除 [start, end) 的元素区间。
// 如果索引无效返回 false。
func (d *Deque[T]) Erase(start, end int) bool {
	// 边界检查
	if start < 0 || end > d.size || start >= end {
		return false
	}

	count := end - start // 要删除的元素数量

	// 决定移动方向
	if start < d.size-end {
		// 向后移动前面的元素
		for i := start - 1; i >= 0; i-- {
			d.Set(i+count, d.At(i))
		}
		// 更新头部指针
		for i := 0; i < count; i++ {
			d.PopFront()
		}
	} else {
		// 向前移动后面的元素
		for i := end; i < d.size; i++ {
			d.Set(i-count, d.At(i))
		}
		// 更新尾部指针
		for i := 0; i < count; i++ {
			d.PopBack()
		}
	}

	return true
}

// Resize 调整 Deque 的大小。
// 如果新大小大于当前大小，使用零值填充新元素。
// 如果提供了 fillValue，则使用该值填充新元素。
// 如果新大小小于当前大小，删除多余的元素。
func (d *Deque[T]) Resize(newSize int, fillValue ...T) {
	if newSize < 0 {
		return
	}

	if newSize > d.size {
		// 扩展 Deque
		var value T
		if len(fillValue) > 0 {
			value = fillValue[0]
		}

		for i := d.size; i < newSize; i++ {
			d.PushBack(value)
		}
	} else if newSize < d.size {
		// 收缩 Deque
		for i := d.size; i > newSize; i-- {
			d.PopBack()
		}
	}
}

// Contains 检查 Deque 是否包含指定的元素。
// 需要类型 T 支持相等比较。
func Contains[T comparable](d *Deque[T], elem T) bool {
	for i := 0; i < d.size; i++ {
		if d.At(i) == elem {
			return true
		}
	}
	return false
}

// IndexOf 返回元素在 Deque 中第一次出现的索引。
// 如果元素不存在，返回 -1。
// 需要类型 T 支持相等比较。
func IndexOf[T comparable](d *Deque[T], elem T) int {
	for i := 0; i < d.size; i++ {
		if d.At(i) == elem {
			return i
		}
	}
	return -1
}

// Equals 检查两个 Deque 是否相等。
// 需要类型 T 支持相等比较。
func Equals[T comparable](d1, d2 *Deque[T]) bool {
	if d1.size != d2.size {
		return false
	}

	for i := 0; i < d1.size; i++ {
		if d1.At(i) != d2.At(i) {
			return false
		}
	}

	return true
}

// Reverse 反转 Deque 中的元素顺序。
func (d *Deque[T]) Reverse() {
	if d.size <= 1 {
		return
	}

	// 从两端向中间交换元素
	for i := 0; i < d.size/2; i++ {
		j := d.size - 1 - i
		temp := d.At(i)
		d.Set(i, d.At(j))
		d.Set(j, temp)
	}
}

// Swap 交换两个 Deque 的内容。
func (d *Deque[T]) Swap(other *Deque[T]) {
	*d, *other = *other, *d
}
