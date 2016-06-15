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
	return s.len
}
