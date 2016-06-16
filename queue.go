package lfc

import (
	"sync/atomic"
	"unsafe"
)

type Element struct {
	value interface{}
	next  unsafe.Pointer
}

type Queue struct {
	head, tail unsafe.Pointer
	len        int64
}

func NewQueue() *Queue {
	nn := &Element{}
	return &Queue{
		head: unsafe.Pointer(nn),
		tail: unsafe.Pointer(nn),
	}
}

func (q *Queue) Enqueue(value interface{}) {
	nElem := &Element{
		value: value,
	}
	for {
		tail := q.tail
		next := (*Element)(tail).next
		if tail == q.tail {
			if next == nil {
				if atomic.CompareAndSwapPointer(&(*Element)(q.tail).next, next, unsafe.Pointer(nElem)) {
					atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(nElem))
					atomic.AddInt64(&q.len, 1)
					break
				}
			} else {
				atomic.CompareAndSwapPointer(&q.tail, tail, next)
			}
		}
	}
}

func (q *Queue) Dequeue() (value interface{}, ok bool) {
	for {
		head := q.head
		tail := q.tail
		next := (*Element)(head).next
		if head == q.head {
			if head == tail {
				if next == nil {
					return nil, false
				}
				atomic.CompareAndSwapPointer(&q.tail, tail, next)
			} else {
				value := (*Element)(next).value
				if atomic.CompareAndSwapPointer(&q.head, head, next) {
					atomic.AddInt64(&q.len, -1)
					return value, true
				}
			}
		}
	}
}

func (q *Queue) Len() int64 {
	return atomic.LoadInt64(&q.len)
}

//Get n element from queue without delete it. If n > q.Len return all elems.
//Need tests for this function.
func (q *Queue) Get(n uint) (values []interface{}) {
	values = make([]interface{}, 0, n)
	head := atomic.LoadPointer(&q.head)
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

func (q *Queue) GetAll() (values []interface{}) {
	values = make([]interface{}, 0, q.Len())
	next := (*Element)(atomic.LoadPointer(&q.head)).next
	for ; next != nil; next = (*Element)(next).next {
		values = append(values, (*Element)(next).value)
	}
	return
}
