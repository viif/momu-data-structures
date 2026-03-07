package ds

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	size     int
	key2node map[int]*list.Element
	nodes    *list.List
}

type intNode struct {
	key   int
	value int
}

func NewLRUCache(capacity int) LRUCache {
	lruCache := &LRUCache{
		capacity: capacity,
		key2node: make(map[int]*list.Element),
		nodes:    list.New(),
	}
	return *lruCache
}

func (lc *LRUCache) Get(key int) int {
	elem, exist := lc.key2node[key]
	if !exist {
		return -1
	}
	lc.promoteNodeToRecent(key)
	return elem.Value.(*intNode).value
}

func (lc *LRUCache) Put(key int, value int) {
	if elem, exist := lc.key2node[key]; exist {
		elem.Value.(*intNode).value = value
		lc.promoteNodeToRecent(key)
	} else {
		if lc.size >= lc.capacity {
			lc.evictLeastRecentlyUsed()
		}
		newNode := &intNode{
			key:   key,
			value: value,
		}
		lc.nodes.PushBack(newNode)
		lc.key2node[key] = lc.nodes.Back()
		lc.size++
	}
}

func (lc *LRUCache) promoteNodeToRecent(key int) {
	elem := lc.key2node[key]
	lc.nodes.PushBack(elem.Value.(*intNode))
	lc.nodes.Remove(elem)
	lc.key2node[key] = lc.nodes.Back()
}

func (lc *LRUCache) evictLeastRecentlyUsed() {
	key := lc.nodes.Front().Value.(*intNode).key
	lc.nodes.Remove(lc.nodes.Front())
	delete(lc.key2node, key)
	lc.size--
}
