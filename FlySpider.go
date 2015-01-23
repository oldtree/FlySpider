package main

import (
	"fmt"
	"FlySpider/messager"
	"runtime"
	//"time"
	"FlySpider/Patten"
)


func main() {

    if false{
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
	}else if true{
		Patten.PTest()
	}
	
	
	//TestMain()
	//time.Sleep(time.Second * 4)
}
