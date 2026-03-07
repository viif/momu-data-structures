package ds_test

import (
	"momu/ds"
	"testing"
)

func TestLRUCache_GetMissing(t *testing.T) {
	lc := ds.NewLRUCache(2)
	val := lc.Get(1)
	if val != -1 {
		t.Errorf("Expected -1 for missing key, got %d", val)
	}
}

func TestLRUCache_PutAndGet(t *testing.T) {
	lc := ds.NewLRUCache(2)

	lc.Put(1, 10)
	lc.Put(2, 20)

	if v := lc.Get(1); v != 10 {
		t.Errorf("Expected 10 for key 1, got %d", v)
	}
	if v := lc.Get(2); v != 20 {
		t.Errorf("Expected 20 for key 2, got %d", v)
	}
}

func TestLRUCache_Eviction(t *testing.T) {
	lc := ds.NewLRUCache(2)

	lc.Put(1, 10)
	lc.Put(2, 20)

	// Cache is now full: [1, 2] (2 is most recent)

	// Add a new key, should evict key 1 (least recently used)
	lc.Put(3, 30)

	if v := lc.Get(1); v != -1 {
		t.Errorf("Expected key 1 to be evicted, got value %d", v)
	}
	if v := lc.Get(2); v != 20 {
		t.Errorf("Expected key 2 to exist with value 20, got %d", v)
	}
	if v := lc.Get(3); v != 30 {
		t.Errorf("Expected key 3 to exist with value 30, got %d", v)
	}
}

func TestLRUCache_UpdateExisting(t *testing.T) {
	lc := ds.NewLRUCache(2)

	lc.Put(1, 10)
	lc.Put(2, 20)

	// Update key 1, this makes key 1 the most recently used
	// Order becomes: [2, 1] (1 is most recent)
	lc.Put(1, 100)

	// Add key 3, should evict key 2 (now the least recently used)
	lc.Put(3, 30)

	if v := lc.Get(1); v != 100 {
		t.Errorf("Expected key 1 to be updated to 100, got %d", v)
	}
	if v := lc.Get(2); v != -1 {
		t.Errorf("Expected key 2 to be evicted, got value %d", v)
	}
	if v := lc.Get(3); v != 30 {
		t.Errorf("Expected key 3 to exist, got %d", v)
	}
}

func TestLRUCache_PromoteOnGet(t *testing.T) {
	lc := ds.NewLRUCache(2)

	lc.Put(1, 10)
	lc.Put(2, 20)

	// Access key 1, making it most recently used
	// Order becomes: [2, 1]
	lc.Get(1)

	// Add key 3, should evict key 2
	lc.Put(3, 30)

	if v := lc.Get(1); v != 10 {
		t.Errorf("Expected key 1 to exist, got %d", v)
	}
	if v := lc.Get(2); v != -1 {
		t.Errorf("Expected key 2 to be evicted because key 1 was accessed recently, got %d", v)
	}
	if v := lc.Get(3); v != 30 {
		t.Errorf("Expected key 3 to exist, got %d", v)
	}
}

func TestLRUCache_CapacityOne(t *testing.T) {
	lc := ds.NewLRUCache(1)

	lc.Put(1, 10)
	if v := lc.Get(1); v != 10 {
		t.Errorf("Expected 10, got %d", v)
	}

	// Add new key, should evict the only existing key
	lc.Put(2, 20)

	if v := lc.Get(1); v != -1 {
		t.Errorf("Expected key 1 to be evicted, got %d", v)
	}
	if v := lc.Get(2); v != 20 {
		t.Errorf("Expected key 2 to exist, got %d", v)
	}

	// Overwrite existing key
	lc.Put(2, 200)
	if v := lc.Get(2); v != 200 {
		t.Errorf("Expected key 2 to be updated to 200, got %d", v)
	}
}

func TestLRUCache_Sequence(t *testing.T) {
	// This test mimics a typical LeetCode test case sequence
	lc := ds.NewLRUCache(2)

	lc.Put(2, 1)
	lc.Put(1, 1)
	lc.Put(2, 3) // updates
	lc.Put(4, 1) // evicts key 1

	if v := lc.Get(1); v != -1 {
		t.Errorf("Expected key 1 to be evicted, got %d", v)
	}
	if v := lc.Get(2); v != 3 {
		t.Errorf("Expected key 2 to be 3, got %d", v)
	}
}
