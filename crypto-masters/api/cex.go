package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"example.com/crypto-masters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {

	if len(currency) != 3 {
		return nil, fmt.Errorf("minimum 3 characters required")
	}

	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch rate for %s", currency)
	} else {

		bytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		var cexRate CEXRateResponse
		err = json.Unmarshal(bytes, &cexRate)
		if err != nil {
			return nil, err
		}

		floatRate, _ := strconv.ParseFloat(cexRate.Last, 64)

		rate := datatypes.Rate{Currency: currency, Price: floatRate}
		return &rate, nil

	}
}
