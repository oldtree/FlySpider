package messager


import (
	"net"
	"fmt"
	"io"
//	"os"
	"errors"
//	"time"
	"bytes"
//	"math/rand"
)

type Message chan string

const (
	MINCLientNUM = 128
)


type Server struct{
	ClientTable map[int64]Client
	listener net.Listener
	pender  chan net.Conn
	quiter  chan net.Conn
	down    chan bool
}


func NewServer()*Server{
	//ser := &Server{
	//	ClientTable:    make(map[net.Conn]Client,MINCLientNUM),
	//	pender:		 	make(chan net.Conn,16),
	//	quiter:			make(chan net.Conn,16),
		
	//}
	
	ser := new(Server)
	ser.ClientTable = make(map[int64]Client,MINCLientNUM)
	ser.pender = make(chan net.Conn,16)
	ser.quiter = make(chan net.Conn,16)
	
	return ser
}

func (s *Server)CloseServer(){
	s.listener.Close()
}

func (s *Server)Listen(addr_port string){
	var err error
	fmt.Println(addr_port)
	s.listener ,err=net.Listen("tcp","10.0.0.3:8001")
	
	if err!=nil{
		errors.New("server listen port set failure")
	}
	
	defer s.CloseServer()
	
	for {
		nConn ,err := s.listener.Accept()
		
		if err!=nil{
			errors.New("a new connection is wrong")
			continue
		}
		
		//s.pender <- nConn
		go s.RecvMsg(nConn)
	}
}


func (s *Server)JoinClient(){
	
	for {
		select{
			case conn := <-s.pender:
				client := NewClient()
				client.conn  = conn
//				s.ClientTable[conn] = *client
			case state:=<-s.down:
				if state{
					fmt.Printf("server state change to : %s ",state)
			    	break
				}else {
					fmt.Printf("server state change to : %s ",state)
				}
		}
	
	}
}


func (s *Server)LeaveClient(){
	for {
		select{
			case conn := <-s.quiter:
				//delete(s.ClientTable,conn)
				fmt.Println(conn)
			case state:=<-s.down:
			    if state{
					fmt.Printf("server state change to : %s ",state)
			    	break
				}else {
					fmt.Printf("server state change to : %s ",state)
				}
		}
	}
}




func (s *Server)SendMsg(cConn net.Conn,msg []byte)(err error){
	
    num , err:=cConn.Write(msg)
	fmt.Println(num)
	if err!=nil{
		return errors.New("send massage wrong")
	}
	
	return nil
	
}

func (s *Server)RecvMsg(cConn net.Conn)(err error){
	

	var u Shell_message
	fmt.Println("RecvMsg been called ")
	for{
		var msgbyte [128]byte
	    var readBuffer = bytes.NewBuffer(msgbyte[0:])
		var newcoming Client
		le, err:=cConn.Read(msgbyte[0:])
		
		if err!=nil{
			if err != io.EOF{
				return errors.New("read data error")
			}
		}
		
		u.ParseContent(readBuffer.Bytes()[0:le])
		
		if u.Messgeid == LOGIN_MESSAGE_ID{
			var temp UserLogin_message
			temp.ParseContent(u.MessageBody)
			newcoming.conn = cConn
			newcoming.id = temp.Id
			s.ClientTable[newcoming.id] = newcoming 
			fmt.Println(s.ClientTable[newcoming.id])
			s.SendMsg(newcoming.conn,[]byte("grapes"))
		}	
		
		if u.Messgeid == LOGOUT_MESSAGE_ID{
			var temp UserLogin_message
			temp.ParseContent(u.MessageBody)
			delete(s.ClientTable,temp.Id)
		}	
		
	}
	return nil
	
	
}