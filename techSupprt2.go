package main

import (
//	"fmt"
	"time"
)

var ccc = make(chan int ,1)
var tick = time.NewTicker(1 * time.Second) 
func Send(){
	for {
		select{
			case <-tick.C:
			    ccc<-time.Now().Second()
			
		}
	}
}


func Resc(){
	for{
		select {
			case <-time.After(time.Second*6):
				close(ccc)
			//case temp:=<-ccc:
			//	fmt.Println(temp)
			
		}
	}
}

func Funny(){
	go Send()
	go Resc()
}
/*
type Task func() error

type Task interface {
	Do() error
}

type TaskFunc func() error

func (f TaskFunc) Do() error {
	return f()
}

type TaskSet []Task

func (ts TaskSet) Do() error {
	for _, task := range ts {
		if err := task.Do(); err != nil {
			return err
		}
	}
	return nil
}

type Worker <-chan Task

func (w Worker) Run() (done chan<- error) {
	go func() {
		for task := range w {
			done <- task.Do()
		}
	}()
	return
}

type FileTask struct {
	Task
	Input  string
	Output string
}

func (ft *FileTask) Do() error {
}*/