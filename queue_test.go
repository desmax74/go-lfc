package lfc

import "testing"

func BenchmarkQueueEnq(b *testing.B) {
	Q := NewQueue()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Q.Enqueue(2)
		}
	},
	)
}

func BenchmarkQueueDeq(b *testing.B) {
	Q := NewQueue()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Q.Enqueue(2)
		}
	},
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Q.Dequeue()
		}
	},
	)
}

func TestQueue(t *testing.T) {
	Q := NewQueue()
	Q.Enqueue(52)
	if Q.Len() != 1 {
		t.FailNow()
	}
	res, ok := Q.Dequeue()
	if Q.Len() != 0 {
		t.FailNow()
	}
	if !ok {
		t.FailNow()
	}
	resInt, ok := res.(int)
	if !ok {
		t.FailNow()
	}
	if resInt != 52 {
		t.FailNow()
	}

}
