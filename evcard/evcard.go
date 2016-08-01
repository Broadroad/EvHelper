package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"log"
	//	"net/url"
	_ "strings"
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
	Token    string
	Vin      string
	AuthId   string
	ShopSeq  int
}

type ShopRealInfo struct {
	SerialNum   int        `json:",omitempty"`
	Status      int        `json:",omitempty"`
	Message     string     `json:",omitempty"`
	ServiceName string     `json:",omitempty"`
	Token       string     `json:",omitempty"`
	DataList    []dataList `json:",omitempty"`
	ListSize    int        `json:",omitempty"`
}

type dataList struct {
	ShopSeq        int `json:",omitempty"`
	StackCnt       int `json:",omitempty"`
	AllowStackCnt  int `json:",omitempty"`
	AllowVehileCnt int `json:",omitempyt"`
	CanParkNum     int `json:",omitempty"`
	ParkCnt        int `json:",omitempty"`
}

type OrderResp struct {
	SerialNum   int    `json:",omitempty"`
	Status      int    `json:",omitempty"`
	Message     string `json:",omitempty"`
	ServiceName string `json:",omitempty"`
	Token       string `json:",omitempty"`
	ReturnCode  int    `json:",omitempty"`
	OrderSeq    int    `json:",omitempty"`
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err

	}
	return buf.Bytes(), nil

}

func GetAllShopInfoList() []map[string]interface{} {
	resp, err := session.Post("http://www.evcardchina.cn/api/proxy/getShopInfoList", &grequests.RequestOptions{JSON: map[string]string{"token": "", "shopSeq": "-1"}, RedirectLocationTrusted: false})
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	respData := []byte(resp.String())

	var dat []map[string]interface{}
	if err := json.Unmarshal(respData, &dat); err != nil {
		panic(err)
	}
	return dat
}

func (evcard *EvcardUser) login() {
	resp, err := session.Get("http://www.evcardchina.com", nil)
	if err != nil {
		log.Println(err)
	}
	if resp.Ok != true {
		log.Println("Request did not return OK")
	}

	postresp, err := session.Post("http://58.32.252.60:8080/isv2/evcardapp?service=login", &grequests.RequestOptions{JSON: map[string]string{"loginName": "13162502908", "password": "zkZK1234"}, Headers: map[string]string{"Origin": "http://www.evcardchina.com"}, UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"})
	if err != nil {
		log.Fatal("Unable to make request:", err)
	}
	respReader := []byte(postresp.String())
	var dat map[string]interface{}

	if err := json.Unmarshal(respReader, &dat); err != nil {
		panic(err)
	}
	evcard.Token = dat["token"].(string)
}

func (evcard *EvcardUser) getCarByShopSeq() <-chan string {
	vehicalch := make(chan string)
	go func(vehical chan string) {
		defer close(vehicalch)
		for {
			resp, err := session.Post("http://www.evcardchina.com/api/proxy/getVehicleInfoList", &grequests.RequestOptions{JSON: map[string]int{"shopSeq": evcard.ShopSeq, "canRent": 1}})
			if err != nil {
				log.Fatal("Unable to make request:", err)
			}
			fmt.Println(resp.String())
			respData := []byte(resp.String())
			var dat []map[string]interface{}
			if err = json.Unmarshal(respData, &dat); err != nil {
				panic(err)
			}
			if dat == nil {
				continue
			} else {
				vehicalch <- getMaxDrivingVin(dat)
				break
			}
		}
	}(vehicalch)
	return vehicalch
}

func getMaxDrivingVin(vinlist []map[string]interface{}) (vin string) {
	maxDriving := 0.0
	for _, vs := range vinlist {
		if vs["drivingRange"].(float64) > maxDriving {
			maxDriving = vs["drivingRange"].(float64)
			vin = vs["vin"].(string)
		}
	}
	return vin
}

func (evcard *EvcardUser) getLastShopList() (shopSeq int) {
	return
}

func (evcard *EvcardUser) getShopRealInfo() {
	resp, err := session.Post("http://58.32.252.61:8082/isv2/evcardapp?service=getShopRealInfo", &grequests.RequestOptions{JSON: map[string]interface{}{"token": evcard.Token, "shopSeq": -1}})
	respData := []byte(resp.String())
	var dat ShopRealInfo
	if err = json.Unmarshal(respData, &dat); err != nil {
		panic(err)
	}
}

func (evcard *EvcardUser) orderCar() {
	resp, _ := session.Post("http://58.32.252.60:8080/isv2/evcardapp?service=orderVehicle", &grequests.RequestOptions{JSON: map[string]interface{}{"planpickupstoreseq": evcard.ShopSeq, "vin": evcard.Vin, "authId": evcard.AuthId, "token": evcard.Token}, UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"})
	fmt.Println(resp.String())
}

func main() {
	session = grequests.NewSession(nil)
	evcard := &EvcardUser{Username: "", Password: "", ShopSeq: 1}
	evcard.login()
	// You can modify the request by passing an optional RequestOptions struct
	GetAllShopInfoList()
	evcard.Vin = <-evcard.getCarByShopSeq()
	evcard.getShopRealInfo()
	evcard.orderCar()
}
