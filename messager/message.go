package messager


import (

	"time"

)



func MsgMain(){
	
	
	go SMsgMain()
	
	for {
		select{
			case <-time.After(10):
			go CMsgMain()
			case <-time.After(100):
			break
		}
	}
	
	
}