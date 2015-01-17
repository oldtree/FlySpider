package messager


import (
    "fmt"
	"time"

)

func MsgMain(){
	s := NewServer()
	go s.Listen("")
	select{
		case <-time.After(time.Second * 4):
		fmt.Println("CMsgMain function")
		go CMsgMain()
		
		fmt.Println("all run")
		case <-time.After(time.Second * 10):
		break
	}
}