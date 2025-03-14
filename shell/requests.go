package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type PricesJSONResponse struct {
	Mins     int    `json:"mins"`
	Price    string `json:"price"`
	Currency string `json:"currency"`
}

type ExchangeInformationJSONResponse struct {
	Timezone        string        `json:"timezone"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	ServerTime      int           `json:"serverTime"`
	Symbols         []struct {    // TODO: Refactor to reduce indentation.
		Symbol              string     `json:"symbol"`
		Status              string     `json:"status"`
		BaseAsset           string     `json:"baseAsset"`
		BaseAssetPrecision  int        `json:"baseAssetPrecision"`
		QuoteAsset          string     `json:"quoteAsset"`
		QuoteAssetPrecision int        `json:"quoteAssetPrecision"`
		OrderTypes          []string   `json:"orderTypes"`
		Filters             []struct { // TODO: Refactor to reduce indentation.
			FilterType        string `json:"filterType"`
			MinPrice          string `json:"minPrice,omitempty"`
			MaxPrice          string `json:"maxPrice,omitempty"`
			TickSize          string `json:"tickSize,omitempty"`
			MinQty            string `json:"minQty,omitempty"`
			MaxQty            string `json:"maxQty,omitempty"`
			StepSize          string `json:"stepSize,omitempty"`
			MinNotional       string `json:"minNotional,omitempty"`
			MaxNotional       string `json:"maxNotional,omitempty"`
			PriceUp           string `json:"priceUp,omitempty"`
			PriceDown         string `json:"priceDown,omitempty"`
			BidMultiplierUp   string `json:"bidMultiplierUp,omitempty"`
			BidMultiplierDown string `json:"bidMultiplierDown,omitempty"`
			AskMultiplierUp   string `json:"askMultiplierUp,omitempty"`
			AskMultiplierDown string `json:"askMultiplierDown,omitempty"`
			MultiplierUp      string `json:"multiplierUp,omitempty"`
			MultiplierDown    string `json:"multiplierDown,omitempty"`
			MaxNumOrders      int    `json:"maxNumOrders,omitempty"`
			MaxNumAlgoOrders  int    `json:"maxNumAlgoOrders,omitempty"`
		} `json:"filters"`
	} `json:"symbols"`
}

type TradeQuoteJSONResponse struct {
	Status int                  `json:"status"`
	Error  string               `json:"error"`
	Data   TradeQuoteDataStruct `json:"data"`
}

type TradeQuoteDataStruct struct {
	QuoteID        string `json:"quoteId"`
	SourceCurrency string `json:"sourceCurrency"`
	TargetCurrency string `json:"targetCurrency"`
	SourceAmount   string `json:"sourceAmount"`
	TargetAmount   string `json:"targetAmount"`
	Price          string `json:"price"`
	Expiry         string `json:"expiry"`
}

// Gets the latest price and returns raw JSON output.
func GETPrices(currency string, printJSON bool) []byte {
	cmd := exec.Command("./check_prices.sh", currency)
	cmdOutput, _ := cmd.Output()

	var f PricesJSONResponse
	json.Unmarshal([]byte(string(cmdOutput)), &f)
	f.Currency = currency

	if printJSON {
		fmt.Println(f)
	}

	cmdOutput, _ = json.Marshal(f)

	return cmdOutput
}

func POSTTradeQuote(targetCurrency string, sourceCurrency string, printJSON bool) []byte {
	cmd := exec.Command("./fetch_quote.sh", targetCurrency, sourceCurrency)
	cmdOutput, _ := cmd.Output()

	if printJSON {
		var f TradeQuoteJSONResponse
		json.Unmarshal([]byte(string(cmdOutput)), &f)
		fmt.Println(f)

		// Normally, f == {0 {    }} if there is an IP-related error. Hence f.Status will not be 200.
		if f.Status != 200 {
			fmt.Println("There could be an IP-related error. Please check. File name is fetch_quote.sh.")
		}
	}

	return cmdOutput
}

func GETExchangeInformation(currency string, printJSON bool) []byte {
	cmd := exec.Command("./check_prices.sh", currency)
	cmdOutput, _ := cmd.Output()
	jsonOutput, _ := json.MarshalIndent(string(cmdOutput), "", "\t")

	if printJSON {
		var f ExchangeInformationJSONResponse
		json.Unmarshal([]byte(string(cmdOutput)), &f)
		fmt.Println(f.Symbols[0].Symbol)
	}

	return jsonOutput
}
