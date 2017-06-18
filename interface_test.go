package main

import (
	"testing"
)

func TestFunctionCall(t *testing.T) {

	c := []struct{
		in string
		function string
		params string
	}{
		{"version", "version", ""},
		{"getStatus", "getStatus", ""},
		{"prepLoadout steamid", "prepLoadout", "steamid"},
	}

	for i := 0; i < len(c); i++  {

		function, param := parseFunctionCall(c[i].in)

		if function != c[i].function || param != c[i].params {
			t.Error("Function call test failed")
		}

	}
}