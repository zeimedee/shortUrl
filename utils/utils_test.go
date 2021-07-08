package utils

import "testing"

func TestShortener(t *testing.T) {
	shorten := Shortener()

	if len(shorten) != 6 {
		t.Errorf("length of shortened url should be 6")
	}
}
