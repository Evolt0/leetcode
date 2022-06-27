package lfu

type lfuNode struct {
	K     int
	V     int
	count int
	pre   *lfuNode
	next  *lfuNode
}
type LFUCache struct {
	cache    map[int]*lfuNode
	head     *lfuNode
	capacity int
}

func Constructor(capacity int) LFUCache {
	lfuCache := LFUCache{}
	lfuCache.capacity = capacity
	lfuCache.cache = make(map[int]*lfuNode, capacity)
	return lfuCache
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}
	node := &lfuNode{
		K:     key,
		V:     v.V,
		count: v.count + 1,
	}
	if this.head.count >= node.count {
		node.pre = nil
		node.next = this.head
		this.head.pre = node
		this.head = node
	}
	this.cache[key] = node
	return v.V
}

func (this *LFUCache) Put(key int, value int) {
	v, ok := this.cache[key]
	if ok {
		node := &lfuNode{
			K:     key,
			V:     value,
			count: v.count + 1,
		}
		this.cache[key] = node
		if this.head.count >= node.count {
			node.pre = nil
			node.next = this.head
			this.head.pre = node
			this.head = node
		}
		return
	}
	if len(this.cache) < this.capacity {
		node := &lfuNode{
			K:     key,
			V:     value,
			count: 1,
		}
		this.cache[key] = node
		if this.head == nil {
			this.head = node
		} else if this.head.count >= node.count {
			node.pre = nil
			node.next = this.head
			this.head.pre = node
			this.head = node
		}
		return
	}
	delete(this.cache, this.head.K)
	node := &lfuNode{
		K:     key,
		V:     value,
		count: 1,
	}
	this.cache[key] = node
	if this.head.count >= node.count {
		node.pre = nil
		node.next = this.head
		this.head.pre = node
		this.head = node
	}
	return
}
