/*A linked list is a linear data structure.
 Unlike arrays, linked list elements are not stored at a contiguous location;
the elements are linked using pointers.*/

/*why using linked list instead of arrays?
--arrays need prealocation and should be assigned a fixed length before declaration
linked lists are dynamic structures , i.e. we don't know the size of data or avoid allocating larger array size
*/

/*
-The size of a linked list is not fixed.
-The last node of the linked list points to NULL.
-Memory is not wasted but extra memory is consumed as it also uses pointers to keep track of the next successive node.
-The entry point of a linked list is known as the head
*/

/*
	illustration

we have nodes that are pointing together, each node contains some data and a pointer to the next node

		 HEAD										TAIL
	 ------ ------		  ------  ------	   ------ ------
	| DATA | NEXT |	===> | DATA | NEXT | ===> | DATA | NEXT | ===> NULL
	 ------ ------		  ------ ------		   ------ ------

	 The first node is called the head. If the linked list is empty, then the value of the head is NULL.
*/
package ds

import (
	"errors"
	"fmt"
)

type Node struct {
	data          interface{}
	next_node_ptr *Node // pointer type is the Node structure itself
}

type List struct {
	head *Node // points to the first node
	len  int   // stores the list's length
}

func NewList() *List {
	list := new(List)

	return list
}

func (l *List) Insert(data interface{}) {
	/*sudo code
	-create a new node
	-put data in its data field
	-find the last node
	-make the last node point to the new node(set next_node_ptr of the last node to &newnode)
	-
	*/
	n := Node{} //why not new(Node)? beacuse we dont want reference we want a new  copied object
	// this new node has 2 fields with no value(0, nil)

	n.data = data //put the given data into data filed of our new Node

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	lastNode := l.head
	for i := 0; i < l.len; i++ {

		if lastNode.next_node_ptr == nil {
			lastNode.next_node_ptr = &n
			l.len++
			return
		}
		lastNode = lastNode.next_node_ptr
	}
}

func (l *List) GetNode(position int) *Node {

	if position < 0 || position > l.len {
		fmt.Println("Invalid position")
	}

	node := l.head //start from the head

	for i := 0; i < position; i++ {
		node = node.next_node_ptr
	}

	return node
}

func (l *List) GetData(position int) interface{} {

	node := l.GetNode(position)

	return node.data

}

func (l *List) InsertAtPosition(position int, data interface{}) {

	if position < 0 {
		fmt.Println("Invalid position")
		return
	}

	if position > l.len {
		fmt.Println("Position out of bounds")
		return
	}

	n := Node{}
	n.data = data

	if position == 0 {
		n.next_node_ptr = l.head
		l.head = &n
		l.len++
		return
	}

	targetNode := l.GetNode(position) // current node found
	n.next_node_ptr = targetNode

	prevNode := l.GetNode(position - 1) // pervious node found
	prevNode.next_node_ptr = &n
	l.len++

	/* sudo code
	-find the current node present at given the position (target node at the targeted position)
	-put it's address into the next_node_pointer of the new node

	-find prev node
	-update its next_node_ptr to point to the new node's address

	---if position was == 0
	-put head into the next_node_ptr field of new node
	-update list's head to point to the new node's address
	*/

}

func (l *List) DeleteAtPosition(position int) error {
	/*sudo code
	-find the requested node
	-update its data field to nil
	-find the prev node and update its next_node_prt value to next_node_prt value of requested node
	-decrement list len

	---if position was == 0
	-update head to point to next_node_prt value of requested node
	-decrement

	*/
	if position < 0 || position > (l.len-1) {
		fmt.Println("Invalid position")
		return errors.New("invalid position")
	}

	if l.len == 0 {
		fmt.Println("Empty LinkedList")
		return errors.New("empty linkedList, delete failed")
	}

	targetNode := l.GetNode(position)
	targetNode.data = nil

	if position == 0 {

		l.head = targetNode.next_node_ptr
		l.len--
		return nil
	}

	prevNode := l.GetNode(position - 1)
	prevNode.next_node_ptr = targetNode.next_node_ptr
	l.len--

	return nil
}

func (l *List) DeleteByValue(data interface{}) error {

	/*sudo code
	-Find the data and its respective node
	-find its previous node and update its next_node_ptr to point to next_node_ptr of given node
	-
	*/

	if l.len == 0 {
		fmt.Println("Empty LinkedList")
		return errors.New("empty linkedList, delete failed")
	}

	node := l.head

	for i := 0; i < l.len; i++ {
		if node.data == data {
			if i > 0 {
				targetNode := l.GetNode(i)
				prevNode := l.GetNode(i - 1)
				prevNode.next_node_ptr = targetNode.next_node_ptr
			} else {
				l.head = node.next_node_ptr
			}
			l.len--
			return nil

		}

		node = node.next_node_ptr
	}

	fmt.Println("Node not found")
	return errors.New("node not found")
}

func (l *List) Print() {
	if l.len == 0 {
		fmt.Println("Empty List")
	}

	node := l.head

	for i := 0; i < l.len; i++ {
		if node.next_node_ptr == nil {
			fmt.Println(node.data)
		} else {
			fmt.Println(node.data, "->")
		}
		node = node.next_node_ptr
	}
	fmt.Println()
}

func (l *List) Len() int {
	return l.len
}
