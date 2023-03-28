// https://stackoverflow.com/questions/15334220/encode-decode-base64

package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	var message = "stop"
	str := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(str)

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Printf("%q\n", data)
}
