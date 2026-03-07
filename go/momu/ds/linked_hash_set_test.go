package ds_test

import (
	"momu/ds"
	"testing"
)

// TestLinkedHashSet_BasicOperations 测试基本的 Insert, Contains, Size
func TestLinkedHashSet_BasicOperations(t *testing.T) {
	s := ds.NewLinkedHashSet[string]()

	// 1. 测试空状态
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
	if s.Contains("item1") {
		t.Error("Expected set to not contain item1")
	}

	// 2. 测试插入
	s.Insert("apple")
	s.Insert("banana")
	s.Insert("cherry")

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	// 3. 测试存在性
	if !s.Contains("apple") {
		t.Error("Expected Contains('apple') to be true")
	}
	if !s.Contains("banana") {
		t.Error("Expected Contains('banana') to be true")
	}

	// 4. 测试不存在性
	if s.Contains("durian") {
		t.Error("Expected Contains('durian') to be false")
	}
}

// TestLinkedHashSet_NoDuplicates 测试集合的去重特性 (核心功能)
func TestLinkedHashSet_NoDuplicates(t *testing.T) {
	s := ds.NewLinkedHashSet[int]()

	// 插入重复元素
	s.Insert(1)
	s.Insert(2)
	s.Insert(2) // 重复
	s.Insert(3)
	s.Insert(1) // 重复

	// 验证大小：应该只有 3 个唯一元素
	if s.Size() != 3 {
		t.Errorf("Expected size 3 (no duplicates), got %d", s.Size())
	}

	// 验证内容
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Error("Missing expected elements")
	}
}

// TestLinkedHashSet_Ordering 测试保持插入顺序的特性
func TestLinkedHashSet_Ordering(t *testing.T) {
	s := ds.NewLinkedHashSet[string]()

	// 按特定顺序插入
	items := []string{"first", "second", "third", "fourth"}
	for _, item := range items {
		s.Insert(item)
	}

	// 插入一个已存在的元素，顺序不应改变
	s.Insert("second")

	values := s.Values() // 假设你也会为 Set 实现 Values() 方法，或者用 Keys() 的类比

	// 注意：如果你的 LinkedHashSet 还没有暴露 Values() 方法，
	// 你需要在 ds 包中为 LinkedHashSet 添加一个 Values() []T 方法，
	// 其内部调用 lhs.mapImpl.Keys()。
	// 下面假设你已经添加了该方法，如果没有，请看下方的【重要提示】。

	if len(values) != len(items) {
		t.Fatalf("Expected %d values, got %d", len(items), len(values))
	}

	for i, v := range values {
		if v != items[i] {
			t.Errorf("At index %d: expected %s, got %s", i, items[i], v)
		}
	}
}

// TestLinkedHashSet_Remove 测试删除功能
func TestLinkedHashSet_Remove(t *testing.T) {
	s := ds.NewLinkedHashSet[int]()

	s.Insert(10)
	s.Insert(20)
	s.Insert(30)
	s.Insert(40)

	// 1. 删除中间元素
	if !s.Remove(20) {
		t.Error("Expected Remove(20) to return true")
	}
	if s.Contains(20) {
		t.Error("20 should be removed")
	}
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	// 2. 删除不存在的元素
	if s.Remove(99) {
		t.Error("Expected Remove(99) to return false")
	}

	// 3. 删除头和尾
	s.Remove(10) // 头
	s.Remove(40) // 尾

	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
	if !s.Contains(30) {
		t.Error("Expected 30 to remain")
	}
}

// TestLinkedHashSet_Clear 测试清空功能
func TestLinkedHashSet_Clear(t *testing.T) {
	s := ds.NewLinkedHashSet[string]()

	s.Insert("a")
	s.Insert("b")
	s.Insert("c")

	if s.Size() != 3 {
		t.Fatal("Setup failed")
	}

	s.Clear()

	if s.Size() != 0 {
		t.Errorf("Expected size 0 after Clear, got %d", s.Size())
	}
	if s.Contains("a") {
		t.Error("Set should be empty after Clear")
	}

	// 验证清空后可重用
	s.Insert("new")
	if s.Size() != 1 || !s.Contains("new") {
		t.Error("Failed to insert after Clear")
	}
}

// TestLinkedHashSet_DifferentTypes 测试泛型支持
func TestLinkedHashSet_DifferentTypes(t *testing.T) {
	type Point struct {
		X int
		Y int
	}

	s := ds.NewLinkedHashSet[Point]()
	p1 := Point{1, 2}
	p2 := Point{3, 4}

	s.Insert(p1)
	s.Insert(p2)
	s.Insert(p1) // 重复

	if s.Size() != 2 {
		t.Errorf("Expected size 2 for struct set, got %d", s.Size())
	}
	if !s.Contains(p1) {
		t.Error("Should contain p1")
	}
}

// TestLinkedHashSet_EmptyOperations 测试空集合操作
func TestLinkedHashSet_EmptyOperations(t *testing.T) {
	s := ds.NewLinkedHashSet[float64]()

	if s.Contains(1.1) {
		t.Error("Empty set should not contain anything")
	}
	if s.Remove(1.1) {
		t.Error("Remove on empty set should return false")
	}

	// 安全测试：空集合 Clear 不应 panic
	s.Clear()
	if s.Size() != 0 {
		t.Error("Size should remain 0")
	}
}
