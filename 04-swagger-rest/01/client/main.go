package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/lizrice/secure-connections/utils"
)

type Order struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

var order Order

func getFlags() {
	flag.StringVar(&order.CandyType, "k", "", "type of candy")
	flag.IntVar(&order.CandyCount, "c", 0, "type of candy")
	flag.IntVar(&order.Money, "m", 0, "type of candy")
	flag.Parse()
}

func main() {
	getFlags()
	client := getClient()
	bodyJson, _ := json.Marshal(order)
	bodyPost := bytes.NewReader(bodyJson)
	resp, err := client.Post("https://candy.tld:8080/buy_candy", "application/json", bodyPost)
	must(err)
	if resp.StatusCode == 201 {
		var success InlineResponse201
		json.NewDecoder(resp.Body).Decode(&success)
		fmt.Println(resp.StatusCode, success.Change, success.Thanks)
	} else if resp.StatusCode == 400 {
		var fail InlineResponse400
		json.NewDecoder(resp.Body).Decode(&fail)
		fmt.Println(resp.StatusCode, fail.Error_)
	}
}

func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("minica.pem")
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		RootCAs:               cp,
		GetClientCertificate:  utils.ClientCertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func must(err error) {
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
}
