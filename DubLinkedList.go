package ds

type DNode struct {
	data          interface{}
	next_node_ptr *DNode // pointer type is the Node structure itself
	prev_node_ptr *DNode
}

type DubList struct {
	head *DNode // points to the first node
	len  int    // stores the list's length
}

func NewDubList() *DubList {

	return new(DubList).initiate()

}

func (dl *DubList) initiate() *DubList {
	head := dl.head
	head.prev_node_ptr = nil

	return dl
}

func (dl *DubList) Insert(data interface{}) {

	dn := DNode{}

	if dl.len == 0 {
		dl.head = &dn
		dn.prev_node_ptr = &dl.head
		dl.len++
		return

	}

	node := dl.head

	for i := 0; i < dl.len; i++ {
		if node.next_node_ptr == nil {
			node.next_node_ptr = &dn
			dn.prev_node_ptr = node
			dl.len++
			return
		}
		node = node.next_node_ptr
	}
}
