// FlySpider project FlySpider.go
package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
//	"bytes"
	"encoding/json"
	"strconv"
	"flag"
)


var downloadPath = flag.String("-dpath",".","指定一个下载文件保存")
var targetUrl =flag.String("-durl","http://news-at.zhihu.com/api/3/news/latest","指定一个下载url")
var DOWNLOAD = downloadPath


var ZhihunewsUrl = "http://news-at.zhihu.com/api/3/news/latest"
//var ZhihunewsUrl ="http://news.at.zhihu.com/api/3/news/before/20131119"
var patch_shareUrl = "http://daily.zhihu.com/story/"

//知乎新闻格式
type TodayNews struct{
	Date string `json:"date"`
	Story []Item  `json:"stories"`
	TopStory []Item  `json:"top_stories"`
} 



type Item struct {
	Theme_name     string `json:"theme_name"`
	Subscribed bool `json:"subscribed"`
	Title string `json:"title"`
	Image string `json:"image"`
	Share_url string `json:"share_url"`
	Ga_prefix      string `json:"ga_prefix"`
	Theme_id int64 `json:"theme_id"`
	Images []string `json:"images"`
	Multipic bool `json:"multipic"`
	Typet      int `json:"type"`
	Id    int `json:"id"`
}



func Creator(name string) {
	file ,err :=os.Getwd()
	sep :=strings.LastIndex(name,"/")
	
	err = os.Mkdir(name, 777)
	
	if err != nil {
		panic(err)
	}
}

//将byte流转化为一个结构体
func ParseContent(jsonstream []byte) *TodayNews{
	var content TodayNews
	err :=json.Unmarshal(jsonstream,&content)
	if err!=nil{
		panic( err)
	}
	return &content
}

//判断文件状态
func isExist(path string)bool{
	exist ,err:=os.Stat(path)
	if err!=nil{
		return os.IsExist(err)
	}
	fmt.Println("|^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^|")
	fmt.Println("|file infomation :")
	fmt.Println("|   ",exist.Name())
	fmt.Println("|   ",exist.ModTime())
	fmt.Println("|   ",exist.Size())
	fmt.Println("|   ",exist.IsDir())
	fmt.Println("|____________________________________|")
	return true
}

//根据新闻项的id和固定URL来取得新闻内容，并保存为.html文件
func GetStoryContent(it Item){
	var content []byte
	var filepath string
	defer func(){
		v := recover()
		if v!=nil{
			fmt.Println(v)
		}
	}()
	body ,err:=http.Get(patch_shareUrl+strconv.Itoa(it.Id))
	if err!=nil{
		panic(err)
	}
	filepath = DOWNLOAD+it.Title+it.Theme_name+".html"
	
	content,err = ioutil.ReadAll(body.Body)
    
	if isExist(filepath){
		fmt.Println("file already exist")
	}else{
		file ,err:=os.Create(filepath)
		file.Write(content)
		if err!=nil{
			panic(err)
		}
	}
	
}

//t := time.Date(2013, time.May, 20, 23, 0, 0, 0, time.UTC); t.Before(time.Now()); t = t.AddDate(0, 0, 1)
func GetPageBody(urlpath string){
	urlpath = ZhihunewsUrl
	if urlpath=""{
		fmt.Println("url path are null ")
		return
	}
	resp, err := http.Get(urlpath)
	defer func(){
		v := recover()
		if v!=nil{
			fmt.Println(v)
		}
	}()
	var content *TodayNews
	var imagePath string
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	content = ParseContent([]byte(body))
	fmt.Println(content.Date)
	for _,item :=range  content.Story{
		resp ,err :=http.Get(item.Images[0])
		imgBody ,err:=ioutil.ReadAll(resp.Body)
		if err!=nil{
			panic(err)
		}
		GetStoryContent(item)
		temp := strings.Split(item.Images[0],"/")
		imagePath = DOWNLOAD+temp[len(temp)-1]
		if isExist(imagePath){
			fmt.Println("file already exist")
		}else{
			file ,err:=os.Create(imagePath)
			if err!=nil{
				panic(err)
			}
			file.Write(imgBody)
		}
			
	}
	for _,item :=range content.TopStory{
		fmt.Println(item.Image)
		resp ,err :=http.Get(item.Image)
		if err!=nil{
			panic(err)
		}
	}
	return
}


func main(){
	GetPageBody("http://news-at.zhihu.com/api/3/news/latest")
}
