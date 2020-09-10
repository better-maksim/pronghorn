package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	str := "123:123"
	base := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(base)
}