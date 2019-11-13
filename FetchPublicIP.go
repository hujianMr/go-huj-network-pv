package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Unknwon/goconfig/goconfig-master"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

var IP string

func main() {
	go asyncFetchIp()
	select {}
}

var key string

func asyncFetchIp() {
	cfg, _ := goconfig.LoadConfigFile("config.ini")
	url, _ := cfg.GetValue("server", "url")
	key, _ = cfg.GetValue("server", "key")
	for {
		urlPath := url + "urlPath"
		res, err := http.Get(urlPath)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		ipUrl := string(body)

		publicIp := PublicIp(ipUrl)
		fmt.Println("获取公网ip 地址 = " + publicIp)
		pushUrl := url + "pushIp?ip=" + publicIp + "&key=" + key
		_, _ = http.Get(pushUrl)
		time.Sleep(120 * time.Second)
	}
}

var decodeIp = mahonia.NewDecoder("utf-8")

func PublicIp(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("status code error: %d %s", res.StatusCode, res.Status)
		return ""
	}
	utfBody := decodeIp.NewReader(res.Body)
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Println(err)
		return ""
	}
	ip := doc.Find("title").Text()

	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	return reg.FindString(ip)
}
