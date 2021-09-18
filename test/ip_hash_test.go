package main

import (
	"fmt"
	"proxy/balance"
	"testing"
)

func TestIpHashBalance(t *testing.T) {
	rb := balance.IPHshBalance{}
	rb.Add("127.0.0.1:2003")
	rb.Add("127.0.0.1:2004")
	rb.Add("127.0.0.1:2005")
	rb.Add("127.0.0.1:2006")
	rb.Add("127.0.0.1:2007")

	fmt.Println(rb.Get("127.0.0.1"))
	fmt.Println(rb.Get("192.168.1.12"))
	fmt.Println(rb.Get("127.0.0.1"))
	fmt.Println(rb.Get("192.168.1.12"))

}
