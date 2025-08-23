package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Connecting Crypto API to get details of all supported coins")
	apiKey, found := os.LookupEnv("CRYPTO_KEY")
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/list?x_cg_demo_api_key=%s", apiKey)
	if found {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed due to %s", err)
		}
		defer resp.Body.Close()
		fmt.Println("Result is",resp)

	} else {
		fmt.Println("CRYTPO_KEY not set")
	}
}
