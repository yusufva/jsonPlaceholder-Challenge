package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RequestBody struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		apiUrl := "http://jsonplaceholder.typicode.com/posts"

		data := RequestBody{
			Water: rand.Intn(16),
			Wind:  rand.Intn(16),
		}

		bs, err := json.Marshal(data)

		if err != nil {
			log.Panicf("error converting struct to json => %s \n", err.Error())
		}

		request, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(bs))

		if err != nil {
			log.Panicf("error while defining request instance => %s\n", err.Error())
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		response, err := client.Do(request)

		if err != nil {
			log.Panicf("error while sending request instance => %s\n", err.Error())
		}

		defer response.Body.Close()

		responseBody, err := ioutil.ReadAll(response.Body)

		fmt.Println(string(responseBody))

		switch air := data.Water; {
		case air < 6:
			fmt.Println("status Water: Aman")

		case air >= 6 && air <= 8:
			fmt.Println("status Water: Siaga")

		case air > 8:
			fmt.Println("status Water: Bahaya")
		}

		switch udara := data.Wind; {
		case udara < 7:
			fmt.Println("status Water: Aman")

		case udara >= 7 && udara <= 15:
			fmt.Println("status Water: Siaga")

		case udara > 15:
			fmt.Println("status Water: Bahaya")
		}

		time.Sleep(time.Second * 15)
	}

}
