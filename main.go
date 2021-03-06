package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	url     = flag.String("url", "http://localhost", "URL Address")
	file    = flag.String("file", "test.json", "post")
	contype = flag.String("contype", "application/json", "Content-Type")
	showver = flag.Bool("v", false, "Show version")
)

func main() {

	flag.Parse()
	fmt.Println("Ver:0.1")
	if *showver {

		return
	}
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 4000 * time.Millisecond,
		}).Dial,
		TLSHandshakeTimeout: 4000 * time.Millisecond,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	var nc = &http.Client{Timeout: time.Millisecond * 4000,
		Transport: netTransport}
	fmt.Printf("Import file:\t%s\n", *file)
	filetext, err := ioutil.ReadFile(*file)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("URL:\t\t%s\n", *url)
	req, err := http.NewRequest("POST", *url, strings.NewReader(string(filetext)))

	req.Header.Add("Content-Type", *contype)
	fmt.Printf("Content-Type:\t%s \n", *contype)
	resp, err := nc.Do(req)

	if err == nil {
		fmt.Printf("File Sent %v\n", resp.StatusCode)

		defer resp.Body.Close()

	} else {
		fmt.Println(err)
	}

}
