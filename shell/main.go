package main

import (
	"fmt"
	"encoding/json"
	"math/big"
	"strconv"
	"os"
	"time"
)

func executeTrade(coins []string) {
	for true {
	c := make(chan []byte, 2)

	for _, coin := range coins {
		go func() {
			c <- GETPrices(coin, false)
		}()
	}

	r1, r2, r3, r4, r5 := <-c, <-c, <-c, <-c, <-c
	
	var f PricesJSONResponse
	var g PricesJSONResponse
	var h PricesJSONResponse
	var i PricesJSONResponse
	var j PricesJSONResponse

	json.Unmarshal([]byte(string(r1)), &f)
	json.Unmarshal([]byte(string(r2)), &g)
	json.Unmarshal([]byte(string(r3)), &h)
	json.Unmarshal([]byte(string(r4)), &i)
	json.Unmarshal([]byte(string(r5)), &j)

	fmt.Printf("%s: %s\n%s: %s\n%s: %s\n%s: %s\n%s: %s", f.Currency, f.Price, g.Currency, g.Price, h.Currency, h.Price, i.Currency, i.Price, j.Currency, j.Price)

	btcBase := 5600312.962363588703558076
	usdcBase := 58.07689184275714527
	xrpBase := 138.472816354053522092
	dogeBase := 19.658
	ethBase := 200796.695706111655084865

	fPrice, _ := strconv.ParseFloat(f.Price, 64)
        gPrice, _ := strconv.ParseFloat(g.Price, 64)
        hPrice, _ := strconv.ParseFloat(h.Price, 64)
        iPrice, _ := strconv.ParseFloat(i.Price, 64)
        jPrice, _ := strconv.ParseFloat(j.Price, 64)

//	fIsHigher := big.NewFloat(fPrice).Cmp(big.NewFloat(btcBase))
//	gIsHigher := big.NewFloat(gPrice).Cmp(big.NewFloat(usdcBase))
//	hIsHigher := big.NewFloat(hPrice).Cmp(big.NewFloat(xrpBase))
//	iIsHigher := big.NewFloat(iPrice).Cmp(big.NewFloat(dogeBase))
//	jIsHigher := big.NewFloat(jPrice).Cmp(big.NewFloat(ethBase))

	fmt.Println();
	fmt.Printf("%s: %d\n",f.Currency, big.NewFloat(fPrice).Cmp(big.NewFloat(btcBase)))
	fmt.Printf("%s: %d\n",g.Currency, big.NewFloat(gPrice).Cmp(big.NewFloat(usdcBase)))
	fmt.Printf("%s: %d\n",h.Currency, big.NewFloat(hPrice).Cmp(big.NewFloat(xrpBase)))
	fmt.Printf("%s: %d\n",i.Currency, big.NewFloat(iPrice).Cmp(big.NewFloat(dogeBase)))
	fmt.Printf("%s: %d\n",j.Currency, big.NewFloat(jPrice).Cmp(big.NewFloat(ethBase)))

	currentDate := time.Now()

	formattedDate := currentDate.Format("2006-01-02")

	filename := fmt.Sprintf("%s.csv", formattedDate)

	fi, err := os.OpenFile("dataset/" + filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = fi.WriteString(fmt.Sprintf("%s,%s,%s\n%s,%s,%s\n%s,%s,%s\n%s,%s,%s\n%s,%s,%s\n", f.Currency, f.Price, currentDate.Format("2006-01-02T15:04:05"), g.Currency, g.Price, currentDate.Format("2006-01-02T15:04:05"), h.Currency, h.Price, currentDate.Format("2006-01-02T15:04:05"), i.Currency, i.Price, currentDate.Format("2006-01-02T15:04:05"), j.Currency, j.Price, currentDate.Format("2006-01-02T15:04:05")))
	if err != nil {
		panic(err)
	}
	
	time.Sleep(30 * time.Minute)
	}
}

func main() {
	coins := []string{"BTC", "USDC", "XRP", "DOGE", "ETH", "CBX"}
	executeTrade(coins);
}
