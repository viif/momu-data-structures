// 包 ds 提供了各种数据结构的实现。
package ds

// LinkedHashSet 一个可以保持插入顺序的哈希集合。
type LinkedHashSet[T comparable] struct {
	mapImpl *LinkedHashMap[T, struct{}]
}

func NewLinkedHashSet[T comparable]() *LinkedHashSet[T] {
	return &LinkedHashSet[T]{
		mapImpl: NewLinkedHashMap[T, struct{}](),
	}
}

// Insert 将一个元素插入集合中。
func (lhs *LinkedHashSet[T]) Insert(value T) {
	lhs.mapImpl.Put(value, struct{}{})
}

// Remove 从集合中移除一个元素。
func (lhs *LinkedHashSet[T]) Remove(value T) bool {
	return lhs.mapImpl.Remove(value)
}

// Clear 清空集合中的所有元素。
func (lhs *LinkedHashSet[T]) Clear() {
	lhs.mapImpl.Clear()
}

// Contains 检查集合中是否存在一个元素。
func (lhs *LinkedHashSet[T]) Contains(value T) bool {
	return lhs.mapImpl.Contains(value)
}

// Size 返回集合中元素的数量。
func (lhs *LinkedHashSet[T]) Size() int {
	return lhs.mapImpl.Size()
}

// Values 返回集合中所有元素的切片，保持插入顺序。
func (lhs *LinkedHashSet[T]) Values() []T {
	values := lhs.mapImpl.Keys()
	return values
}
