package ds

import (
	"errors"
	"fmt"
)

type DNode struct {
	data          interface{}
	next_node_ptr *DNode // pointer type is the Node structure itself
	prev_node_ptr *DNode
}

type DoublyLinkedList struct {
	head *DNode // points to the first node
	len  int    // stores the list's length
}

func NewDubList() *DoublyLinkedList {

	return new(DoublyLinkedList).initiate()

}

func (dll *DoublyLinkedList) initiate() *DoublyLinkedList {
	head := dll.head
	head.prev_node_ptr = nil

	return dll
}

// push a new node at the end of list
func (dll *DoublyLinkedList) Insert(data interface{}) {

	dn := DNode{}

	if dll.len == 0 {
		dll.head = &dn
		dll.len++
		return

	}

	node := dll.head

	for i := 0; i < dll.len; i++ {
		if node.next_node_ptr == nil {
			node.next_node_ptr = &dn
			dn.prev_node_ptr = node
			dll.len++
			return
		}
		node = node.next_node_ptr
	}
}

func (dll *DoublyLinkedList) GetNode(position int) *DNode {

	if position < 0 || position > dll.len {
		fmt.Println("Invalid position")
	}

	node := dll.head

	for i := 0; i < position; i++ {
		node = node.next_node_ptr
	}

	return node
}

func (dll *DoublyLinkedList) GetData(position int) interface{} {

	node := dll.GetNode(position)

	return node.data

}

func (dll *DoublyLinkedList) InsertAtPosition(position int, data interface{}) {
	/*pseudo code
	- create a new node,
	 - find the node at the given position
	 - update the prev_node_ptr of (the node at the given position) to point to (the new node)
	 - set the next_node_ptr of your (new node) to point to (the node at the given position)

	 - find the prev node of (the node at the given position)
	 - update the next_node_ptr of the prev node of (the node at the given position) to point to (the new node)
	 - set the prev_node_ptr of the (new node) to point to the prev node of (the node at the given position)
	 -if position == 0
	 	--
	*/
	if position < 0 {
		fmt.Println("Invalid position")
		return
	}

	if position > dll.len {
		fmt.Println("Position out of bounds")
		return
	}

	n := DNode{
		data: data,
	}

	if position == 0 {
		head := dll.head
		n.next_node_ptr = head
		n.prev_node_ptr = nil

		head.prev_node_ptr = &n
		dll.len++

		return

	}

	givenNode := dll.GetNode(position) //(the node at the given position)

	givenNode.prev_node_ptr = &n
	n.next_node_ptr = givenNode

	givenNode_prev := dll.GetNode(position - 1) //the prev node of (the node at the given position)

	givenNode_prev.next_node_ptr = &n
	n.prev_node_ptr = givenNode_prev

	dll.len++

}

func (dll *DoublyLinkedList) DeleteAtPosition(position int) error {

	/*pseudo code
	-find the given node
	-the its prev node and next node
	-set all the values of the given node to nil
	-set the next_node_ptr of the prev node to point to the next node
	-set the prev_node_ptr of the nex node tp point to the prev node

	*/

	if position < 0 || position > (dll.len-1) {
		fmt.Println("Invalid position")
		return errors.New("invalid position")
	}

	if dll.len == 0 {
		fmt.Println("Empty LinkedList")
		return errors.New("empty linkedList, delete failed")
	}

	if position == 0 {
		head := dll.head
		dll.head = head.next_node_ptr
		dll.len--
		return nil

	}
	prevNode := dll.GetNode(position - 1)
	NextNode := dll.GetNode(position + 1)

	prevNode.next_node_ptr = NextNode
	NextNode.prev_node_ptr = prevNode
	dll.len--
	return nil
}

/*
	 func (dll *DoublyLinkedList) DeleteByValue(data interface{}) error {

		// pseudo code//

		// not implemented yet //
	}
*/
func (dn *DNode) Forward() *DNode {
	return dn.next_node_ptr
}

func (dn *DNode) Backward() *DNode {
	return dn.prev_node_ptr
}

func (dll *DoublyLinkedList) Len() int {
	return dll.len
}
