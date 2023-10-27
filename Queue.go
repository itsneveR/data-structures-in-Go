package ds

import "errors"

/*
 The difference between a queue and a stack is the removal order,
 in a stack, we remove the item the most recently added;
 in a queue, we remove the item the least recently added.

 -- Queue follows FIFO order for insertion and removal of data,
*/

type QueueArray []interface{}

type Queue struct {
	data QueueArray
	head int
	tail int
}

func NewQueue(capacity int) *Queue {
	return new(Queue).newQueue(capacity)
}

func (q *Queue) newQueue(capacity int) *Queue {

	q.data = make(QueueArray, capacity)
	q.head = -1
	q.tail = -1

	return q
}

// Add an element to the end of the queue
func (q *Queue) Enqueue(data interface{}) error {
	/*pseudo code
	  - check if queue Isfull()
	  -for the first element, set the value of 'head' to 0
	  -increase the 'tail' index by 1
	  -add the new element in the position pointed to by 'tail'
	*/

	if q.IsFull() {
		return errors.New("queue is full")
	}

	if q.IsEmpty() {
		q.head = 0
	}

	q.data[q.tail] = data
	q.tail++
}

func (q *Queue) Dequeue() (interface{}, error) {

	/*pseudo code
	-check if the queue is empty
	-return the value pointed by head
	-increase the 'head' index by 1
	-for the last element, reset the values of head and tail to -1
	*/

	if q.IsEmpty() {
		return nil, errors.New("queue is full")
	}

	data := q.data[q.head]

	q.head++

	if q.Size() == 0 {
		q.head = -1
		q.tail = -1
	}
	return data, nil

}

func (q *Queue) IsEmpty() bool {
	if q.head == -1 {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	if q.tail == len(q.data) {
		return true
	}
	return false
}

func (q *Queue) Peek() interface{} {
	return q.data[q.tail]
}

func (q *Queue) Size() int {
	return q.tail - q.head
}
