package main

import "testing"

func TestPrintStruct(t *testing.T) {
	mapToTest := map[string]interface{}{
		"id":    1,
		"name":  "nameless",
		"value": 123,
	}

	var in inter

	_ = PrintStruct(in, mapToTest)
}
