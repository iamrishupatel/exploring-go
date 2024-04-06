package main

import (
	"fmt"
	"sync"

	"example.com/crypto-masters/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "LTC", "XRP", "BCH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(curr string) {
			getCurrencyData(curr)
			wg.Done()
		}(currency)
	}
	wg.Wait()

}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %s is %.2f\n", rate.Currency, rate.Price)
	}
}
