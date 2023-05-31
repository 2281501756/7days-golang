package gee

type node struct {
	pattern  string  // 匹配路径只有在匹配成功的叶子节点才值否则都为空字符串
	part     string  // 部分值代表这个节点值
	children []*node // 孩子
	isWild   bool    // 是否精准匹配 : 和 * 开头的一定匹配成功其他都是false
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(pattern string, parts []string, height int) *node {
	if len(parts) == height || (len(n.part) > 0 && n.part[0] == '*') {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, c := range children {
		res := c.search(pattern, parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}

// 插入的匹配
func (n *node) matchChild(part string) *node {
	for _, c := range n.children {
		if c.part == part || c.isWild {
			return c
		}
	}
	return nil
}

// 搜索的匹配
func (n *node) matchChildren(part string) []*node {
	res := make([]*node, 0)
	// 先匹配精准的part放到前面
	for _, c := range n.children {
		if c.part == part {
			res = append(res, c)
		}
	}
	// 后匹配通配的part放到后面
	for _, c := range n.children {
		if c.isWild {
			res = append(res, c)
		}
	}
	return res
}
