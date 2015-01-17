package main

import (
	//"fmt"
	"FlySpider/messager"
	"runtime"
	//"time"
)


func main() {

    if true{
		messager.MsgMain()
	}else if false{
		Funny()
	}
	
	runtime.GOMAXPROCS(4)
	for{
		select{
			
		}
	}
	//TestMain()
	//time.Sleep(time.Second * 4)
}
