package messager

import (
	"encoding/json"
	//"reflect"
)

type messager interface{
	ParseContent(jsonstream []byte)
}

const (
	BROADCAST_MESSAGE_ID  = 0
	PEER2PEER_MESSAGE_ID  = 1
	PEER2N_MESSAGE_ID  = 2
	N2PEER_MESSAGE_ID  = 3
	
	
	LOGIN_MESSAGE_ID  = 128
	LOGINOK_MESSAGE_ID  = 129
	LOGOUT_MESSAGE_ID = 130
	LOGOUTOK_MESSAGE_ID = 131
)
type Shell_message struct{
	Userid int64 `json:"userid"`
	To int64 `json:"to"`
	Messgeid int `json:"messgeid"`
	MessageBody []byte `json:"messagebody"`
}

func (this *Shell_message)ParseContent(jsonstream []byte){
	err := json.Unmarshal(jsonstream, this)
	
	if err != nil {
		panic(err)
	}
	
	return
}


type UserLogin_message struct{
	Id      int64 `json:"id"`
	Login   bool `json:"login"`
	Date    string`json:"date"`
}
func (this *UserLogin_message)ParseContent(jsonstream []byte){
	err := json.Unmarshal(jsonstream, this)
	if err != nil {
		panic(err)
	}
	
	return
}
type UserPost_message struct{
	Id      int64 `json:"id"`
	Space   int `json:"space"`
	To      int64  `json:"to"`
	Date    string`json:"date"`
}

func (this *UserPost_message)ParseContent(jsonstream []byte){
	err := json.Unmarshal(jsonstream, this)
	if err != nil {
		panic(err)
	}
	return
}
type UserGet_message struct{
	Id      int64 `json:"id"`
	Space   int `json:"space"`
	To      string  `json:"to"`
	Date    string`json:"date"`
}
func (this *UserGet_message)ParseContent(jsonstream []byte){
	err := json.Unmarshal(jsonstream, this)
	if err != nil {
		panic(err)
	}
	return
}
 