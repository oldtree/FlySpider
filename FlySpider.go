package main

import (
	//"fmt"
//	"FlySpider/messager"
	"runtime"
	"time"
)

func main() {

	//messager.MsgMain()
	runtime.GOMAXPROCS(4)
	TestMain()
	time.Sleep(time.Second * 4)
}
