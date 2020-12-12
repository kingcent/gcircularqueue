package gcircularqueue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQueueThreadSafeInit(t *testing.T) {
	cq := NewCircularQueueThreadSafe(5)
	if cq.capacity != 6 {
		t.Error("It has wrong size:", cq.capacity)
	}

	if !cq.IsEmpty() {
		t.Error("It is not empty")
	}
}

func TestThreadSafePush(t *testing.T) {
	cq := NewCircularQueueThreadSafe(1)
	firstElement := "First"
	cq.Push(firstElement)
	if cq.elements[0] != firstElement {
		t.Error("It's Push Func is wrong, the first elment is:", cq.elements[0])
	}
}

func TestThreadSafeShift(t *testing.T) {
	cq := NewCircularQueueThreadSafe(5)
	firstElement := "First"
	cq.Push(firstElement)
	e := cq.Shift()
	if e != firstElement {
		t.Error("It can not shift the first element")
	}
	// shift from a empty queueThreadSafe
	e = cq.Shift()
	if e != nil {
		t.Errorf("Can not shift from a empty queueThreadSafe, it is:%v", e)
	}
}

func TestThreadSafeIsEmpty(t *testing.T) {
	cq := NewCircularQueueThreadSafe(5)
	if !cq.IsEmpty() {
		t.Error("It's IsEmpty Func is wrong")
	}
}

func TestThreadSafeIsFull(t *testing.T) {
	cq := NewCircularQueueThreadSafe(1)
	cq.Push("First")
	if !cq.IsFull() {
		t.Error("It's IsFull is wrong")
	}

}

func TestThreadSafeFIFO(t *testing.T) {
	cq := NewCircularQueueThreadSafe(3)
	cq.Push(1)
	cq.Push(2)
	firstElement := cq.Shift()
	if firstElement != 1 {
		t.Error("It doesn't support FIFO")
	}
}

func TestThreadSafeCirculartAility(t *testing.T) {
	cq := NewCircularQueueThreadSafe(3)
	cq.Push(1)
	cq.Push(2)
	cq.Push(3)
	cq.Shift()
	cq.Push(3)
}

func TestCircularQueueThreadSafe_PushKick(t *testing.T) {
	size := 1000
	l := 1000000
	cq := NewCircularQueueThreadSafe(size)
	for i := 1; i < l; i++ {
		cq.PushKick(i)
	}
	s := l%size + (int(l/size)-1)*size
	for i := s; i < s+size; i++ {
		v := cq.Shift()
		if v.(int) != i {
			t.Error("error value ", i, v)
		}
	}
}

func TestCircularQueueThreadSafe_SizeCorrect(t *testing.T) {
	Convey("size should correct 1", t, func() {
		size := 1000
		l := 1000000
		cq := NewCircularQueueThreadSafe(size)
		for i := 1; i < l; i++ {
			cq.PushKick(i)
		}
		So(cq.Len() == size, ShouldBeTrue)
	})
	Convey("size should correct 2", t, func() {
		size := 1000
		l := 100
		cq := NewCircularQueueThreadSafe(size)
		for i := 0; i < l; i++ {
			cq.PushKick(i)
		}
		So(cq.Len() == l, ShouldBeTrue)
	})
	Convey("size should correct 3", t, func() {
		size := 1000
		l := 1000000
		cq := NewCircularQueueThreadSafe(size)
		for i := 1; i < l; i++ {
			cq.PushKick(i)
		}
		cq.Shift()
		cq.Shift()
		So(cq.Len() == size - 2, ShouldBeTrue)
	})
	Convey("size should correct 4", t, func() {
		size := 1000
		l := 1000000
		cq := NewCircularQueueThreadSafe(size)
		for i := 1; i < l; i++ {
			cq.PushKick(i)
		}
		cq.ShiftAll()
		So(cq.Len() == 0, ShouldBeTrue)
	})
}

func TestCircularQueueThreadSafe_ShiftAll(t *testing.T) {
	Convey("ShiftAll", t, func() {
		size := 1000
		l := 1000000
		cq := NewCircularQueueThreadSafe(size)
		for i := 0; i < l; i++ {
			cq.PushKick(i)
		}
		all := cq.ShiftAll()
		So(len(all) == 1000, ShouldBeTrue)
		So(cq.Len() == 0, ShouldBeTrue)
		cq.PushKick(2134)
		cq.PushKick(323)
		So(cq.Len() == 2, ShouldBeTrue)
		all = cq.ShiftAll()
		So(len(all) == 2, ShouldBeTrue)
	})
}

