package main

import (
	"testing"
)

func TestGeneratorQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-2368")

	if result == nil {
		t.Errorf("GeneratedQRCode is nil")
	}
	if len(result) == 0 {
		t.Errorf("GeneratedQRCode has no data.")
	}
}
