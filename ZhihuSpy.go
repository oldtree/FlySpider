// FlySpider project FlySpider.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//	"regexp"
	"strings"
	//	"bytes"
	"encoding/json"
	"flag"
	"strconv"
	"time"
)

var downloadPath = flag.String("-dpath", "download/", "指定一个下载文件保存")
var targetUrl = flag.String("-durl", "http://news-at.zhihu.com/api/3/news/latest", "指定一个下载url")
var DOWNLOAD = *downloadPath

var ZhihunewsUrl = "http://news-at.zhihu.com/api/3/news/latest"
var ZhihuBeforeNewsUrl = "http://news.at.zhihu.com/api/3/news/before/"
var patch_shareUrl = "http://daily.zhihu.com/story/"

//知乎新闻格式
type TodayNews struct {
	Date     string `json:"date"`
	Story    []Item `json:"stories"`
	TopStory []Item `json:"top_stories"`
}

type Item struct {
	Theme_name string   `json:"theme_name"`
	Subscribed bool     `json:"subscribed"`
	Title      string   `json:"title"`
	Image      string   `json:"image"`
	Share_url  string   `json:"share_url"`
	Ga_prefix  string   `json:"ga_prefix"`
	Theme_id   int64    `json:"theme_id"`
	Images     []string `json:"images"`
	Multipic   bool     `json:"multipic"`
	Typet      int      `json:"type"`
	Id         int      `json:"id"`
}

func Creator(name string) error {
	//	file ,err :=os.Getwd()
	//sep :=strings.LastIndex(name,"/")
	var err error
	if isExist(name) {
		fmt.Println("file already exist")
	} else {
		err = os.Mkdir(DOWNLOAD+"/"+name, os.ModePerm)
	}
	return err
}

//将byte流转化为一个结构体
func ParseContent(jsonstream []byte) *TodayNews {
	var content TodayNews
	err := json.Unmarshal(jsonstream, &content)
	if err != nil {
		panic(err)
	}
	return &content
}

//判断文件状态
func isExist(path string) bool {
	exist, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	fmt.Println("|^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^|")
	fmt.Println("|file infomation :")
	fmt.Println("|   ", exist.Name())
	fmt.Println("|   ", exist.ModTime())
	fmt.Println("|   ", exist.Size())
	fmt.Println("|   ", exist.IsDir())
	fmt.Println("|____________________________________|")
	return true
}

//根据新闻项的id和固定URL来取得新闻内容，并保存为.html文件
func GetStoryContent(it Item, path string) {
	var content []byte
	var filepath string
	defer func() {
		v := recover()
		if v != nil {
			fmt.Println(v)
		}
	}()
	body, err := http.Get(patch_shareUrl + strconv.Itoa(it.Id))
	if err != nil {
		panic(err)
	}
	filepath = DOWNLOAD + path + "/" + it.Title + it.Theme_name + ".html"

	content, err = ioutil.ReadAll(body.Body)
	fmt.Println(filepath)
	if isExist(filepath) {
		fmt.Println("file already exist")
	} else {
		file, err := os.Create(filepath)
		file.Write(content)
		if err != nil {
			panic(err)
		}
	}

}

func GetPageBody(urlpath string, date string) {
	urlpath = urlpath
	if urlpath == "" {
		fmt.Println("url path are null ")
		return
	}
	defer func() {
		v := recover()
		if v != nil {
			fmt.Println(v)
		}
	}()

	Creator(date)

	resp, err := http.Get(urlpath)
	defer resp.Body.Close()

	var content *TodayNews
	var imagePath string

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content = ParseContent([]byte(body))

	for _, item := range content.Story {
		resp, err := http.Get(item.Images[0])
		imgBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		GetStoryContent(item, content.Date)
		temp := strings.Split(item.Images[0], "/")

		imagePath = DOWNLOAD + date + "/" + temp[len(temp)-1]
		fmt.Println(imagePath)
		if isExist(imagePath) {
			fmt.Println("file already exist")
		} else {
			file, err := os.Create(imagePath)
			if err != nil {
				panic(err)
			}
			file.Write(imgBody)
		}

	}
	for _, item := range content.TopStory {
		fmt.Println(item.Image)
		resp, err := http.Get(item.Image)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Header)
	}
	return
}

func MonthToNumber(month string) string {
	switch month {
	case "January":
		return "01"
	case "February":
		return "02"
	case "March":
		return "03"
	case "April":
		return "04"
	case "May":
		return "05"
	case "June":
		return "06"
	case "July":
		return "07"
	case "August":
		return "08"
	case "September":
		return "09"
	case "October":
		return "10"
	case "November":
		return "11"
	case "December":
		return "12"
	}
	return "01"
}

func DayToNumber(day int) string {
	switch day {
	case 1:
		return "01"
	case 2:
		return "02"
	case 3:
		return "03"
	case 4:
		return "04"
	case 5:
		return "05"
	case 6:
		return "06"
	case 7:
		return "07"
	case 8:
		return "08"
	case 9:
		return "09"
	default:
		return strconv.Itoa(day)
	}
	return "01"
}

func GetDateString(t time.Time) string {
	/*const layout = "2, 2006 at 3:04pm (MST)"
	var ti time.Time
	ti = time.Now()
	//ti.Year()
	ti.Format(layout)
	fmt.Println(ti.Year())
	fmt.Println(ti.Day())
	fmt.Println(ti.Date())
	*/
	temp := strconv.Itoa(t.Year()) + MonthToNumber(t.Month().String()) + DayToNumber(t.Day())
	return temp
}

func GetBeforeNews(start time.Time, end time.Time) {
	tCreate := time.Date(2013, time.May, 21, 12, 0, 0, 0, time.UTC)

	if start.Before(tCreate) {
		fmt.Println("start time nis before Zhihu's Dayly")
		return
	}

	if time.Now().Before(end) {
		fmt.Println("end time is faraway today")
		return
	}

	for t := start; t.Before(end); t = t.AddDate(0, 0, 1) {
		date := GetDateString(t)
		url := ZhihuBeforeNewsUrl + date
		GetPageBody(url, date)
	}
}

func SpyMain() {
	//GetPageBody("http://news-at.zhihu.com/api/3/news/latest")
	begin := time.Date(2014, time.December, 21, 18, 0, 0, 0, time.UTC)

	end := time.Date(2014, time.December, 25, 18, 0, 0, 0, time.UTC)
	GetBeforeNews(begin, end)
}
