package lru

type Cache struct {
	capacity   int64
	size       int64
	hash       map[string]*Node
	head, tail *Node
	OnDelete   func(key string, value interface{})
}

type Node struct {
	key       string
	value     interface{}
	pre, next *Node
}

func New(capacity int64, onDelete func(key string, value interface{})) *Cache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.pre = head
	return &Cache{
		capacity: capacity,
		size:     0,
		hash:     map[string]*Node{},
		head:     head,
		tail:     tail,
		OnDelete: onDelete,
	}
}

func (c *Cache) Get(key string) (val interface{}, ok bool) {
	if n, ok := c.hash[key]; ok {
		c.moveHead(n)
		return n.value, true
	}
	return nil, false
}
func (c *Cache) Put(key string, value interface{}) {
	if n, ok := c.hash[key]; ok {
		n.value = value
		c.moveHead(n)
	} else {
		n := &Node{key, value, nil, nil}
		c.hash[key] = n
		c.addHead(n)
		c.size++
	}
	for c.size > c.capacity {
		c.removeTail()
		c.size--
	}
}
func (c *Cache) Len() int64 {
	return c.size
}

func (c *Cache) moveHead(n *Node) {
	c.removeNode(n)
	c.addHead(n)
}
func (c *Cache) addHead(n *Node) {
	c.head.next.pre = n
	n.next = c.head.next
	n.pre = c.head
	c.head.next = n
}
func (c *Cache) removeNode(n *Node) {
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (c *Cache) removeTail() {
	t := c.tail.pre
	t.pre.next = t.next
	t.next.pre = t.pre
	delete(c.hash, t.key)
	c.OnDelete(t.key, t.value)
}
