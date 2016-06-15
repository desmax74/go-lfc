package lfc

import "testing"

func BenchmarkStackPush(b *testing.B) {
	S := NewStack()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			S.Push(2)
		}
	})
}

func BenchmarkStackPop(b *testing.B) {
	S := NewStack()
	for i := 0; i < b.N; i++ {
		S.Push(2)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			S.Pop()
		}
	})
}

func TestStack(t *testing.T) {
	S := NewStack()
	S.Push(23)
	if S.Len() != 1 {
		t.FailNow()
	}
	res, ok := S.Pop()
	if !ok {
		t.FailNow()
	}
	resInt, ok := res.(int)
	if !ok {
		t.FailNow()
	}
	if resInt != 23 {
		t.FailNow()
	}
}
