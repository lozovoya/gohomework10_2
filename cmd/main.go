package main

import (
	"context"
	"github.com/lozovoya/gohomework10_2/pkg/qr"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const line = "www.yandex.ru"
const filename = "yandex.png"

func main() {

	qrURL, ok := os.LookupEnv("QRURL")
	if !ok {
		log.Println("No URL available")
		os.Exit(1)
	}
	qrVersion, ok := os.LookupEnv("QRVERSION")
	if !ok {
		log.Println("No version available")
		os.Exit(1)
	}
	qrTimeout, ok := os.LookupEnv("QRTIMEOUT")
	if !ok {
		log.Println("No timeout available")
		os.Exit(1)
	}
	timeout, err := strconv.Atoi(qrTimeout)
	if err != nil {
		log.Println("wrong timeout format")
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	svc := qr.NewService(
		qrURL,
		qrVersion,
		ctx,
		&http.Client{},
	)
	data, err := svc.Encode(line)

	file, err := os.Create("test.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	err = ioutil.WriteFile("test.png", data, 0777)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
