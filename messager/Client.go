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

var tick = time.NewTicker(time.Second)


type Message chan Shell_message

type Client struct{
	conn net.Conn
	inmsg Message  // come from server
	outmsg Message // send into server
	name string
	id   int64
	serv *Server
}

func (c *Client)ClientCircly(){
	go c.RecvMsg()
	go c.SendMsg()
	go c.PassMsgToServer()
}

func (c *Client)PassMsgToServer(){
	for{
		select{
			case c.serv.ClientToCliet<-<-c.outmsg:
				fmt.Println("pass by message")
			//case msg:=<-c.outmsg:
			//    c.serv.ClientToCliet<-msg
			//	fmt.Println(msg)
		}
	}
}

func (c *Client)SendMsg(){
	defer c.Close()
	for{
		select{
			case msg:=<-c.inmsg:
			    data ,err := json.Marshal(msg)
				le , err:=c.conn.Write(data)
				if err!=nil{
					fmt.Println("send massage error")
					goto End
				}
				fmt.Printf("message len : %d   \n " , le)
		}
	}
End:
    fmt.Println("SendMsg failure")
}

func (c *Client)RecvMsg(){
	
	defer c.Close()
	for {
		var msgbyte [256]byte
		var readBuffer = bytes.NewBuffer(msgbyte[0:])
		var msg Shell_message
		le ,err:=c.conn.Read(msgbyte[0:])
		if err!=nil{
			if err != io.EOF{
				fmt.Println(errors.New("read data error"))
				break
			}
		}

		err = json.Unmarshal(readBuffer.Bytes()[0:le],&msg)
		if err!=nil{
			fmt.Println("recv message ok ,Unmarshal wrong")
		}
		if msg.Userid == msg.To{
			c.inmsg <-msg
		}else{
			c.outmsg<-msg
		}
		
	}
}

func (c *Client)Close(){
	close(c.inmsg)
	close(c.outmsg)
	c.conn.Close()
}


func NewClient()*Client{
	return &Client{
		inmsg  :make(chan Shell_message ,8),
		outmsg :make(chan Shell_message,8),
	}
	
}


func CMsgMain1(){

	
	var u Shell_message
	var uu UserLogin_message
	
	conn ,err:=net.Dial("tcp","10.0.0.2:8001")
	defer conn.Close()
	
	uu.Id = 1234
	uu.Login = true
	uu.Date = "20150116"
	
	u.Userid = 1234
	u.Messgeid = 128
	u.To = 12345
	u.MessageBody,_ = json.Marshal(uu)
	
	msg ,err :=json.Marshal(u)
	
	if err!=nil{
		fmt.Println(err)
	}
	for{
		var msgbyte [256]byte
	    var readBuffer = bytes.NewBuffer(msgbyte[0:])
		
		select{
			case <-tick.C:
			conn.Write(msg)	
		    //go conn.Read(msgbyte[0:])
		    
			if err!=nil{
				if err != io.EOF{
					errors.New("read data error")
				}
			}
		}
		fmt.Println(string(readBuffer.String()))
	}
	
}

func CMsgMain2(){

	
	var u Shell_message
	var uu UserLogin_message
	
	conn ,err:=net.Dial("tcp","10.0.0.2:8001")
	defer conn.Close()
	
	uu.Id = 12345
	uu.Login = true
	uu.Date = "20150116"
	
	u.Userid = 12345
	u.Messgeid = 128
	u.To = 1234
	u.MessageBody,_ = json.Marshal(uu)
	
	msg ,err :=json.Marshal(u)
	
	if err!=nil{
		fmt.Println(err)
	}
	for{
		var msgbyte [256]byte
	    var readBuffer = bytes.NewBuffer(msgbyte[0:])
		
		select{
			case <-tick.C:
			conn.Write(msg)	
		    //go conn.Read(msgbyte[0:])
		    
			if err!=nil{
				if err != io.EOF{
					errors.New("read data error")
				}
			}
		}
		fmt.Println(string(readBuffer.String()))
	}
	
}
