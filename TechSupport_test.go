package main

import(
	"testing"
	"strconv"
	"runtime"
)


func TestSharemap(t *testing.T){
	share :=NewShareMap()
	for i:=0;i<10;i++{
		share.AddOrSet(strconv.Itoa(i),"world")
	}
	for i:=0;i<10;i++{
		temp,_:=share.Get(strconv.Itoa(i))
		if temp!="world"{
			t.Error("get function wrong")
		}
	}
	//share.Clear()
}

func BenchmarkShareMap(tb*testing.B){
	runtime.GOMAXPROCS(2)
	share :=NewShareMap()
	for i:=0;i<tb.N;i++{
		go share.AddOrSet(strconv.Itoa(i),"world")
		go share.Get(strconv.Itoa(i))
		go share.Remove(strconv.Itoa(i))
	}
	share.Clear()
}