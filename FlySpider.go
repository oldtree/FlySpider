package main

import (
	"fmt"
	"FlySpider/messager"
	"runtime"
	//"time"
)


func main() {

    if true{
		messager.MsgMain()
		runtime.GOMAXPROCS(4)
		for{
			select{
			
			}
		}
	}else if false{
		Funny()
	}else if false{
		TechMain3()
		fmt.Println("techmain3")
	}
	
	
	//TestMain()
	//time.Sleep(time.Second * 4)
}
