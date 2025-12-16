package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q == nil {
		t.Fatal("NewQueue returned nil")
	}
	if q.Len() != 0 {
		t.Fatalf("expected length 0, got %d", q.Len())
	}
	if !q.IsEmpty() {
		t.Fatal("queue should be empty")
	}
}

func TestPushPopOrder(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	if q.Len() != 5 {
		t.Fatalf("expected length 5, got %d", q.Len())
	}
	for i := 0; i < 5; i++ {
		val, ok := q.Pop()
		if !ok {
			t.Fatalf("Pop #%d failed", i)
		}
		if val != i {
			t.Fatalf("Pop #%d expected %d, got %d", i, i, val)
		}
	}
	if !q.IsEmpty() {
		t.Fatal("queue should be empty after pops")
	}
	if _, ok := q.Pop(); ok {
		t.Fatal("Pop on empty queue should fail")
	}
}

func TestFrontBack(t *testing.T) {
	q := NewQueue[int]()
	q.Push(10)
	q.Push(20)
	q.Push(30)
	if front, _ := q.Front(); front != 10 {
		t.Fatalf("Front expected 10, got %d", front)
	}
	if back, _ := q.Back(); back != 30 {
		t.Fatalf("Back expected 30, got %d", back)
	}
}

func TestClear(t *testing.T) {
	q := NewQueue[string]()
	q.Push("a")
	q.Push("b")
	q.Push("c")
	q.Clear()
	if !q.IsEmpty() {
		t.Fatal("queue should be empty after Clear")
	}
	q.Push("d")
	if q.Len() != 1 {
		t.Fatalf("expected length 1 after reuse, got %d", q.Len())
	}
}

func TestZeroValueQueue(t *testing.T) {
	var q Queue[int]
	q.Push(1)
	q.Push(2)
	if q.Len() != 2 {
		t.Fatalf("expected length 2, got %d", q.Len())
	}
	if front, _ := q.Front(); front != 1 {
		t.Fatalf("Front expected 1, got %d", front)
	}
}

func TestToSliceIndependence(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 4; i++ {
		q.Push(i)
	}
	slice := q.ToSlice()
	if len(slice) != 4 {
		t.Fatalf("ToSlice expected length 4, got %d", len(slice))
	}
	for i, v := range slice {
		if v != i {
			t.Fatalf("slice[%d] expected %d, got %d", i, i, v)
		}
	}
	slice[0] = 99
	if front, _ := q.Front(); front != 0 {
		t.Fatal("modifying ToSlice result should not affect queue")
	}
}

func TestSwap(t *testing.T) {
	q1 := NewQueue[int]()
	q2 := NewQueue[int]()
	for i := 0; i < 3; i++ {
		q1.Push(i)
	}
	for i := 10; i < 13; i++ {
		q2.Push(i)
	}
	q1.Swap(q2)
	if front, _ := q1.Front(); front != 10 {
		t.Fatalf("q1 front expected 10 after swap, got %d", front)
	}
	if front, _ := q2.Front(); front != 0 {
		t.Fatalf("q2 front expected 0 after swap, got %d", front)
	}
}

func TestClone(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	clone := q.Clone()
	if clone == q {
		t.Fatal("Clone should return different instance")
	}
	if !Equal(q, clone) {
		t.Fatal("Clone content should equal original")
	}
	clone.Push(99)
	if Equal(q, clone) {
		t.Fatal("mutating clone should not keep equality")
	}
}

func TestEqual(t *testing.T) {
	a := NewQueue[int]()
	b := NewQueue[int]()
	if !Equal(a, b) {
		t.Fatal("two empty queues should be equal")
	}
	for i := 0; i < 3; i++ {
		a.Push(i)
		b.Push(i)
	}
	if !Equal(a, b) {
		t.Fatal("queues with same content should be equal")
	}
	b.Push(99)
	if Equal(a, b) {
		t.Fatal("queues with different lengths should not be equal")
	}
	a.Push(100)
	if Equal(a, b) {
		t.Fatal("queues with different elements should not be equal")
	}
}
