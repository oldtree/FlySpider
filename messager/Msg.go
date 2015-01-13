package messager


import (
	"net"
	"fmt"
	"io"
	"os"
)

type Point struct{
	PointType string
	Address string
	
}

var ConnManager map[string]Point

func SMsgMain(){
	
	conn ,err :=net.Listen("tcp","10.0.0.3:8001")
	if err!=nil{
		fmt.Println(err)
	}
	defer conn.Close()
	for {
		var b []byte
		c ,err:=conn.Accept()
		if err!=nil{
		    fmt.Println(err)
	    }
			fmt.Println(c.Read(b))
			fmt.Println(string(b))
			io.Copy(os.Stdout,c)
	}
}
