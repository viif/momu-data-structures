package ds_test

import (
	"momu/ds"
	"testing"
)

// helperLess 用于 int 类型的比较函数
func helperLess(a, b int) bool {
	return a < b
}

// helperLessStr 用于 string 类型的比较函数
func helperLessStr(a, b string) bool {
	return a < b
}

func TestTreeMap_BasicPutGet(t *testing.T) {
	tm := ds.NewTreeMap[int, string](helperLess)

	// 空树测试
	if tm.Size() != 0 {
		t.Errorf("Expected size 0, got %d", tm.Size())
	}
	if _, ok := tm.Get(1); ok {
		t.Error("Expected key 1 not to exist in empty map")
	}

	// 插入测试
	tm.Put(5, "Five")
	tm.Put(3, "Three")
	tm.Put(7, "Seven")
	tm.Put(3, "Three-Updated") // 更新已有 key

	if tm.Size() != 3 {
		t.Errorf("Expected size 3, got %d", tm.Size())
	}

	// 获取测试
	val, ok := tm.Get(3)
	if !ok || val != "Three-Updated" {
		t.Errorf("Expected 'Three-Updated', got '%v', ok=%v", val, ok)
	}

	val, ok = tm.Get(5)
	if !ok || val != "Five" {
		t.Errorf("Expected 'Five', got '%v', ok=%v", val, ok)
	}

	_, ok = tm.Get(10)
	if ok {
		t.Error("Expected key 10 not to exist")
	}
}

func TestTreeMap_ContainsAndClear(t *testing.T) {
	tm := ds.NewTreeMap[int, int](helperLess)
	tm.Put(1, 1)
	tm.Put(2, 2)

	if !tm.Contains(1) {
		t.Error("Expected Contains(1) to be true")
	}
	if tm.Contains(3) {
		t.Error("Expected Contains(3) to be false")
	}

	tm.Clear()
	if tm.Size() != 0 {
		t.Error("Expected size 0 after Clear")
	}
	if tm.Contains(1) {
		t.Error("Expected Contains(1) to be false after Clear")
	}
}

func TestTreeMap_MinMaxKey(t *testing.T) {
	tm := ds.NewTreeMap[int, string](helperLess)

	// 空树
	if _, ok := tm.MinKey(); ok {
		t.Error("MinKey on empty tree should return false")
	}
	if _, ok := tm.MaxKey(); ok {
		t.Error("MaxKey on empty tree should return false")
	}

	// 单节点
	tm.Put(10, "Ten")
	minK, ok := tm.MinKey()
	if !ok || minK != 10 {
		t.Errorf("MinKey failed: %v, %v", minK, ok)
	}
	maxK, ok := tm.MaxKey()
	if !ok || maxK != 10 {
		t.Errorf("MaxKey failed: %v, %v", maxK, ok)
	}

	// 多节点
	tm.Put(5, "Five")
	tm.Put(15, "Fifteen")
	tm.Put(2, "Two")
	tm.Put(20, "Twenty")

	minK, _ = tm.MinKey()
	if minK != 2 {
		t.Errorf("Expected MinKey 2, got %d", minK)
	}

	maxK, _ = tm.MaxKey()
	if maxK != 20 {
		t.Errorf("Expected MaxKey 20, got %d", maxK)
	}
}

func TestTreeMap_FloorCeilingKey(t *testing.T) {
	tm := ds.NewTreeMap[int, string](helperLess)
	// 插入: 10, 20, 30, 40, 50
	for i := 1; i <= 5; i++ {
		tm.Put(i*10, "")
	}

	tests := []struct {
		key        int
		wantFloor  int
		floorExist bool
		wantCeil   int
		ceilExist  bool
	}{
		{key: 25, wantFloor: 20, floorExist: true, wantCeil: 30, ceilExist: true},
		{key: 20, wantFloor: 20, floorExist: true, wantCeil: 20, ceilExist: true}, // Exact match
		{key: 5, wantFloor: 0, floorExist: false, wantCeil: 10, ceilExist: true},  // Smaller than min
		{key: 55, wantFloor: 50, floorExist: true, wantCeil: 0, ceilExist: false}, // Larger than max
		{key: 35, wantFloor: 30, floorExist: true, wantCeil: 40, ceilExist: true},
	}

	for _, tt := range tests {
		fk, fOk := tm.FloorKey(tt.key)
		if fOk != tt.floorExist || (fOk && fk != tt.wantFloor) {
			t.Errorf("FloorKey(%d): got (%v, %v), want (%v, %v)", tt.key, fk, fOk, tt.wantFloor, tt.floorExist)
		}

		ck, cOk := tm.CeilingKey(tt.key)
		if cOk != tt.ceilExist || (cOk && ck != tt.wantCeil) {
			t.Errorf("CeilingKey(%d): got (%v, %v), want (%v, %v)", tt.key, ck, cOk, tt.wantCeil, tt.ceilExist)
		}
	}
}

func TestTreeMap_RankAndSelect(t *testing.T) {
	tm := ds.NewTreeMap[int, string](helperLess)
	// 插入乱序: 50, 30, 70, 20, 40, 60, 80
	keys := []int{50, 30, 70, 20, 40, 60, 80}
	for _, k := range keys {
		tm.Put(k, "")
	}

	// Rank 测试 (0-based index of sorted keys: 20, 30, 40, 50, 60, 70, 80)
	// 20->0, 30->1, 40->2, 50->3, 60->4, 70->5, 80->6
	rankTests := []struct {
		key  int
		want uint
	}{
		{20, 0}, {30, 1}, {40, 2}, {50, 3}, {60, 4}, {70, 5}, {80, 6},
		{10, 0}, // Not exists, smaller than all
		{90, 7}, // Not exists, larger than all (rank is count of smaller keys)
	}

	for _, rt := range rankTests {
		if r := tm.Rank(rt.key); r != rt.want {
			t.Errorf("Rank(%d): got %d, want %d", rt.key, r, rt.want)
		}
	}

	// Select 测试 (0-based)
	selectTests := []struct {
		k    uint
		want int
		ok   bool
	}{
		{0, 20, true}, {1, 30, true}, {2, 40, true}, {3, 50, true},
		{6, 80, true},
		{7, 0, false}, // Out of bounds
	}

	for _, st := range selectTests {
		k, ok := tm.SelectKey(st.k)
		if ok != st.ok || (ok && k != st.want) {
			t.Errorf("SelectKey(%d): got (%v, %v), want (%v, %v)", st.k, k, ok, st.want, st.ok)
		}
	}
}

func TestTreeMap_RangeKeys(t *testing.T) {
	tm := ds.NewTreeMap[int, string](helperLess)
	for i := 1; i <= 10; i++ {
		tm.Put(i, "")
	}

	// Range [3, 7] -> 3, 4, 5, 6, 7
	keys := tm.RangeKeys(3, 7)
	expected := []int{3, 4, 5, 6, 7}
	if len(keys) != len(expected) {
		t.Fatalf("RangeKeys length mismatch: got %d, want %d", len(keys), len(expected))
	}
	for i, k := range keys {
		if k != expected[i] {
			t.Errorf("RangeKeys[%d]: got %d, want %d", i, k, expected[i])
		}
	}

	// Range [1, 1] -> 1
	keys = tm.RangeKeys(1, 1)
	if len(keys) != 1 || keys[0] != 1 {
		t.Errorf("RangeKeys(1, 1) failed: %v", keys)
	}

	// Range [11, 20] -> empty
	keys = tm.RangeKeys(11, 20)
	if len(keys) != 0 {
		t.Errorf("RangeKeys(11, 20) should be empty, got %v", keys)
	}
}

func TestTreeMap_Remove(t *testing.T) {
	tm := ds.NewTreeMap[int, string](helperLess)
	// 构建一个稍微复杂的树
	//       50
	//     /    \
	//   30      70
	//  /  \    /  \
	// 20  40  60  80
	keys := []int{50, 30, 70, 20, 40, 60, 80}
	for _, k := range keys {
		tm.Put(k, "")
	}

	// 1. 删除叶子节点 (20)
	tm.Remove(20)
	if tm.Contains(20) {
		t.Error("Failed to remove leaf node 20")
	}
	if tm.Size() != 6 {
		t.Errorf("Size after removing leaf: got %d, want 6", tm.Size())
	}
	// 验证 BST 性质未破坏 (Min/Max/Rank)
	if min, _ := tm.MinKey(); min != 30 {
		t.Errorf("MinKey after removing 20: got %d, want 30", min)
	}

	// 2. 删除只有一个子节点的节点 (30, 剩下 40)
	tm.Remove(30)
	if tm.Contains(30) {
		t.Error("Failed to remove node 30")
	}
	if !tm.Contains(40) {
		t.Error("Child 40 should still exist after removing 30")
	}
	if tm.Size() != 5 {
		t.Errorf("Size after removing node with one child: got %d, want 5", tm.Size())
	}

	// 3. 删除有两个子节点的节点 (50 - 根节点)
	// 此时树结构大致为: 40 (或 60/70 取决于实现) 成为根
	tm.Remove(50)
	if tm.Contains(50) {
		t.Error("Failed to remove root node 50")
	}
	if tm.Size() != 4 {
		t.Errorf("Size after removing root: got %d, want 4", tm.Size())
	}

	// 验证剩余元素依然有序且完整
	remainingKeys := tm.Keys()
	expectedRemaining := []int{40, 60, 70, 80} // 假设删除逻辑正确，剩下的应该是这些
	if len(remainingKeys) != len(expectedRemaining) {
		t.Fatalf("Remaining keys count mismatch: got %d, want %d", len(remainingKeys), len(expectedRemaining))
	}

	// 简单验证顺序
	for i := 0; i < len(remainingKeys)-1; i++ {
		if remainingKeys[i] >= remainingKeys[i+1] {
			t.Errorf("Keys not sorted: %v >= %v", remainingKeys[i], remainingKeys[i+1])
		}
	}

	// 4. 删除不存在的键
	tm.Remove(999)
	if tm.Size() != 4 {
		t.Error("Size changed after removing non-existent key")
	}

	// 5. 删空树
	tm.Remove(40)
	tm.Remove(60)
	tm.Remove(70)
	tm.Remove(80)
	if tm.Size() != 0 {
		t.Error("Tree should be empty")
	}
}

func TestTreeMap_KeysOrder(t *testing.T) {
	tm := ds.NewTreeMap[string, int](helperLessStr)
	tm.Put("c", 3)
	tm.Put("a", 1)
	tm.Put("b", 2)
	tm.Put("z", 26)
	tm.Put("m", 13)

	keys := tm.Keys()
	expected := []string{"a", "b", "c", "m", "z"}

	if len(keys) != len(expected) {
		t.Fatalf("Keys length mismatch")
	}
	for i, k := range keys {
		if k != expected[i] {
			t.Errorf("Index %d: got %s, want %s", i, k, expected[i])
		}
	}
}

func TestTreeMap_StringComparison(t *testing.T) {
	tm := ds.NewTreeMap[string, bool](helperLessStr)

	// 插入数据: apple, banana, cherry (按字典序排序)
	tm.Put("banana", true)
	tm.Put("apple", true)
	tm.Put("cherry", true)

	// 1. 测试 MinKey
	if min, ok := tm.MinKey(); !ok || min != "apple" {
		t.Errorf("MinKey failed: got %s (ok=%v), want apple", min, ok)
	}

	// 2. 测试 MaxKey
	if max, ok := tm.MaxKey(); !ok || max != "cherry" {
		t.Errorf("MaxKey failed: got %s (ok=%v), want cherry", max, ok)
	}

	// 3. 测试 FloorKey
	// 当前 Keys: [apple, banana, cherry]
	// 查询 "blueberry":
	// "apple" < "blueberry" -> candidate = "apple", 往右找
	// "banana" < "blueberry" -> candidate = "banana", 往右找
	// "cherry" > "blueberry" -> 往左找 (nil)
	// 结果应为 "banana"
	floor, ok := tm.FloorKey("blueberry")
	if !ok || floor != "banana" {
		t.Errorf("FloorKey('blueberry') failed: got %s (ok=%v), want banana", floor, ok)
	}

	// 4. 测试 CeilingKey (顺便验证对称逻辑)
	// 查询 "blueberry":
	// "apple" < "blueberry" -> 往右找
	// "banana" < "blueberry" -> 往右找
	// "cherry" > "blueberry" -> candidate = "cherry", 往左找 (nil)
	// 结果应为 "cherry"
	ceiling, ok := tm.CeilingKey("blueberry")
	if !ok || ceiling != "cherry" {
		t.Errorf("CeilingKey('blueberry') failed: got %s (ok=%v), want cherry", ceiling, ok)
	}

	// 5. 测试边界情况：小于所有键
	floorSmall, okSmall := tm.FloorKey("aaa")
	if okSmall {
		t.Errorf("FloorKey('aaa') should not exist, got %s", floorSmall)
	}

	// 6. 测试边界情况：大于所有键
	ceilingLarge, okLarge := tm.CeilingKey("zzz")
	if okLarge {
		t.Errorf("CeilingKey('zzz') should not exist, got %s", ceilingLarge)
	}
}
