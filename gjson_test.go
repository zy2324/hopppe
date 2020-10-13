package main

import (
	"github.com/tidwall/gjson"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())
}
