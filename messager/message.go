package messager


import (
	"time"

)

func MsgMain(){
	s := NewServer()
	go s.Listen("")
	select{
		case <-time.After(time.Second * 4):
		go CMsgMain1()
		go CMsgMain2()
		
		case <-time.After(time.Second * 10):
		break
	}
}