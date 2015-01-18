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
//每次我只是把这两天听的歌的歌名挂在那里


func (c *Client)SendMsg(){
	defer c.Close()
	for{
		select{
			case msg:=<-c.inmsg:
			    data ,err := json.Marshal(msg)
				le , err:=c.conn.Write(data)
				if err!=nil{
					fmt.Println("send massage error")
					break
				}
				fmt.Printf("message len : %d   \n " , le)
		}
	}
}

func (c *Client)RecvMsg(){
	
	defer c.Close()
	for {
		var msgbyte [128]byte
		var readBuffer = bytes.NewBuffer(msgbyte[0:])
		var msg Shell_message
		le ,err:=c.conn.Read(msgbyte[0:])
		if err!=nil{
			if err != io.EOF{
				fmt.Println(errors.New("read data error"))
				break
			}
		}
		err = json.Unmarshal(readBuffer.Bytes()[0:le],msg)
		if err!=nil{
			fmt.Println("recv message ok ,Unmarshal wrong")
		}
		c.outmsg<-msg
	}
}

func (c *Client)Close(){
	c.conn.Close()
}


func NewClient()*Client{
	return &Client{
		inmsg:make(chan Shell_message ,8),
		outmsg :make(chan Shell_message,8),
	}
	
}


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
