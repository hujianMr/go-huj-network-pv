package main

import (
	"encoding/json"
	"fmt"
	"github.com/Unknwon/goconfig/goconfig-master"
	"log"
	"net/http"
)

var IpMap = make(map[string]string)

func main() {
	http.HandleFunc("/urlPath", urlPath)
	http.HandleFunc("/publicIp", publicIp)
	http.HandleFunc("/pushIp", pushIp)
	http.HandleFunc("/listIp", listIp)
	cfg, _ := goconfig.LoadConfigFile("config.ini")
	addr, _ := cfg.GetValue("host", "addr")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func pushIp(w http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	ip := getFormValue("ip", req)
	key := getFormValue("key", req)
	fmt.Println("ip=" + ip + "  key=" + key)
	IpMap[key] = ip
	_, _ = fmt.Fprintf(w, "ok")
}

func publicIp(w http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	key := getFormValue("key", req)
	_, _ = fmt.Fprintf(w, IpMap[key])
}

func listIp(w http.ResponseWriter, req *http.Request) {
	bytes, _ := json.Marshal(IpMap)
	_, _ = fmt.Fprintf(w, string(bytes))
}

func urlPath(w http.ResponseWriter, req *http.Request) {
	cfg, _ := goconfig.LoadConfigFile("config.ini")
	url, _ := cfg.GetValue("publicIP", "url")
	_, _ = fmt.Fprintf(w, url)
}

func getFormValue(p string, req *http.Request) string {
	v, f := req.Form[p]
	if !f {
		log.Println("..接收参数异常:" + p + "  ")
		return ""
	}
	return v[0]
}
