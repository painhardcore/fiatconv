package conv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiResponse struct {
	Rates map[string]float64
	Error string
}

func Currency(amount float64, from, to string) (float64, error) {
	req := fmt.Sprintf("https://api.exchangeratesapi.io/latest?base=%s&symbols=%s", from, to)
	resp, err := http.Get(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var r ApiResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return 0, err
	}

	if r.Error != "" {
		return 0, errors.New(r.Error)
	}

	return amount * r.Rates[to], nil
}
