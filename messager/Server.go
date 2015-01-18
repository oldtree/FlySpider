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



const (
	MINCLientNUM = 128
)


type Server struct{
	ClientTable map[int64]Client
	listener net.Listener
	FromClient  Message
	ToCliet  Message
	down    chan bool
}


func NewServer()*Server{
	ser := new(Server)
	ser.ClientTable = make(map[int64]Client,MINCLientNUM)
	ser.FromClient = make(Message,16*8)
	ser.ToCliet = make(Message,16*8)
	
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
		fmt.Println("before call go s.RecvMsg(nConn)")
		go s.RecvMsg(nConn)
		fmt.Println("after call go s.RecvMsg(nConn)")
	}
}


func (s *Server)MsgFromClient(){
	
	for {
		select{
			case msg := <-s.FromClient:
			    s.ClientTable[msg.Userid].inmsg<-msg
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


func (s *Server)MsgToClient(){
	for {
		select{
			case msg := <-s.ToCliet:
			    s.ClientTable[msg.Userid].outmsg<-msg
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