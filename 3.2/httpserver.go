package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>cloud native</h1>"))
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	//get version
	os.Setenv("version", "v9")
	version := os.Getenv("version")
	w.Header().Add("version", version)
	//get client ip,如果有LB则取得的LB的地址或者是proxy代理地址
	//X-REAL-IP   X-FORWARD-FOR真实用户ip
	clientIP := getrealip(r)
	httpCode := http.StatusOK
	log.Println("", clientIP, httpCode)
}

func getrealip(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")

	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "ok")

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("start server failed, %s \n", err.Error())
	}
}
