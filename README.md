# momu-data-structures

一些常用数据结构。

| 类名                       | 说明                                                         | C++ 实现 | Go 实现 |
| -------------------------- | ------------------------------------------------------------ | :------: | :-----: |
| Array                   | 定长数组                                                     |    ✅    |       |
| ArrayList               | 变长数组                                                     |    ✅    |       |
| LinkedList              | 双向链表                                                     |    ✅    |       |
| ArrayStack              | 栈，基于变长数组                                             |    ✅    |       |
| LinkedStack             | 栈，基于双向链表                                             |    ✅    |       |
| ArrayDeque              | 双端队列，基于环形数组                                       |    ✅    |       |
| ArrayQueue              | 队列，基于双端队列                                           |    ✅    |       |
| LinkedQueue             | 队列，基于双向链表                                           |    ✅    |       |
| RingBuffer              | 环形缓冲区                                                   |    ✅    |       |
| SeparateChainingHashMap | 哈希映射，使用拉链法解决冲突                                 |    ✅    |       |
| LinearProbingHashMap    | 哈希映射，使用线性探查法解决冲突                             |    ✅    |       |
| HashSet                 | 哈希集合，使用线性探查法解决冲突                             |    ✅    |       |
| LinkedHashMap           | 映射，基于哈希链表，特性：可以顺序性访问所有 key，返回顺序即插入顺序 |    ✅    | ✅ |
| LinkedHashSet           | 集合，基于哈希链表，特性：可以顺序性访问所有 value，返回顺序即插入顺序 |    ✅    | ✅ |
| ArrayHashMap            | 映射，基于哈希数组，特性：可以在 O(1) 时间内等概率地随机返回一个 key |    ✅    |       |
| ArrayHashSet            | 集合，基于哈希数组，特性：可以在 O(1) 时间内等概率地随机返回一个 value |    ✅    |       |
| RecursiveList           | 单向链表，各种操作以递归实现                                 |    ✅    |       |
| TreeMap                 | 映射，基于普通 BST                                           |    ✅    |       |
| TrieMap                 | 映射，基于前缀树                                             |    ✅    |       |
| TrieSet                 | 集合，基于前缀树                                             |    ✅    |       |
| LRUCache                | LRU（Least Recently Used，最近最少使用）缓存，对应 [146. LRU 缓存 - 力扣（LeetCode）](https://leetcode.cn/problems/lru-cache/) |    ✅    |       |
| LFUCache                | LFU（Least Frequently Used，最不经常使用）缓存，对应 [460. LFU 缓存 - 力扣（LeetCode）](https://leetcode.cn/problems/lfu-cache/description/) |    ✅    |       |
