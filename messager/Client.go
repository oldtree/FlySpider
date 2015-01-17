package messager


import (
	"net"
	"fmt"
	"time"
	"encoding/json"
	"bytes"
	"io"
	"errors"
)


type Client struct{
	conn net.Conn
	inmsg Message
	outmsg Message
	name string
	id   int64
}

func NewClient()*Client{
	return &Client{
		inmsg:make(chan string ,8),
		outmsg :make(chan string,8),
		name :"hello",
	}
	
}

var tick = time.NewTicker(time.Second)


func CMsgMain(){

	
	var u Shell_message
	var uu UserLogin_message
	
	conn ,err:=net.Dial("tcp","10.0.0.3:8001")
	defer conn.Close()
	
	uu.Id = 1234
	uu.Login = true
	uu.Date = "20150116"
	
	u.Userid = 1234
	u.Messgeid = 128
	u.MessageBody,_ = json.Marshal(uu)
	
	msg ,err :=json.Marshal(u)
	
	if err!=nil{
		fmt.Println(err)
	}
	for{
		
		var msgbyte [128]byte
	    var readBuffer = bytes.NewBuffer(msgbyte[0:])
		
		select{
			case <-tick.C:
			conn.Write(msg)	
		    conn.Read(msgbyte[0:])
		
			if err!=nil{
				if err != io.EOF{
					errors.New("read data error")
				}
			}
		}
		fmt.Println(string(readBuffer.String()))
	}
	
}
