package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	if s == nil {
		t.Fatal("NewStack returned nil")
	}
	if s.Len() != 0 || !s.IsEmpty() {
		t.Fatalf("expected new stack empty, len=%d", s.Len())
	}
}

func TestPushPop(t *testing.T) {
	s := NewStack[int]()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	if s.Len() != 5 {
		t.Fatalf("expected len 5, got %d", s.Len())
	}
	for i := 4; i >= 0; i-- {
		val, ok := s.Pop()
		if !ok || val != i {
			t.Fatalf("Pop expected (%d,true), got (%v,%v)", i, val, ok)
		}
	}
	if !s.IsEmpty() {
		t.Fatal("stack should be empty after pops")
	}
	if _, ok := s.Pop(); ok {
		t.Fatal("Pop on empty stack should fail")
	}
}

func TestTop(t *testing.T) {
	s := NewStack[string]()
	s.Push("a")
	s.Push("b")
	if top, _ := s.Top(); top != "b" {
		t.Fatalf("Top expected b, got %s", top)
	}
	s.Pop()
	if top, _ := s.Top(); top != "a" {
		t.Fatalf("Top expected a, got %s", top)
	}
}

func TestClear(t *testing.T) {
	s := NewStack[int]()
	for i := 0; i < 3; i++ {
		s.Push(i)
	}
	s.Clear()
	if !s.IsEmpty() {
		t.Fatal("stack should be empty after Clear")
	}
	s.Push(42)
	if s.Len() != 1 {
		t.Fatalf("expected len 1 after reuse, got %d", s.Len())
	}
}

func TestZeroValueStack(t *testing.T) {
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	if s.Len() != 2 {
		t.Fatalf("expected len 2, got %d", s.Len())
	}
	if top, _ := s.Top(); top != 20 {
		t.Fatalf("Top expected 20, got %d", top)
	}
}

func TestToSliceIndependence(t *testing.T) {
	s := NewStack[int]()
	for i := 0; i < 4; i++ {
		s.Push(i)
	}
	slice := s.ToSlice()
	if len(slice) != 4 {
		t.Fatalf("ToSlice expected len 4, got %d", len(slice))
	}
	for i, v := range slice {
		if v != i {
			t.Fatalf("slice[%d] expected %d, got %d", i, i, v)
		}
	}
	slice[0] = 99
	if bottom := s.ToSlice()[0]; bottom != 0 {
		t.Fatal("modifying ToSlice result should not affect stack data")
	}
}

func TestSwap(t *testing.T) {
	s1 := NewStack[int]()
	s2 := NewStack[int]()
	for i := 0; i < 3; i++ {
		s1.Push(i)
	}
	for i := 10; i < 13; i++ {
		s2.Push(i)
	}
	s1.Swap(s2)
	if top, _ := s1.Top(); top != 12 {
		t.Fatalf("s1 top expected 12 after swap, got %d", top)
	}
	if top, _ := s2.Top(); top != 2 {
		t.Fatalf("s2 top expected 2 after swap, got %d", top)
	}
}

func TestClone(t *testing.T) {
	s := NewStack[int]()
	for i := 0; i < 4; i++ {
		s.Push(i)
	}
	clone := s.Clone()
	if clone == s {
		t.Fatal("Clone should return new instance")
	}
	if !Equal(s, clone) {
		t.Fatal("Clone content should equal original")
	}
	clone.Push(99)
	if Equal(s, clone) {
		t.Fatal("mutating clone should not keep equality")
	}
}

func TestEqual(t *testing.T) {
	a := NewStack[int]()
	b := NewStack[int]()
	if !Equal(a, b) {
		t.Fatal("two empty stacks should be equal")
	}
	for i := 0; i < 3; i++ {
		a.Push(i)
		b.Push(i)
	}
	if !Equal(a, b) {
		t.Fatal("stacks with same content should be equal")
	}
	b.Push(99)
	if Equal(a, b) {
		t.Fatal("stacks with different lengths should not be equal")
	}
	a.Push(100)
	if Equal(a, b) {
		t.Fatal("stacks with different elements should not be equal")
	}
}

func TestGenericStack(t *testing.T) {
	type point struct {
		X int
		Y int
	}
	s := NewStack[point]()
	s.Push(point{X: 1, Y: 2})
	s.Push(point{X: 3, Y: 4})
	if top, _ := s.Top(); top.X != 3 || top.Y != 4 {
		t.Fatalf("Top returned unexpected value: %+v", top)
	}
}
