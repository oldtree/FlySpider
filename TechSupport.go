package main


import (
    "fmt"
    "math/rand"
	"time"
	"sync"
	"encoding/json"
)

func test(ch chan int, i int) {
    ch <- rand.Int()
    fmt.Println(i, "go...")
}

func TechMain() {
    chs := make([]chan int, 10)
    for i := 0; i < 10; i++ {
        chs[i] = make(chan int)
        go test(chs[i], i)
    }
    for _, ch := range chs {
        value := <-ch
        fmt.Print(value)
		fmt.Println("----")
    }
    var i int
    fmt.Scan(&i)
    fmt.Println("done")
}


var chans = make(chan int,10)

func OpChans(){
	for i:=1;i<10;i++{
		
		test(chans,i)
	}
	for i:=1;i<10;i++{
		fmt.Println(<-chans)
	}
	
}



type F func(op1 int,op2 int)

var Opqueue = make(chan F,10)


func Add(op1 int,op2 int){
	var re = op1 + op2
	fmt.Printf(" %d + %d = %d  \n",op1,op2,re)
}

func Sub(op1 int,op2 int){
	var re = op1 - op2
	fmt.Printf(" %d - %d = %d  \n",op1,op2,re)
}

func Double(op1 int,op2 int){
	var re = op1 * op2
	fmt.Printf(" %d * %d = %d  \n",op1,op2,re)
}

func Worker(){
	var funcCall F
	for{
		select{
			case funcCall =<-Opqueue:
			    fmt.Printf("funcCall is  :%p  \n",funcCall)
				funcCall(4,2)
			case <-time.After(time.Second*3):
			    fmt.Println("bad commander!")
		}
	}
	fmt.Println("worker exit!")
	
}

func Commander(){
	var funcCall1 = Add
	var funcCall2 = Sub
	var funcCall3 = Double
	
	Opqueue<-funcCall1
	Opqueue<-funcCall1
	Opqueue<-funcCall2
	Opqueue<-funcCall1
	Opqueue<-funcCall3
	Opqueue<-funcCall1
	Opqueue<-funcCall2
	Opqueue<-funcCall3
	Opqueue<-funcCall2
}



func TestMain(){
	go Commander()
	go Worker()
}

type WorkTask func()


type ShareMap struct{
	list map[string]interface{}
	sync.RWMutex
	task chan WorkTask
}

func (sh * ShareMap)WorkerQueue(){
	
	var f WorkTask
	for {
		select{
			case f=<-sh.task:
			    go f()
			case <-time.After(time.Second*10):
			    fmt.Println("bad commander!")
		}
	}
	for f:=range sh.task{
		go f()
	}
}




func NewShareMap()(*ShareMap){
	sh :=new(ShareMap)
	sh.list = make(map[string]interface{})
	sh.task = make(chan WorkTask,100)
	return sh
}


func (sh *ShareMap)ClearUnsafe()(suc bool){
	suc = false
	for k:= range sh.list{
		delete(sh.list,k)
	}
	suc =true
	return suc
}


func (sh *ShareMap)GetUnsafe(key string)(elem interface{},suc bool){
	elem,suc =sh.list[key]
	if suc{
		return 
	}
	return nil,false
}

func (sh *ShareMap)AddOrSetUnsafe(key string,elem interface{})(suc bool){
	sh.list[key] = elem	
	return true
}

func (sh *ShareMap)RemoveUnsafe(key string)(elem interface{},suc bool){	
	elem,suc=sh.Get(key)
	if suc{
		delete(sh.list,key)
		return 
	}
	
	return nil,true
}


func (sh *ShareMap)CountUnsafe()(num int){
	return len(sh.list)
}


func (sh *ShareMap)ToJsonUnsafe()(re []byte,suc bool){	
	re ,err:=json.Marshal(sh)
	if err!=nil{
		return nil,false
	}
	
	return re,true
}




func (sh *ShareMap)Clear()(suc bool){
	sh.Lock()
	//defer sh.Unlock()
	suc = false
	for k:= range sh.list{
		delete(sh.list,k)
	}
	sh.Unlock()
	suc =true
	return suc
}


func (sh *ShareMap)Get(key string)(elem interface{},suc bool){
	sh.RLock()
	//defer sh.RUnlock()
	
	elem,suc =sh.list[key]
	if suc{
		return 
	}
	sh.RUnlock()
	
	return nil,false
}

func (sh *ShareMap)AddOrSet(key string,elem interface{})(suc bool){
	sh.Lock()
	//defer  sh.Unlock()

	sh.list[key] = elem
	sh.Unlock()
	return true
}

func (sh *ShareMap)Remove(key string)(elem interface{},suc bool){
	sh.RLock()
	//defer sh.RUnlock()
	
	elem,suc=sh.Get(key)
	if suc{
		
		delete(sh.list,key)
		return 
	}
	sh.RUnlock()
	return nil,true
}


func (sh *ShareMap)Count()(num int){
	var length int
	sh.Lock()
	//defer sh.Unlock()
	length = len(sh.list)
	sh.Unlock()
	return length
}



func (sh *ShareMap)IsEmpty()(empty bool){
	if sh.Count() == 0{
		return true
	}
	return false
}

func (sh *ShareMap)ToJson()(re []byte,suc bool){
	sh.Lock()
	//defer sh.Unlock()
	
	re ,err:=json.Marshal(sh)
	if err!=nil{
		return nil,false
	}
	sh.Unlock()
	return re,true
}








//here we go
