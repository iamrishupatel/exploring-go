package api_test

import (
	"testing"

	"example.com/crypto-masters/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("No Error was returned when an empty string was passed")
	}
}
