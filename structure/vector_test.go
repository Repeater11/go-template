package structure

import (
	"cmp"
	"testing"
)

func TestConstructors(t *testing.T) {
	v1 := NewVector[int]()
	if v1.Len() != 0 || !v1.IsEmpty() {
		t.Errorf("Expected empty vector, got length %d", v1.Len())
	}

	v2 := NewVector(1, 2, 3)
	if v2.Len() == 3 {
		for i, val := range []int{1, 2, 3} {
			if v2.At(i) != val {
				t.Errorf("Expected v2[%d] to be %d, got %d", i, val, v2.At(i))
			}
		}
	} else {
		t.Errorf("Expected vector of length 3, got %d", v2.Len())
	}

	v3 := NewVectorFill(5, 7)
	if v3.Len() == 5 {
		for i := 0; i < 5; i++ {
			if v3.At(i) != 7 {
				t.Errorf("Expected v3[%d] to be 7, got %d", i, v3.At(i))
			}
		}
	} else {
		t.Errorf("Expected vector of length 5, got %d", v3.Len())
	}

	v4 := NewVectorFromSlice([]string{"a", "b", "c"})
	if v4.Len() == 3 {
		for i, val := range []string{"a", "b", "c"} {
			if v4.At(i) != val {
				t.Errorf("Expected v4[%d] to be %s, got %s", i, val, v4.At(i))
			}
		}
	} else {
		t.Errorf("Expected vector of length 3, got %d", v4.Len())
	}
}

func TestPushPopOperations(t *testing.T) {
	v := NewVector[int]()

	v.PushBack(1, 2, 3)
	if v.Len() != 3 {
		t.Errorf("Expected vector length 3 after PushBack, got %d", v.Len())
	}

	val, ok := v.PopBack()
	if !ok || val != 3 || v.Len() != 2 {
		t.Errorf("Expected to pop 3, got %d, length now %d", val, v.Len())
	}

	v.Clear()
	_, ok = v.PopBack()
	if ok {
		t.Errorf("Expected PopBack on empty vector to return false")
	}
}

func TestInsertAndErase(t *testing.T) {
	v := NewVector(1, 2, 5)

	if !v.Insert(2, 3, 4) {
		t.Errorf("Insert failed")
	}

	if v.Len() != 5 {
		t.Errorf("Expected vector length 5 after Insert, got %d", v.Len())
	}
	for i, val := range []int{1, 2, 3, 4, 5} {
		if v.At(i) != val {
			t.Errorf("Expected v[%d] to be %d, got %d", i, val, v.At(i))
		}
	}

	v.Insert(0, 0)
	if v.At(0) != 0 || v.Len() != 6 {
		t.Errorf("Expected v[0] to be 0 and length 6, got %d and length %d", v.At(0), v.Len())
	}

	v.Insert(v.Len(), 6)
	if v.At(v.Len()-1) != 6 || v.Len() != 7 {
		t.Errorf("Expected last element to be 6 and length 7, got %d and length %d", v.At(v.Len()-1), v.Len())
	}

	v = NewVector(1, 2, 3, 4, 5)
	if !v.Erase(1, 4) {
		t.Errorf("Erase failed")
	}
	if v.Len() != 2 || v.At(0) != 1 || v.At(1) != 5 {
		t.Errorf("Expected vector to be [1, 5] after Erase, got length %d", v.Len())
	}

	if v.Insert(-1, 99) || v.Insert(100, 99) {
		t.Errorf("Expected Insert with invalid index to fail")
	}
}

func TestGetSetAt(t *testing.T) {
	v := NewVector("a", "b", "c")

	if v.At(1) != "b" {
		t.Errorf("Expected v.At(1) to be 'b', got %s", v.At(1))
	}

	val, ok := v.Get(2)
	if !ok || val != "c" {
		t.Errorf("Expected Get(2) to return 'c', got %s", val)
	}

	if v.Set(0, "z") {
		if v.At(0) != "z" {
			t.Errorf("Expected v.At(0) to be 'z' after Set, got %s", v.At(0))
		}
	} else {
		t.Errorf("Set failed")
	}

	_, ok = v.Get(5)
	if ok {
		t.Errorf("Expected Get with invalid index to return false")
	}

	if v.Set(5, "x") {
		t.Errorf("Expected Set with invalid index to return false")
	}
}

func TestFrontBack(t *testing.T) {
	v := NewVector(10, 20, 30)

	front, ok := v.Front()
	if !ok || front != 10 {
		t.Errorf("Expected Front to return 10, got %d", front)
	}

	back, ok := v.Back()
	if !ok || back != 30 {
		t.Errorf("Expected Back to return 30, got %d", back)
	}

	v.Clear()
	_, ok = v.Front()
	if ok {
		t.Errorf("Expected Front on empty vector to return false")
	}

	_, ok = v.Back()
	if ok {
		t.Errorf("Expected Back on empty vector to return false")
	}
}

func TestCapacityManagement(t *testing.T) {
	v := NewVector[int]()

	v.Reserve(100)
	if v.Capacity() < 100 {
		t.Errorf("Expected capacity at least 100 after Reserve, got %d", v.Capacity())
	}
	if v.Len() != 0 {
		t.Errorf("Expected length to remain 0 after Reserve, got %d", v.Len())
	}

	v.Resize(5, 42)
	if v.Len() != 5 || v.At(0) != 42 || v.At(4) != 42 {
		t.Errorf("Expected vector of length 5 with all elements 42 after Resize, got length %d", v.Len())
	}

	v.Resize(2)
	if v.Len() != 2 {
		t.Errorf("Expected vector of length 2 after Resize down, got %d", v.Len())
	}

	v.Resize(4)
	if v.Len() != 4 || v.At(2) != 0 || v.At(3) != 0 {
		t.Errorf("Expected vector of length 4 with new elements as zero after Resize up, got length %d", v.Len())
	}
}

func TestSort(t *testing.T) {
	v1 := NewVector(3, 1, 4, 1, 5, 9, 2, 6, 5)
	Sort(v1)
	expected1 := []int{1, 1, 2, 3, 4, 5, 5, 6, 9}
	for i, val := range expected1 {
		if v1.At(i) != val {
			t.Errorf("Expected sorted v1[%d] to be %d, got %d", i, val, v1.At(i))
		}
	}

	v2 := NewVector("banana", "apple", "cherry")
	v2.Sort(func(a, b string) int { return cmp.Compare(b, a) })
	expected2 := []string{"cherry", "banana", "apple"}
	for i, val := range expected2 {
		if v2.At(i) != val {
			t.Errorf("Expected sorted v2[%d] to be %s, got %s", i, val, v2.At(i))
		}
	}
}

func TestClone(t *testing.T) {
	original := NewVector(1, 2, 3, 4, 5)
	clone := original.Clone()

	if clone.Len() != original.Len() {
		t.Errorf("Expected cloned vector length %d, got %d", original.Len(), clone.Len())
	}

	for i := 0; i < original.Len(); i++ {
		if clone.At(i) != original.At(i) {
			t.Errorf("Expected cloned vector element %d to be %d, got %d", i, original.At(i), clone.At(i))
		}
	}

	clone.Set(0, 99)
	if original.At(0) == clone.At(0) {
		t.Errorf("Modifying clone should not affect original")
	}
}

func TestEqual(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(1, 2, 3)
	v3 := NewVector(4, 5, 6)

	if !Equal(v1, v2) {
		t.Errorf("Expected v1 and v2 to be equal")
	}

	if Equal(v1, v3) {
		t.Errorf("Expected v1 and v3 to be not equal")
	}
}

func TestContainsAndIndexOf(t *testing.T) {
	v := NewVector("x", "y", "z", "x")

	if !Contains(v, "y") {
		t.Errorf("Expected vector to contain 'y'")
	}
	if Contains(v, "a") {
		t.Errorf("Expected vector to not contain 'a'")
	}

	if idx := IndexOf(v, "z"); idx != 2 {
		t.Errorf("Expected index of 'z' to be 2, got %d", idx)
	}
	if idx := IndexOf(v, "a"); idx != -1 {
		t.Errorf("Expected index of 'a' to be -1, got %d", idx)
	}
}

func TestReverse(t *testing.T) {
	v := NewVector(1, 2, 3, 4, 5)
	v.Reverse()

	expected := []int{5, 4, 3, 2, 1}
	for i, val := range expected {
		if v.At(i) != val {
			t.Errorf("Expected reversed v[%d] to be %d, got %d", i, val, v.At(i))
		}
	}
}
