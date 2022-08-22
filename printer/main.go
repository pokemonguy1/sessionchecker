package main

import (
	"encoding/json"
	"fmt"
	"github.com/hennedo/escpos"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type Configuration struct {
	Host string `json:"host"`
}

func main() {

	jsonFile, err := os.Open("configure.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	config := Configuration{}

	json.Unmarshal(byteValue, &config)

	socket, err := net.Dial("tcp", config.Host)
	if err != nil {
		println(err.Error())
	}
	defer socket.Close()

	p := escpos.New(socket)

	//p.LineFeed()
	//p.QRCode("https://github.com/hennedo/escpos", true, 10, escpos.QRCodeErrorCorrectionLevelH)

	imgFile, err := os.Open("test2.png")
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot read file:", err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}

	p.PrintImage(img)
	p.Cut()

}
