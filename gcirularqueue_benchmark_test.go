/*
@Time : 2020-12-11 08:20
@Author : kc
@File : gcirularqueue_benchmark
@Software: GoLand
*/
package gcircularqueue

import (
	"testing"
)

func BenchmarkCircularQueueThreadSafe_PushKick(b *testing.B) {
	cq := NewCircularQueueThreadSafe(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cq.PushKick(i)
	}
	for true {
		v := cq.Shift()
		if v == nil {
			break
		}
	}
}
func BenchmarkCircularQueue_PushKick(b *testing.B) {
	cq := NewCircularQueue(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cq.PushKick(i)
	}
	for true {
		v := cq.Shift()
		if v == nil {
			break
		}
	}
}


func BenchmarkCircularQueueThreadSafe_ShiftAll1(b *testing.B) {
	cq := NewCircularQueueThreadSafe(1000)
	for i := 0; i < 1000000; i++ {
		cq.PushKick(i)
	}
	b.ResetTimer()
	cq._ShiftAll1()
}

func BenchmarkCircularQueueThreadSafe_ShiftAll2(b *testing.B) {
	cq := NewCircularQueueThreadSafe(1000)
	for i := 0; i < 100000; i++ {
		cq.PushKick(i)
	}
	b.ResetTimer()
	cq._ShiftAll2()
}

func BenchmarkCircularQueueThreadSafe_ShiftAll3(b *testing.B) {
	cq := NewCircularQueueThreadSafe(1000)
	for i := 0; i < 100000; i++ {
		cq.PushKick(i)
	}
	b.ResetTimer()
	cq._ShiftAll3()
}

func BenchmarkSliceAppendInitedSize(b *testing.B) {
	a := make([]int, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = append(a, i)
	}
}

func BenchmarkSliceAppendZeroSize(b *testing.B) {
	a := make([]int, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = append(a, i)
	}
}