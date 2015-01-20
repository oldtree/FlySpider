package main

import(
	"log"
	"log/syslog"
	"io"
	"reflect"
	"fmt"
	"time"
//	"io/ioutil"
)


func LogFunction(){
	sl,err:=syslog.New(syslog.LOG_ERR,"FlySpider")
	//sl,err:=syslog.Dial("udp","localhost",syslog.LOG_ERR,"FlySpider")
	defer sl.Close()
	
	if err!=nil{
		log.Fatal("error")
	}
	
	sl.Alert("hello")
	sl.Crit("crit")
	sl.Err("Error")
	sl.Warning("waring")
	sl.Notice("notice")
	sl.Info("info")
	sl.Write([]byte("write"))
}


type ByteReader byte

func (b ByteReader)Read(buf []byte)(int ,error){
	for i := range buf{
		buf[i] = byte(b)
	}
	fmt.Println(reflect.TypeOf(b).Name())
	return len(buf),nil
}


type LogReader struct{
	
	io.Reader
	ByteReader
}


func (l LogReader)Read(buf [] byte)(int ,error){
	n , err := l.Reader.Read(buf)
	log.Printf("read %d bytes ,error :%v ",n,err )
	fmt.Println(reflect.TypeOf(l).Name())
	return n ,err
}


func Readfuntion(){
	r := LogReader{ByteReader('A'),ByteReader('A')}
	fmt.Println(reflect.TypeOf(r))
	b :=make([]byte,10)
	r.Read(b)
	log.Printf("b : %q ",b)
}


func ParamChan( param chan int){
	var temp = 0 
	for{
		temp ++ 
		select{
			case param<-temp:
				
				if temp > 500 {
					goto LL
				}
			case <-time.After(time.Second *3): // will be never going here stupid
			    goto LL
		}
	}
LL:
	fmt.Println(temp)
}



func ParamChanTest(){
	var p  = make(chan int,1)
	go ParamChan(p)
	fmt.Println(reflect.TypeOf(p))
	
	for{
		select{
			case v:=<-p:
			    fmt.Println(v)
			case <-time.After(time.Second *3):
			    fmt.Println("hell")
			    goto LL
		}
	}
LL:
	fmt.Println("mession down")

}


func TechMain3(){
	ParamChanTest()
}