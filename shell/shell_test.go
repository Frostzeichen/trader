package shell

import (
	"testing"
	"reflect"
	"encoding/json"
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

		response1 := GETPrices("BTC", true)
		response2 := GETPrices("USDC", true)

		json.Unmarshal([]byte(string(response1)), &f)
		json.Unmarshal([]byte(string(response2)), &g)

		got := f.Price
		want := g.Price

		if got == want {
			t.Errorf("Seems like both values got the same value of %s.", got)
		}
	})
}

func assertStatus(t *testing.T, got reflect.Type, want reflect.Type) {
	if want != got {
		t.Errorf("got %s expected %s", got, want)
	}
}
