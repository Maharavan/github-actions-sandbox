package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func main() {
	fmt.Println("Connecting Crypto API to get details of all supported coins")
	apiKey, found := os.LookupEnv("CRYPTO_KEY")
	fmt.Println(apiKey)
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/list?x_cg_demo_api_key=%s", apiKey)
	if found {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed due to %s", err)
			panic("error")
		}
		defer resp.Body.Close()
		fmt.Println("Result is", resp)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed due to %s", err)
			panic("error")
		}

		var c []Coin
		err = json.Unmarshal(body, &c)

		if err != nil {
			fmt.Printf("Failed due to %s", err)
			panic("error")
		}

		for _, val := range c {
			fmt.Println(val)
		}
	} else {
		fmt.Println("CRYTPO_KEY not set")
	}
}
