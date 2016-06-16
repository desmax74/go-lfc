package lfc

import (
	"sync/atomic"
	"unsafe"
)

type Stack struct {
	head unsafe.Pointer
	len  int64
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value interface{}) {
	nElem := &Element{}
	nElem.value = value
	for {
		nElem.next = s.head
		if atomic.CompareAndSwapPointer(&s.head, nElem.next, unsafe.Pointer(nElem)) {
			atomic.AddInt64(&s.len, 1)
			return
		}
	}
}

func (s *Stack) Pop() (value interface{}, ok bool) {
	for {
		oldHead := s.head
		if oldHead == nil {
			return nil, false
		}
		if atomic.CompareAndSwapPointer(&s.head, oldHead, (*Element)(oldHead).next) {
			atomic.AddInt64(&s.len, -1)
			return (*Element)(oldHead).value, true
		}
	}
}

func (s *Stack) Len() int64 {
	return atomic.LoadInt64(&s.len)
}

//Get n element from queue without delete it. If n > s.Len return all elems.
//Need tests for this function.
func (s *Stack) Get(n uint) (values []interface{}) {
	values = make([]interface{}, 0, n)
	head := atomic.LoadPointer(&s.head)
	for ; n > 0; n-- {
		next := (*Element)(head).next
		if next == nil {
			return
		}
		values = append(values, (*Element)(next).value)
		head = next
	}
	return
}

func (s *Stack) GetAll() (values []interface{}) {
	values = make([]interface{}, 0, s.Len())
	next := (*Element)(atomic.LoadPointer(&s.head)).next
	for ; next != nil; next = (*Element)(next).next {
		values = append(values, (*Element)(next).value)
	}
	return
}
