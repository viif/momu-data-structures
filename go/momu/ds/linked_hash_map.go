// 包 ds 提供了各种数据结构的实现。
package ds

import (
	"container/list"
)

// LinkedHashMap 哈希链表，一个保持插入顺序的哈希映射。它使用一个双向链表来维护元素的顺序，同时使用一个哈希表来实现快速访问。
type LinkedHashMap[K comparable, V any] struct {
	nodes    *list.List
	key2node map[K]*list.Element // Element 内部存储的是 *node[K, V]
}

type node[K comparable, V any] struct {
	key   K
	value V
}

func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		nodes:    list.New(),
		key2node: make(map[K]*list.Element),
	}
}

func (lhm *LinkedHashMap[K, V]) Get(key K) (V, bool) {
	if elem, ok := lhm.key2node[key]; ok {
		return elem.Value.(*node[K, V]).value, true
	}
	var zero V
	return zero, false
}

func (lhm *LinkedHashMap[K, V]) Put(key K, value V) {
	if elem, ok := lhm.key2node[key]; ok {
		elem.Value.(*node[K, V]).value = value
	} else {
		newNode := &node[K, V]{
			key:   key,
			value: value,
		}
		element := lhm.nodes.PushBack(newNode)
		lhm.key2node[key] = element
	}
}

func (lhm *LinkedHashMap[K, V]) Remove(key K) bool {
	if elem, ok := lhm.key2node[key]; ok {
		lhm.nodes.Remove(elem)
		delete(lhm.key2node, key)
		return true
	}
	return false
}

func (lhm *LinkedHashMap[K, V]) Clear() {
	lhm.nodes.Init()
	lhm.key2node = make(map[K]*list.Element)
}

func (lhm *LinkedHashMap[K, V]) Contains(key K) bool {
	_, ok := lhm.key2node[key]
	return ok
}

func (lhm *LinkedHashMap[K, V]) Size() int {
	return lhm.nodes.Len()
}

func (lhm *LinkedHashMap[K, V]) Keys() []K {
	keys := make([]K, 0, lhm.nodes.Len())
	for elem := lhm.nodes.Front(); elem != nil; elem = elem.Next() {
		keys = append(keys, elem.Value.(*node[K, V]).key)
	}
	return keys
}
