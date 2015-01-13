package messager


import (
	"net"
	"fmt"
)

func CMsgMain(){
	conn ,err:=net.Dial("tcp","10.0.0.3:8001")
	
	if err!=nil{
		fmt.Println(err)
	}
	
	conn.Write([]byte("hello \n"))	
	defer conn.Close()
}