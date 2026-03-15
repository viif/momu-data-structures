package ds

type treeNode[K comparable, V any] struct {
	key   K
	value V
	size  uint
	left  *treeNode[K, V]
	right *treeNode[K, V]
}

type TreeMap[K comparable, V any] struct {
	root *treeNode[K, V]
	less func(K, K) bool
}

func NewTreeMap[K comparable, V any](less func(K, K) bool) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		root: nil,
		less: less,
	}
}

func (tm *TreeMap[K, V]) Put(key K, value V) {
	tm.root = tm.insertNode(tm.root, key, value)
}

func (tm *TreeMap[K, V]) Get(key K) (value V, exist bool) {
	node := tm.findNode(tm.root, key)
	if node == nil {
		return
	}
	return node.value, true
}

func (tm *TreeMap[K, V]) Remove(key K) {
	tm.root = tm.removeNode(tm.root, key)
}

func (tm *TreeMap[K, V]) Contains(key K) bool {
	return tm.findNode(tm.root, key) != nil
}

func (tm *TreeMap[K, V]) Clear() {
	tm.root = nil
}

func (tm *TreeMap[K, V]) Size() uint {
	return tm.getNodeSize(tm.root)
}

func (tm *TreeMap[K, V]) Keys() (keys []K) {
	var traverse func(*treeNode[K, V])
	traverse = func(node *treeNode[K, V]) {
		if node == nil {
			return
		}
		traverse(node.left)
		keys = append(keys, node.key)
		traverse(node.right)
	}
	traverse(tm.root)
	return keys
}

func (tm *TreeMap[K, V]) MinKey() (key K, exist bool) {
	if tm.root == nil {
		return
	}
	cur := tm.root
	for cur.left != nil {
		cur = cur.left
	}
	return cur.key, true
}

func (tm *TreeMap[K, V]) MaxKey() (key K, exist bool) {
	if tm.root == nil {
		return
	}
	cur := tm.root
	for cur.right != nil {
		cur = cur.right
	}
	return cur.key, true
}

func (tm *TreeMap[K, V]) FloorKey(key K) (floorKey K, exist bool) {
	if tm.root == nil {
		return
	}
	cur := tm.root
	var candidate *treeNode[K, V]
	for cur != nil {
		if tm.less(key, cur.key) {
			// key < cur.key: 当前节点太大，去左子树找更小的
			cur = cur.left
		} else if tm.less(cur.key, key) {
			// cur.key < key: 当前节点是一个候选者（比 key 小）
			// 记录下来，并尝试去右子树找更大的（但仍小于 key）
			candidate = cur
			cur = cur.right
		} else {
			// cur.key == key: 完美匹配，直接返回
			return cur.key, true
		}
	}
	// 循环结束未找到相等值，检查是否有候选者
	if candidate == nil {
		return
	}
	return candidate.key, true
}

func (tm *TreeMap[K, V]) CeilingKey(key K) (ceilingKey K, exist bool) {
	if tm.root == nil {
		return
	}
	cur := tm.root
	var candidate *treeNode[K, V]
	for cur != nil {
		if tm.less(cur.key, key) {
			// cur.key < key: 当前节点太小，去右子树找更大的
			cur = cur.right
		} else if tm.less(key, cur.key) {
			// key < cur.key: 当前节点是一个候选者（比 key 大）
			// 记录下来，并尝试去左子树找更小的（但仍大于 key）
			candidate = cur
			cur = cur.left
		} else {
			// cur.key == key: 完美匹配，直接返回
			return cur.key, true
		}
	}
	// 循环结束未找到相等值，检查是否有候选者
	if candidate == nil {
		return
	}
	return candidate.key, true
}

func (tm *TreeMap[K, V]) SelectKey(k uint) (key K, exist bool) {
	if tm.root == nil {
		return
	}
	node := tm.selectNode(tm.root, k)
	if node == nil {
		return
	}
	return node.key, true
}

func (tm *TreeMap[K, V]) Rank(key K) uint {
	return tm.rankNode(tm.root, key)
}

func (tm *TreeMap[K, V]) RangeKeys(min, max K) (keys []K) {
	var traverse func(*treeNode[K, V], K, K)
	traverse = func(node *treeNode[K, V], min, max K) {
		if node == nil {
			return
		}
		if tm.less(min, node.key) {
			// min < node.key，才有必要去左子树找
			traverse(node.left, min, max)
		}
		if !tm.less(node.key, min) && !tm.less(max, node.key) {
			// min <= node.key <= max
			keys = append(keys, node.key)
		}
		if tm.less(node.key, max) {
			// max > node.key，才有必要去右子树找
			traverse(node.right, min, max)
		}
	}
	traverse(tm.root, min, max)
	return keys
}

func (tm *TreeMap[K, V]) insertNode(node *treeNode[K, V], key K, value V) *treeNode[K, V] {
	if node == nil {
		return &treeNode[K, V]{
			key:   key,
			value: value,
			size:  1,
		}
	}
	if tm.less(key, node.key) {
		// key < node.key
		node.left = tm.insertNode(node.left, key, value)
	} else if tm.less(node.key, key) {
		// key > node.key
		node.right = tm.insertNode(node.right, key, value)
	} else {
		// key == node.key
		node.value = value
	}
	node.size = 1 + tm.getNodeSize(node.left) + tm.getNodeSize(node.right)
	return node
}

func (tm *TreeMap[K, V]) getNodeSize(node *treeNode[K, V]) uint {
	if node == nil {
		return 0
	}
	return node.size
}

func (tm *TreeMap[K, V]) findNode(node *treeNode[K, V], key K) *treeNode[K, V] {
	if node == nil {
		return nil
	}
	if tm.less(key, node.key) {
		// key < node.key
		return tm.findNode(node.left, key)
	} else if tm.less(node.key, key) {
		// key > node.key
		return tm.findNode(node.right, key)
	} else {
		// key == node.key
		return node
	}
}

func (tm *TreeMap[K, V]) removeNode(node *treeNode[K, V], key K) *treeNode[K, V] {
	if node == nil {
		return nil
	}
	if tm.less(key, node.key) {
		// key < node.key
		node.left = tm.removeNode(node.left, key)
	} else if tm.less(node.key, key) {
		// key > node.key
		node.right = tm.removeNode(node.right, key)
	} else {
		// key == node.key
		if node.left == nil {
			// node 无左子树
			return node.right
		}
		if node.right == nil {
			// node 无右子树
			return node.left
		}
		// node 有左右子树
		// 不直接交换节点中的数据，而是交换节点，实现解耦
		// 复制左子树的最大节点作为新的根节点
		leftMax := tm.findMax(node.left)
		leftMax.left = tm.removeMax(node.left)
		leftMax.right = node.right
		node = leftMax
	}
	node.size = 1 + tm.getNodeSize(node.left) + tm.getNodeSize(node.right)
	return node
}

func (tm *TreeMap[K, V]) findMax(node *treeNode[K, V]) *treeNode[K, V] {
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

func (tm *TreeMap[K, V]) removeMax(node *treeNode[K, V]) *treeNode[K, V] {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node.left
	}
	node.right = tm.removeMax(node.right)
	node.size--
	return node
}

func (tm *TreeMap[K, V]) selectNode(node *treeNode[K, V], k uint) *treeNode[K, V] {
	if node == nil {
		return nil
	}
	n := tm.getNodeSize(node.left)
	if k < n {
		return tm.selectNode(node.left, k)
	} else if k > n {
		return tm.selectNode(node.right, k-n-1)
	} else {
		return node
	}
}

func (tm *TreeMap[K, V]) rankNode(node *treeNode[K, V], key K) uint {
	if node == nil {
		return 0
	}
	if tm.less(key, node.key) {
		// key < node.key，node及其右子树中的键都大于key，故只去左子树中找
		return tm.rankNode(node.left, key)
	} else if tm.less(node.key, key) {
		// key > node.key，node及其左子树中的键都小于key，还要去右子树中找
		return tm.getNodeSize(node.left) + 1 + tm.rankNode(node.right, key)
	} else {
		// key == node.key，只有左子树中的键小于key
		return tm.getNodeSize(node.left)
	}
}
