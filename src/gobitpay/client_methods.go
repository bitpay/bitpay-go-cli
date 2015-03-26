package main

import (
	"encoding/json"
	bp "github.com/bitpay/bitpay-go/client"
	ku "github.com/bitpay/bitpay-go/key_utils"
	"io/ioutil"
	"os"
	"strconv"
)

func getInvoice(invId string) (price float64, currency string, err error) {
	rootPath := os.ExpandEnv("$HOME")
	client := clientFromFiles(rootPath)
	inv, err := client.GetInvoice(invId)
	if err == nil {
		price = inv.Price
		currency = inv.Currency
		return price, currency, err
	}
	return price, currency, err
}

func createInvoice(price string, currency string) (invId string, err error) {
	rootPath := os.ExpandEnv("$HOME")
	client := clientFromFiles(rootPath)
	parsedPrice, _ := strconv.ParseFloat(price, 64)
	tok, _ := ioutil.ReadFile(rootPath + "/.bp/token.json")
	var token bp.Token
	json.Unmarshal(tok, &token)
	client.Token = token
	inv, err := client.CreateInvoice(parsedPrice, currency)
	invId = inv.Id
	return invId, err
}

func pairClient(code string) (token string, err error) {
	rootPath := os.ExpandEnv("$HOME")
	client := clientFromFiles(rootPath)
	var toke bp.Token
	toke, err = client.PairWithCode(code)
	var tokenByte []byte
	tokenByte, err = json.Marshal(toke)
	ioutil.WriteFile(rootPath+"/.bp/token.json", tokenByte, 0644)
	token = string(tokenByte)
	return token, err
}

func clientFromFiles(rootPath string) bp.Client {
	pm, _ := ioutil.ReadFile(rootPath + "/.bp/bitpay.pem")
	uri, _ := ioutil.ReadFile(rootPath + "/.bp/uri.txt")
	notsecure, _ := ioutil.ReadFile(rootPath + "/.bp/insecure.txt")
	var insecure bool
	if string(notsecure) == "true" {
		insecure = true
	} else {
		insecure = false
	}
	client := bp.Client{ApiUri: string(uri), Insecure: insecure, Pem: string(pm)}
	return client
}

func newClient(uri string, insecure bool) {
	pm := ku.GeneratePem()
	notSecure := "false"
	if insecure {
		notSecure = "true"
	}
	rootPath := os.ExpandEnv("$HOME")
	os.Mkdir(rootPath+"/.bp", 0755)
	pmByte := []byte(pm)
	uriByte := []byte(uri)
	secByte := []byte(notSecure)
	ioutil.WriteFile(rootPath+"/.bp/bitpay.pem", pmByte, 0644)
	ioutil.WriteFile(rootPath+"/.bp/uri.txt", uriByte, 0644)
	ioutil.WriteFile(rootPath+"/.bp/insecure.txt", secByte, 0644)
}
