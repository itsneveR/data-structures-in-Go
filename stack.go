package ds

import (
	"errors"
	"fmt"
)

type Stack struct {
	top      int
	cap      int
	stackArr []interface{}
}

/*******************initializer operations***********************/

func New(cap int) *Stack {
	return new(Stack).New(cap)
}

func (s *Stack) New(cap int) *Stack {
	s.top = -1
	s.cap = cap
	s.stackArr = make([]interface{}, cap)
	return s
}

/*******************main operations**********************/

func (s *Stack) Push(data interface{}) error {
	if s.IsFull() {
		return errors.New("over flow")
	}
	s.top++
	s.stackArr[s.top] = data
	return nil
}

//Pop removes the data from the top of the stack

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	value := s.stackArr[s.top]
	s.top--
	return value, nil
}

//Peek returns the data without removing from the top of stac

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}

	return s.stackArr[s.top], nil
}

func (s *Stack) Clear() {
	s.stackArr = nil
	s.top = -1
}

func (s *Stack) Print() {
	fmt.Print('[')
	for i := 0; i <= s.top; i++ {
		fmt.Print(s.stackArr[i])
		if i != s.top {
			fmt.Print(",")
		}
	}
	fmt.Print("[\n")

}

/***************************auxilary operations************************/
func (s *Stack) IsFull() bool {
	return s.top == (s.cap - 1)
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Size() uint {
	return uint(s.top + 1)
}
