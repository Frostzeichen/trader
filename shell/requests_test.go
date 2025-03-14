package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRunGETPrices(t *testing.T) {
	t.Run("GETPrices returns price", func(t *testing.T) {
		response := GETPrices("BTC", false)
		var f PricesJSONResponse
		json.Unmarshal([]byte(string(response)), &f)
		got := reflect.TypeOf(f.Price)
		want := reflect.TypeOf("1.001")

		assertStatus(t, got, want)
	})

	t.Run("parameter works", func(t *testing.T) {
		var f PricesJSONResponse
		var g PricesJSONResponse

		response1 := GETPrices("BTC", false)
		response2 := GETPrices("USDC", false)

		json.Unmarshal([]byte(string(response1)), &f)
		json.Unmarshal([]byte(string(response2)), &g)

		got := f.Price
		want := g.Price

		if got == want {
			t.Errorf("Seems like both values got the same value of %s.", got)
		}
	})
}

func TestRunPOSTTradeQuote(t *testing.T) {
	t.Run("trade request should return JSON", func(t *testing.T) {
		response := POSTTradeQuote("BTC", "PHP", false)
		got := reflect.TypeOf(response)
		want := reflect.TypeOf([]byte(string(`"status": 0`)))

		assertStatus(t, got, want)
	})

//	t.Run("response should not be hardcoded", func(t *testing.T) {
//		var f TradeQuoteJSONResponse
//		var g TradeQuoteJSONResponse
//
//		response1 := POSTTradeQuote("BTC", "PHP", false)
//		response2 := POSTTradeQuote("USDC", "PHP", false)
//
//		json.Unmarshal([]byte(string(response1)), &f)
//		json.Unmarshal([]byte(string(response2)), &g)
//
//		got := f.Data.Price
//		want := g.Data.Price
//
//		if got == want {
//			t.Errorf("Seems like both values got the same value of %s.", got)
//		}
//	})
}

func assertStatus(t *testing.T, got reflect.Type, want reflect.Type) {
	if got != want {
		t.Errorf("got %s expected %s", got, want)
	}
}
