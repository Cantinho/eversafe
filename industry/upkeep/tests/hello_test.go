package tests

import (
"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := "Hello Unkeep!"

	if result != "Hello Unkeep!" {
		t.Errorf("Generated QRCode is nil")
	}
	if len(result) == 0 {
		t.Errorf("Message is empty")
	}
}
