package deque

import (
	"testing"
)

// TestNewDeque 测试创建新的 Deque
func TestNewDeque(t *testing.T) {
	d := NewDeque[int]()
	if d == nil {
		t.Fatal("NewDeque returned nil")
	}
	if d.Len() != 0 {
		t.Errorf("Expected length 0, got %d", d.Len())
	}
	if !d.IsEmpty() {
		t.Error("Expected deque to be empty")
	}
}

// TestPushBackPopBack 测试尾部插入和删除
func TestPushBackPopBack(t *testing.T) {
	d := NewDeque[int]()

	// 测试单个元素
	d.PushBack(1)
	if d.Len() != 1 {
		t.Errorf("Expected length 1, got %d", d.Len())
	}

	val, ok := d.Back()
	if !ok || val != 1 {
		t.Errorf("Expected Back() to return 1, got %v, %v", val, ok)
	}

	val, ok = d.PopBack()
	if !ok || val != 1 {
		t.Errorf("Expected PopBack() to return 1, got %v, %v", val, ok)
	}
	if d.Len() != 0 {
		t.Errorf("Expected length 0 after pop, got %d", d.Len())
	}

	// 测试多个元素
	for i := 0; i < 200; i++ {
		d.PushBack(i)
	}
	if d.Len() != 200 {
		t.Errorf("Expected length 200, got %d", d.Len())
	}

	for i := 199; i >= 0; i-- {
		val, ok = d.PopBack()
		if !ok || val != i {
			t.Errorf("Expected PopBack() to return %d, got %v, %v", i, val, ok)
		}
	}
	if !d.IsEmpty() {
		t.Error("Expected deque to be empty after popping all elements")
	}
}

// TestPushFrontPopFront 测试头部插入和删除
func TestPushFrontPopFront(t *testing.T) {
	d := NewDeque[int]()

	// 测试单个元素
	d.PushFront(1)
	if d.Len() != 1 {
		t.Errorf("Expected length 1, got %d", d.Len())
	}

	val, ok := d.Front()
	if !ok || val != 1 {
		t.Errorf("Expected Front() to return 1, got %v, %v", val, ok)
	}

	val, ok = d.PopFront()
	if !ok || val != 1 {
		t.Errorf("Expected PopFront() to return 1, got %v, %v", val, ok)
	}
	if d.Len() != 0 {
		t.Errorf("Expected length 0 after pop, got %d", d.Len())
	}

	// 测试多个元素
	for i := 0; i < 200; i++ {
		d.PushFront(i)
	}
	if d.Len() != 200 {
		t.Errorf("Expected length 200, got %d", d.Len())
	}

	for i := 199; i >= 0; i-- {
		val, ok = d.PopFront()
		if !ok || val != i {
			t.Errorf("Expected PopFront() to return %d, got %v, %v", i, val, ok)
		}
	}
	if !d.IsEmpty() {
		t.Error("Expected deque to be empty after popping all elements")
	}
}

// TestMixedOperations 测试混合操作
func TestMixedOperations(t *testing.T) {
	d := NewDeque[int]()

	// 混合 PushBack 和 PushFront
	d.PushBack(1)   // [1]
	d.PushFront(0)  // [0, 1]
	d.PushBack(2)   // [0, 1, 2]
	d.PushFront(-1) // [-1, 0, 1, 2]

	if d.Len() != 4 {
		t.Errorf("Expected length 4, got %d", d.Len())
	}

	// 验证元素顺序
	expected := []int{-1, 0, 1, 2}
	for i, exp := range expected {
		if d.At(i) != exp {
			t.Errorf("At(%d) expected %d, got %d", i, exp, d.At(i))
		}
	}

	// 混合删除
	val, _ := d.PopFront() // [0, 1, 2]
	if val != -1 {
		t.Errorf("Expected PopFront() to return -1, got %d", val)
	}

	val, _ = d.PopBack() // [0, 1]
	if val != 2 {
		t.Errorf("Expected PopBack() to return 2, got %d", val)
	}

	if d.Len() != 2 {
		t.Errorf("Expected length 2, got %d", d.Len())
	}
}

// TestAtGetSet 测试随机访问
func TestAtGetSet(t *testing.T) {
	d := NewDeque[int]()

	// 添加元素
	for i := 0; i < 10; i++ {
		d.PushBack(i)
	}

	// 测试 At
	for i := 0; i < 10; i++ {
		if d.At(i) != i {
			t.Errorf("At(%d) expected %d, got %d", i, i, d.At(i))
		}
	}

	// 测试 Get
	for i := 0; i < 10; i++ {
		val, ok := d.Get(i)
		if !ok || val != i {
			t.Errorf("Get(%d) expected %d, got %v, %v", i, i, val, ok)
		}
	}

	// 测试越界
	_, ok := d.Get(-1)
	if ok {
		t.Error("Get(-1) should return false")
	}
	_, ok = d.Get(10)
	if ok {
		t.Error("Get(10) should return false")
	}

	// 测试 Set
	for i := 0; i < 10; i++ {
		if !d.Set(i, i*10) {
			t.Errorf("Set(%d, %d) failed", i, i*10)
		}
	}

	// 验证 Set 的结果
	for i := 0; i < 10; i++ {
		if d.At(i) != i*10 {
			t.Errorf("After Set, At(%d) expected %d, got %d", i, i*10, d.At(i))
		}
	}

	// 测试 Set 越界
	if d.Set(-1, 999) {
		t.Error("Set(-1) should return false")
	}
	if d.Set(10, 999) {
		t.Error("Set(10) should return false")
	}
}

// TestClear 测试清空操作
func TestClear(t *testing.T) {
	d := NewDeque[int]()

	for i := 0; i < 100; i++ {
		d.PushBack(i)
	}

	d.Clear()

	if !d.IsEmpty() {
		t.Error("Expected deque to be empty after Clear()")
	}
	if d.Len() != 0 {
		t.Errorf("Expected length 0 after Clear(), got %d", d.Len())
	}

	// 清空后应该还能正常使用
	d.PushBack(1)
	if d.Len() != 1 {
		t.Errorf("Expected length 1 after Clear() and PushBack(), got %d", d.Len())
	}
}

// TestClone 测试克隆
func TestClone(t *testing.T) {
	d := NewDeque[int]()

	for i := 0; i < 50; i++ {
		d.PushBack(i)
	}

	clone := d.Clone()

	// 验证长度
	if clone.Len() != d.Len() {
		t.Errorf("Clone length expected %d, got %d", d.Len(), clone.Len())
	}

	// 验证内容
	for i := 0; i < d.Len(); i++ {
		if clone.At(i) != d.At(i) {
			t.Errorf("Clone At(%d) expected %d, got %d", i, d.At(i), clone.At(i))
		}
	}

	// 修改克隆不应影响原始
	clone.Set(0, 999)
	if d.At(0) == 999 {
		t.Error("Modifying clone affected original deque")
	}
}

// TestToSlice 测试转换为切片
func TestToSlice(t *testing.T) {
	d := NewDeque[int]()

	// 空 deque
	slice := d.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Empty deque ToSlice() expected length 0, got %d", len(slice))
	}

	// 添加元素
	for i := 0; i < 20; i++ {
		d.PushBack(i)
	}

	slice = d.ToSlice()
	if len(slice) != 20 {
		t.Errorf("ToSlice() expected length 20, got %d", len(slice))
	}

	// 验证内容
	for i := 0; i < 20; i++ {
		if slice[i] != i {
			t.Errorf("ToSlice()[%d] expected %d, got %d", i, i, slice[i])
		}
	}

	// 修改切片不应影响原始 deque
	slice[0] = 999
	if d.At(0) == 999 {
		t.Error("Modifying slice affected original deque")
	}
}

// TestFrontBackEmpty 测试空 deque 的 Front 和 Back
func TestFrontBackEmpty(t *testing.T) {
	d := NewDeque[int]()

	_, ok := d.Front()
	if ok {
		t.Error("Front() on empty deque should return false")
	}

	_, ok = d.Back()
	if ok {
		t.Error("Back() on empty deque should return false")
	}
}

// TestPopEmpty 测试空 deque 的 Pop 操作
func TestPopEmpty(t *testing.T) {
	d := NewDeque[int]()

	_, ok := d.PopFront()
	if ok {
		t.Error("PopFront() on empty deque should return false")
	}

	_, ok = d.PopBack()
	if ok {
		t.Error("PopBack() on empty deque should return false")
	}
}

// TestLargeDeque 测试大量元素
func TestLargeDeque(t *testing.T) {
	d := NewDeque[int]()
	n := 10000

	// 从尾部添加大量元素
	for i := 0; i < n; i++ {
		d.PushBack(i)
	}

	if d.Len() != n {
		t.Errorf("Expected length %d, got %d", n, d.Len())
	}

	// 验证元素
	for i := 0; i < n; i++ {
		if d.At(i) != i {
			t.Errorf("At(%d) expected %d, got %d", i, i, d.At(i))
		}
	}

	// 从头部删除所有元素
	for i := 0; i < n; i++ {
		val, ok := d.PopFront()
		if !ok || val != i {
			t.Errorf("PopFront() expected %d, got %v, %v", i, val, ok)
		}
	}

	if !d.IsEmpty() {
		t.Error("Expected deque to be empty")
	}
}

// TestAlternatingOperations 测试交替操作
func TestAlternatingOperations(t *testing.T) {
	d := NewDeque[int]()

	// 交替从两端添加
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			d.PushBack(i)
		} else {
			d.PushFront(i)
		}
	}

	if d.Len() != 100 {
		t.Errorf("Expected length 100, got %d", d.Len())
	}

	// 交替从两端删除
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			_, ok := d.PopFront()
			if !ok {
				t.Errorf("PopFront() failed at iteration %d", i)
			}
		} else {
			_, ok := d.PopBack()
			if !ok {
				t.Errorf("PopBack() failed at iteration %d", i)
			}
		}
	}

	if !d.IsEmpty() {
		t.Error("Expected deque to be empty after alternating pops")
	}
}

// TestStringDeque 测试字符串类型的 Deque
func TestStringDeque(t *testing.T) {
	d := NewDeque[string]()

	words := []string{"hello", "world", "foo", "bar"}
	for _, word := range words {
		d.PushBack(word)
	}

	if d.Len() != len(words) {
		t.Errorf("Expected length %d, got %d", len(words), d.Len())
	}

	for i, word := range words {
		if d.At(i) != word {
			t.Errorf("At(%d) expected %s, got %s", i, word, d.At(i))
		}
	}
}

// TestStructDeque 测试结构体类型的 Deque
func TestStructDeque(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	d := NewDeque[Person]()

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	for _, p := range people {
		d.PushBack(p)
	}

	if d.Len() != len(people) {
		t.Errorf("Expected length %d, got %d", len(people), d.Len())
	}

	for i, p := range people {
		got := d.At(i)
		if got.Name != p.Name || got.Age != p.Age {
			t.Errorf("At(%d) expected %v, got %v", i, p, got)
		}
	}
}

// BenchmarkPushBack 基准测试：尾部插入
func BenchmarkPushBack(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
}

// BenchmarkPushFront 基准测试：头部插入
func BenchmarkPushFront(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
	}
}

// BenchmarkPopBack 基准测试：尾部删除
func BenchmarkPopBack(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PopBack()
	}
}

// BenchmarkPopFront 基准测试：头部删除
func BenchmarkPopFront(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PopFront()
	}
}

// BenchmarkRandomAccess 基准测试：随机访问
func BenchmarkRandomAccess(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < 10000; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.At(i % 10000)
	}
}

// TestInsert 测试插入功能
func TestInsert(t *testing.T) {
	d := NewDeque[int]()

	// 空deque插入
	if !d.Insert(0, 100) {
		t.Error("Insert at index 0 in empty deque should succeed")
	}
	if d.Len() != 1 || d.At(0) != 100 {
		t.Error("Insert at index 0 failed")
	}

	// 头部插入
	d.Clear()
	for i := 0; i < 5; i++ {
		d.PushBack(i) // [0, 1, 2, 3, 4]
	}
	if !d.Insert(0, 99) {
		t.Error("Insert at front should succeed")
	}
	if d.At(0) != 99 || d.Len() != 6 {
		t.Error("Insert at front failed")
	}

	// 尾部插入
	if !d.Insert(d.Len(), 88) {
		t.Error("Insert at back should succeed")
	}
	if d.At(d.Len()-1) != 88 || d.Len() != 7 {
		t.Error("Insert at back failed")
	}

	// 中间插入
	d.Clear()
	for i := 0; i < 10; i++ {
		d.PushBack(i) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	}
	if !d.Insert(5, 99) { // [0, 1, 2, 3, 4, 99, 5, 6, 7, 8, 9]
		t.Error("Insert at middle should succeed")
	}
	if d.Len() != 11 {
		t.Errorf("Expected length 11, got %d", d.Len())
	}
	if d.At(5) != 99 {
		t.Errorf("Expected 99 at index 5, got %d", d.At(5))
	}
	if d.At(6) != 5 {
		t.Errorf("Expected 5 at index 6, got %d", d.At(6))
	}

	// 越界插入
	if d.Insert(-1, 1) {
		t.Error("Insert at negative index should fail")
	}
	if d.Insert(100, 1) {
		t.Error("Insert at out-of-bounds index should fail")
	}
}

// TestErase 测试删除功能
func TestErase(t *testing.T) {
	d := NewDeque[int]()

	// 空deque删除
	if d.Erase(0, 1) {
		t.Error("Erase from empty deque should fail")
	}

	// 添加元素 [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	for i := 0; i < 10; i++ {
		d.PushBack(i)
	}

	// 删除头部
	if !d.Erase(0, 2) { // 删除 [0, 1]，剩余 [2, 3, 4, 5, 6, 7, 8, 9]
		t.Error("Erase from front should succeed")
	}
	if d.Len() != 8 || d.At(0) != 2 {
		t.Errorf("Erase from front failed: len=%d, first=%d", d.Len(), d.At(0))
	}

	// 删除尾部
	if !d.Erase(6, 8) { // 删除 [8, 9]，剩余 [2, 3, 4, 5, 6, 7]
		t.Error("Erase from back should succeed")
	}
	if d.Len() != 6 || d.At(d.Len()-1) != 7 {
		t.Errorf("Erase from back failed: len=%d, last=%d", d.Len(), d.At(d.Len()-1))
	}

	// 删除中间
	if !d.Erase(2, 4) { // 删除 [4, 5]，剩余 [2, 3, 6, 7]
		t.Error("Erase from middle should succeed")
	}
	if d.Len() != 4 {
		t.Errorf("Expected length 4, got %d", d.Len())
	}
	expected := []int{2, 3, 6, 7}
	for i, exp := range expected {
		if d.At(i) != exp {
			t.Errorf("At(%d) expected %d, got %d", i, exp, d.At(i))
		}
	}

	// 越界删除
	if d.Erase(-1, 1) {
		t.Error("Erase with negative start should fail")
	}
	if d.Erase(0, 100) {
		t.Error("Erase with out-of-bounds end should fail")
	}
	if d.Erase(3, 2) {
		t.Error("Erase with start >= end should fail")
	}
}

// TestResize 测试调整大小功能
func TestResize(t *testing.T) {
	d := NewDeque[int]()

	// 扩大（零值填充）
	d.Resize(5)
	if d.Len() != 5 {
		t.Errorf("Expected length 5, got %d", d.Len())
	}
	for i := 0; i < 5; i++ {
		if d.At(i) != 0 {
			t.Errorf("At(%d) expected 0, got %d", i, d.At(i))
		}
	}

	// 扩大（指定值填充）
	d.Resize(8, 42)
	if d.Len() != 8 {
		t.Errorf("Expected length 8, got %d", d.Len())
	}
	for i := 5; i < 8; i++ {
		if d.At(i) != 42 {
			t.Errorf("At(%d) expected 42, got %d", i, d.At(i))
		}
	}

	// 缩小
	d.Resize(3)
	if d.Len() != 3 {
		t.Errorf("Expected length 3, got %d", d.Len())
	}
	// 前3个元素应该保留（都是0）
	for i := 0; i < 3; i++ {
		if d.At(i) != 0 {
			t.Errorf("At(%d) expected 0, got %d", i, d.At(i))
		}
	}

	// 调整为0
	d.Resize(0)
	if d.Len() != 0 || !d.IsEmpty() {
		t.Error("Resize to 0 should make deque empty")
	}

	// 负数大小
	d.Resize(-5)
	if d.Len() != 0 {
		t.Error("Resize with negative size should be ignored")
	}
}

// TestContains 测试包含检查
func TestContains(t *testing.T) {
	d := NewDeque[int]()

	// 空deque
	if Contains(d, 1) {
		t.Error("Contains should return false for empty deque")
	}

	// 添加元素
	for i := 0; i < 10; i++ {
		d.PushBack(i)
	}

	// 存在的元素
	if !Contains(d, 5) {
		t.Error("Contains should return true for existing element")
	}

	// 不存在的元素
	if Contains(d, 100) {
		t.Error("Contains should return false for non-existing element")
	}

	// 字符串类型
	ds := NewDeque[string]()
	ds.PushBack("hello")
	ds.PushBack("world")
	if !Contains(ds, "hello") {
		t.Error("Contains should return true for 'hello'")
	}
	if Contains(ds, "foo") {
		t.Error("Contains should return false for 'foo'")
	}
}

// TestIndexOf 测试查找索引
func TestIndexOf(t *testing.T) {
	d := NewDeque[int]()

	// 空deque
	if IndexOf(d, 1) != -1 {
		t.Error("IndexOf should return -1 for empty deque")
	}

	// 添加元素 [1, 2, 3, 2, 5]
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	d.PushBack(2)
	d.PushBack(5)

	// 找到第一个出现
	if idx := IndexOf(d, 2); idx != 1 {
		t.Errorf("IndexOf(2) expected 1, got %d", idx)
	}

	// 不存在的元素
	if idx := IndexOf(d, 100); idx != -1 {
		t.Errorf("IndexOf(100) expected -1, got %d", idx)
	}

	// 第一个和最后一个元素
	if idx := IndexOf(d, 1); idx != 0 {
		t.Errorf("IndexOf(1) expected 0, got %d", idx)
	}
	if idx := IndexOf(d, 5); idx != 4 {
		t.Errorf("IndexOf(5) expected 4, got %d", idx)
	}
}

// TestEquals 测试相等性比较
func TestEquals(t *testing.T) {
	d1 := NewDeque[int]()
	d2 := NewDeque[int]()

	// 两个空deque
	if !Equals(d1, d2) {
		t.Error("Two empty deques should be equal")
	}

	// 相同内容
	for i := 0; i < 5; i++ {
		d1.PushBack(i)
		d2.PushBack(i)
	}
	if !Equals(d1, d2) {
		t.Error("Deques with same content should be equal")
	}

	// 不同长度
	d2.PushBack(99)
	if Equals(d1, d2) {
		t.Error("Deques with different lengths should not be equal")
	}

	// 长度相同但内容不同
	d1.PushBack(100)
	if Equals(d1, d2) {
		t.Error("Deques with different content should not be equal")
	}
}

// TestReverse 测试反转
func TestReverse(t *testing.T) {
	d := NewDeque[int]()

	// 空deque反转
	d.Reverse()
	if d.Len() != 0 {
		t.Error("Reversing empty deque should work")
	}

	// 单个元素
	d.PushBack(1)
	d.Reverse()
	if d.At(0) != 1 {
		t.Error("Reversing single element deque should not change it")
	}

	// 多个元素
	d.Clear()
	for i := 0; i < 10; i++ {
		d.PushBack(i) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	}
	d.Reverse() // [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]

	for i := 0; i < 10; i++ {
		expected := 9 - i
		if d.At(i) != expected {
			t.Errorf("At(%d) expected %d, got %d", i, expected, d.At(i))
		}
	}

	// 反转两次应该恢复原状
	d.Reverse()
	for i := 0; i < 10; i++ {
		if d.At(i) != i {
			t.Errorf("After double reverse, At(%d) expected %d, got %d", i, i, d.At(i))
		}
	}
}

// TestSwap 测试交换
func TestSwap(t *testing.T) {
	d1 := NewDeque[int]()
	d2 := NewDeque[int]()

	// 准备数据
	for i := 0; i < 5; i++ {
		d1.PushBack(i)
	}
	for i := 10; i < 13; i++ {
		d2.PushBack(i)
	}

	// 记录原始状态
	d1Len, d2Len := d1.Len(), d2.Len()
	d1First, d2First := d1.At(0), d2.At(0)

	// 交换
	d1.Swap(d2)

	// 验证交换结果
	if d1.Len() != d2Len {
		t.Errorf("After swap, d1.Len() expected %d, got %d", d2Len, d1.Len())
	}
	if d2.Len() != d1Len {
		t.Errorf("After swap, d2.Len() expected %d, got %d", d1Len, d2.Len())
	}
	if d1.At(0) != d2First {
		t.Errorf("After swap, d1.At(0) expected %d, got %d", d2First, d1.At(0))
	}
	if d2.At(0) != d1First {
		t.Errorf("After swap, d2.At(0) expected %d, got %d", d1First, d2.At(0))
	}

	// 验证完整内容
	expected1 := []int{10, 11, 12}
	expected2 := []int{0, 1, 2, 3, 4}
	for i := 0; i < d1.Len(); i++ {
		if d1.At(i) != expected1[i] {
			t.Errorf("d1.At(%d) expected %d, got %d", i, expected1[i], d1.At(i))
		}
	}
	for i := 0; i < d2.Len(); i++ {
		if d2.At(i) != expected2[i] {
			t.Errorf("d2.At(%d) expected %d, got %d", i, expected2[i], d2.At(i))
		}
	}
}
