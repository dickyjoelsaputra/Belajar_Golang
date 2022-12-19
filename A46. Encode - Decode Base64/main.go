package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)


	//  Encode & Decode Data URL
	datas := "https://kalipare.com/"

	var encodedStrings = base64.URLEncoding.EncodeToString([]byte(datas))
	fmt.Println(encodedStrings)

	var decodedBytes, _ = base64.URLEncoding.DecodeString(encodedStrings)
	var decodedStrings = string(decodedBytes)
	fmt.Println(decodedStrings)
}
