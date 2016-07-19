package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

var (
	host    = "http://www.evcardchina.com"
	session *grequests.Session
)

type Car struct {
	shopSeq      string
	drivingRange int
}

type EvcardUser struct {
	Username string
	Password string
}

func GetAllShopInfoList() []map[string]interface{} {
	resp, err := session.Post("http://www.evcardchina.cn/api/proxy/getShopInfoList", &grequests.RequestOptions{JSON: map[string]string{"token": "", "shopSeq": "-1"}})
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	//	fmt.Println(resp.String())
	respData := []byte(resp.String())
	//fmt.Printf("%T", respData)

	var dat []map[string]interface{}
	if err := json.Unmarshal(respData, &dat); err != nil {
		panic(err)
	}
	return dat
	/*
		for _, vs := range dat {
			for key, val := range vs {
				fmt.Println(key, val)
			}
		}
	*/
}

func (evcard *EvcardUser) login() {
	resp, err := session.Get("http://www.evcardchina.cn", nil)
	if err != nil {
		log.Println(err)
	}
	if resp.Ok != true {
		log.Println("Request did not return OK")
	}

	_, posterr := session.Post("http://www.evcardchina.com/login", &grequests.RequestOptions{JSON: map[string]string{"login": "13162502908", "password": "zkZK1234"}})
	if posterr != nil {
		log.Fatal("Unable to make request:", posterr)
	}
	//	fmt.Println(postresp.String())
}

func (evcard *EvcardUser) getCar(shopSeq int) <-chan string {
	fmt.Println("test")
	vehicalch := make(chan string)
	go func(vehical chan string) {
		defer close(vehicalch)
		resp, err := session.Post("http://www.evcardchina.com/api/proxy/getVehicleInfoList", &grequests.RequestOptions{JSON: map[string]int{"shopSeq": shopSeq, "canRent": 1}})
		if err != nil {
			log.Fatal("Unable to make request:", err)
		}
		fmt.Println(resp.Error)
		vehicalch <- resp.String()
	}(vehicalch)
	return vehicalch
}

func main() {
	session = grequests.NewSession(nil)
	evcard := &EvcardUser{Username: "13162502908", Password: "zkZK1234"}
	evcard.login()
	// You can modify the request by passing an optional RequestOptions struct
	GetAllShopInfoList()
	ret := <-evcard.getCar(1)
	fmt.Println(ret)
}
