package ds_test

import (
	"momu/ds"
	"testing"
)

// TestLinkedHashMap_BasicOperations 测试基本的 Put, Get, Contains, Size
func TestLinkedHashMap_BasicOperations(t *testing.T) {
	m := ds.NewLinkedHashMap[string, int]()

	// 1. 测试空状态
	if m.Size() != 0 {
		t.Errorf("Expected size 0, got %d", m.Size())
	}
	if m.Contains("key1") {
		t.Error("Expected map to not contain key1")
	}

	// 2. 测试插入
	m.Put("one", 1)
	m.Put("two", 2)
	m.Put("three", 3)

	if m.Size() != 3 {
		t.Errorf("Expected size 3, got %d", m.Size())
	}

	// 3. 测试获取存在的值
	val, ok := m.Get("two")
	if !ok {
		t.Error("Expected to find key 'two'")
	}
	if val != 2 {
		t.Errorf("Expected value 2, got %d", val)
	}

	// 4. 测试获取不存在的值
	_, ok = m.Get("four")
	if ok {
		t.Error("Expected not to find key 'four'")
	}

	// 5. 测试 Contains
	if !m.Contains("one") {
		t.Error("Expected Contains('one') to be true")
	}
	if m.Contains("four") {
		t.Error("Expected Contains('four') to be false")
	}
}

// TestLinkedHashMap_UpdateExistingKey 测试更新已存在的 Key
func TestLinkedHashMap_UpdateExistingKey(t *testing.T) {
	m := ds.NewLinkedHashMap[string, string]()

	m.Put("key", "original")
	val, _ := m.Get("key")
	if val != "original" {
		t.Errorf("Expected 'original', got %s", val)
	}

	// 更新值
	m.Put("key", "updated")

	// 验证值已改变
	val, ok := m.Get("key")
	if !ok {
		t.Error("Key should still exist after update")
	}
	if val != "updated" {
		t.Errorf("Expected 'updated', got %s", val)
	}

	// 验证大小没有增加
	if m.Size() != 1 {
		t.Errorf("Expected size 1 after update, got %d", m.Size())
	}
}

// TestLinkedHashMap_Ordering 测试 LinkedHashMap 的核心特性：保持插入顺序
func TestLinkedHashMap_Ordering(t *testing.T) {
	m := ds.NewLinkedHashMap[int, string]()

	// 插入顺序: 10, 20, 30, 40
	m.Put(10, "ten")
	m.Put(20, "twenty")
	m.Put(30, "thirty")
	m.Put(40, "forty")

	// 更新中间的键 (20)，顺序不应改变
	m.Put(20, "twenty-updated")

	keys := m.Keys()
	expectedKeys := []int{10, 20, 30, 40}

	if len(keys) != len(expectedKeys) {
		t.Fatalf("Expected %d keys, got %d", len(expectedKeys), len(keys))
	}

	for i, k := range keys {
		if k != expectedKeys[i] {
			t.Errorf("At index %d: expected key %d, got %d", i, expectedKeys[i], k)
		}
	}

	// 验证更新后的值也在正确的位置
	val, _ := m.Get(20)
	if val != "twenty-updated" {
		t.Errorf("Expected updated value, got %s", val)
	}
}

// TestLinkedHashMap_Remove 测试删除功能（头、中、尾）
func TestLinkedHashMap_Remove(t *testing.T) {
	m := ds.NewLinkedHashMap[string, int]()

	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	m.Put("d", 4)

	// 1. 删除中间的键
	removed := m.Remove("b")
	if !removed {
		t.Error("Expected Remove('b') to return true")
	}
	if m.Contains("b") {
		t.Error("Key 'b' should be removed")
	}
	if m.Size() != 3 {
		t.Errorf("Expected size 3, got %d", m.Size())
	}

	// 2. 验证顺序是否保持 (跳过被删除的)
	keys := m.Keys()
	expected := []string{"a", "c", "d"}
	for i, k := range keys {
		if k != expected[i] {
			t.Errorf("Order mismatch at %d: expected %s, got %s", i, expected[i], k)
		}
	}

	// 3. 删除不存在的键
	removed = m.Remove("z")
	if removed {
		t.Error("Expected Remove('z') to return false")
	}

	// 4. 删除头节点和尾节点
	m.Remove("a") // 头
	m.Remove("d") // 尾

	keys = m.Keys()
	if len(keys) != 1 || keys[0] != "c" {
		t.Errorf("Expected only ['c'] remaining, got %v", keys)
	}
}

// TestLinkedHashMap_Clear 测试清空功能 (新增)
func TestLinkedHashMap_Clear(t *testing.T) {
	m := ds.NewLinkedHashMap[string, int]()

	// 填充数据
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	if m.Size() != 3 {
		t.Fatal("Setup failed: expected size 3")
	}

	// 执行 Clear
	m.Clear()

	// 验证状态
	if m.Size() != 0 {
		t.Errorf("Expected size 0 after Clear, got %d", m.Size())
	}

	if len(m.Keys()) != 0 {
		t.Error("Expected Keys() to be empty after Clear")
	}

	if m.Contains("a") {
		t.Error("Expected map to not contain 'a' after Clear")
	}

	// 验证 Clear 后可以重新插入
	m.Put("new", 100)
	if m.Size() != 1 {
		t.Error("Expected size 1 after re-inserting post-Clear")
	}
	val, ok := m.Get("new")
	if !ok || val != 100 {
		t.Error("Failed to get re-inserted value")
	}
}

// TestLinkedHashMap_DifferentTypes 测试不同类型的 K 和 V
func TestLinkedHashMap_DifferentTypes(t *testing.T) {
	// 测试 Key 为 int, Value 为 struct
	type User struct {
		Name string
		Age  int
	}

	m := ds.NewLinkedHashMap[int, User]()
	m.Put(1, User{Name: "Alice", Age: 25})
	m.Put(2, User{Name: "Bob", Age: 30})

	user, ok := m.Get(1)
	if !ok {
		t.Fatal("Failed to get user")
	}
	if user.Name != "Alice" || user.Age != 25 {
		t.Errorf("Unexpected user data: %+v", user)
	}
}

// TestLinkedHashMap_EmptyOperations 测试空 哈希链表的操作
func TestLinkedHashMap_EmptyOperations(t *testing.T) {
	m := ds.NewLinkedHashMap[string, int]()

	_, ok := m.Get("nonexistent")
	if ok {
		t.Error("Get on empty map should return false")
	}

	if m.Remove("nonexistent") {
		t.Error("Remove on empty map should return false")
	}

	keys := m.Keys()
	if len(keys) != 0 {
		t.Error("Keys on empty map should be empty slice")
	}

	// 额外测试：在空 哈希链表上调用 Clear 不应 panic
	m.Clear()
	if m.Size() != 0 {
		t.Error("Clear on empty map should remain size 0")
	}
}
