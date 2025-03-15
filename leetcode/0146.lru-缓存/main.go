/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 *
 * https://leetcode.cn/problems/lru-cache/description/
 *
 * algorithms
 * Medium (54.29%)
 * Likes:    3421
 * Dislikes: 0
 * Total Accepted:    791.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
 *
 * 实现 LRUCache 类：
 *
 *
 *
 *
 * LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
 * key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
 *
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * 输出
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * 解释
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // 缓存是 {1=1}
 * lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
 * lRUCache.get(1);    // 返回 1
 * lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
 * lRUCache.get(2);    // 返回 -1 (未找到)
 * lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
 * lRUCache.get(1);    // 返回 -1 (未找到)
 * lRUCache.get(3);    // 返回 3
 * lRUCache.get(4);    // 返回 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 10000
 * 0 <= value <= 10^5
 * 最多调用 2 * 10^5 次 get 和 put
 *
 *
*/

package main

import (
	"fmt"
)

// 运用你所掌握的数据结构，设计和实现一个 LRU 缓存机制 。
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
// 当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

// 要求：在 O(1) 时间复杂度内完成这两种操作

// 解题思路： 利用map和双链表
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	fmt.Printf("cache1: %#+v\n", cache)
	cache.Put(2, 2)
	fmt.Printf("cache2: %#+v\n", cache)
	cache.Get(1)
	fmt.Printf("cache3: %#+v\n", cache)
	cache.Put(3, 3)
	fmt.Printf("cache4: %#+v\n", cache)
}

// @lc code=start
type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func newNode(key, value int) *node {
	return &node{
		key:   key,
		value: value,
	}
}

type doubleList struct {
	head *node
	tail *node
	len  int
}

func newDoubleList() *doubleList {
	return &doubleList{}
}

func (l *doubleList) add(n *node) {
	if l.tail == nil {
		l.head = n
		l.tail = n
		l.len++
		return
	}
	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.len++
}

func (l *doubleList) deleteFirst() {
	if l.head == nil {
		return
	}
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
		l.len--
		return
	}

	l.head = l.head.next
	l.head.prev = nil
	l.len--
}

func (l *doubleList) delete(n *node) {
	if n.prev == nil {
		l.deleteFirst()
		return
	}
	if n.next == nil {
		l.tail = n.prev
	} else {
		n.next.prev = n.prev
	}

	n.prev.next = n.next
	n.prev = nil
	n.next = nil
	l.len--
}

type LRUCache struct {
	capacity int
	m        map[int]*node
	l        *doubleList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*node, capacity),
		l:        newDoubleList(),
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.m[key]
	if !ok {
		return -1
	}

	this.l.delete(n)
	this.l.add(n)
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	new := newNode(key, value)
	if old, ok := this.m[key]; ok {
		this.l.delete(old)
		this.l.add(new)
		this.m[key] = new
		return
	}
	if this.capacity == len(this.m) {
		delete(this.m, this.l.head.key)
		this.l.deleteFirst()
	}
	this.m[key] = new
	this.l.add(new)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// @lc code=end
