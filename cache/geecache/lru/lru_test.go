package lru

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	c := New(3, func(key string, value interface{}) {
		fmt.Println("删除了", key, "值是", value)
	})
	c.Put("c", 3)
	c.Put("a", 1)
	c.Put("c", "覆盖")
	c.Put("b", 2)
	c.Put("d", 4)
	val, _ := c.Get("c")
	c.Put("f", 5)
	c.Put("e", 6)
	println(fmt.Sprintf("%v", val))

}
