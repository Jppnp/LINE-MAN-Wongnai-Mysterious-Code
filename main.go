package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}
	whatIsIt = reverse(string(sd))
	whatIsIt = strings.Replace(whatIsIt, ":", " ", -1)
	fmt.Println(whatIsIt)
}
