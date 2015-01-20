package messager


import (
	"net"
	"fmt"
	"io"
//	"os"
	"errors"
//	"time"
	"bytes"

)



const (
	MINCLientNUM = 1024
)


type Server struct{
	ClientTable map[int64] *Client
	listener net.Listener
	//FromClient  Message
	ClientToCliet  Message
	down    chan bool
}

func (s *Server)CreateClientByServer()*Client{
	return &Client{
		inmsg  :make(chan Shell_message,8),
		outmsg :make(chan Shell_message,8),
		serv   :s,
	}
}


func NewServer()*Server{
	ser := new(Server)
	ser.ClientTable = make(map[int64]*Client,MINCLientNUM)
	//ser.FromClient = make(Message,16*8)
	ser.ClientToCliet = make(Message,8)
	
	return ser
}

func (s *Server)CloseServer(){
	s.listener.Close()
}

func (s *Server)Listen(addr_port string){
	var err error
	fmt.Println(addr_port)
	s.listener ,err=net.Listen("tcp","10.0.0.2:8001")
	
	if err!=nil{
		errors.New("server listen port set failure")
	}
	
	defer s.CloseServer()
	go s.MsgToClient()
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


func (s *Server)MsgToClient(){
	defer fmt.Println("MsgToClient loop down")
	for {
		select{
			case msg := <-s.ClientToCliet:
			    fmt.Println("--------------------")
				fmt.Printf("from %d -----> %d \n",msg.Userid,msg.To)
				fmt.Println("--------------------")
			    //s.ClientTable[msg.To].inmsg<-msg
				 // here will be an empty pointer ,if the target client not online
				if _,ok :=s.ClientTable[msg.To];ok{
					s.ClientTable[msg.To].inmsg<-msg
					//fmt.Println(s.ClientTable)
				}
			case state:=<-s.down:
				if state{
					fmt.Printf("server state change to : %s ",state)
			    	goto End
				}else {
					fmt.Printf("server state change to : %s ",state)
				}
		}
	
	}
End:
    fmt.Println("SendMsg failure")
}

func (s *Server)SendMsg(cConn net.Conn,msg []byte)(err error){
	defer fmt.Println("server's SendMsg down")
    num , err:=cConn.Write(msg)
	fmt.Println(num)
	if err!=nil{
		return errors.New("send massage wrong")
	}
	
	return nil
	
}

func (s *Server)RecvMsg(cConn net.Conn)(err error){
	
    defer fmt.Println("server's RecvMsg down")
	var u Shell_message

	for{
		var msgbyte [256]byte
	    var readBuffer = bytes.NewBuffer(msgbyte[0:])
		//var newcoming Client
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
			
			newcoming := s.CreateClientByServer()
			newcoming.conn = cConn
			newcoming.id = temp.Id
			s.ClientTable[newcoming.id] = newcoming 
			
			//fmt.Println(s.ClientTable[newcoming.id])
			//s.SendMsg(newcoming.conn,[]byte("grapes"))
			
			go s.ClientTable[newcoming.id].ClientCircly()
			fmt.Println("add finished")
			return nil
		}	
		
		
		if u.Messgeid == LOGOUT_MESSAGE_ID{
			var temp UserLogin_message
			temp.ParseContent(u.MessageBody)
			delete(s.ClientTable,temp.Id)
		}	
		
	}
	return nil
	
	
}