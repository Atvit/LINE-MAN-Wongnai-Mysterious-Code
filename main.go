package main

import (
	"fmt"
	"lmwn/mysterious"
	"log"
)

func main() {
	m := mysterious.NewMysterious()

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, err := m.Decode(secret)
	if err != nil {
		log.Printf("error can not decode base64 string (%s) %s", secret, err.Error())
	}

	result := m.Reverse(sd)
	fmt.Println(result)
}
