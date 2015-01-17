package main

import (
	"runtime"
	"strconv"
	"testing"
)

func TestSharemapSync(t *testing.T) {
	share := NewShareMap()
	for i := 0; i < 10; i++ {
		share.AddOrSet(strconv.Itoa(i), "world")
	}
	for i := 0; i < 10; i++ {
		temp, _ := share.Get(strconv.Itoa(i))
		if temp != "world" {
			t.Error("get function wrong")
		}
	}
	share.Clear()
}

func TestSharemapChannel(t *testing.T) {
	runtime.GOMAXPROCS(3)
	share := NewShareMap()
	var temp string
	go share.WorkerQueue()
	for i := 0; i < 10; i++ {
		temp = strconv.Itoa(i)
		f := func() { share.AddOrSetUnsafe(temp, "world") }
		share.task <- f
	}
	share.Clear()
	close(share.task)
}

func BenchmarkShareMapSync(tb *testing.B) {
	runtime.GOMAXPROCS(3)
	share := NewShareMap()
	var temp string
	for i := 0; i < tb.N; i++ {
		temp = strconv.Itoa(i)
		share.AddOrSet(temp, "world")
		share.Get(temp)
		share.Remove(temp)
	}
	share.Clear()
}

func BenchmarkShareMapChannel(tb *testing.B) {
	runtime.GOMAXPROCS(3)
	share := NewShareMap()
	var temp string
	go share.WorkerQueue()
	for i := 0; i < tb.N; i++ {
		temp = strconv.Itoa(i)
		f := func() { share.AddOrSetUnsafe(temp, "world") }
		share.task <- f
	}
	share.Clear()
	close(share.task)
}
