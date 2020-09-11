package main

import (
	"fmt"
	"proxy/balance"
	"testing"
)

func TestRandomBalance(t *testing.T) {
	rb := &balance.RandomBalance{}
	rb.Add("127.0.0.1:2003")
	rb.Add("127.0.0.1:2004")
	rb.Add("127.0.0.1:2005")
	rb.Add("127.0.0.1:2006")
	rb.Add("127.0.0.1:2007")

	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
}