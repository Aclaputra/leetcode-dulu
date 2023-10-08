package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	convertedIP, err := convertIP(address)
	if err != nil {
		fmt.Println(err)
	}

	return convertedIP
}

func convertIP(address string) (string, error) {
	var (
		leftBraces  = "["
		period      = '.'
		rightBraces = "]"
		emptyStr    = ""
		strArr      = []string{}
	)

	for _, addr := range address {
		if addr == period {
			strArr = append(strArr, leftBraces, string(addr), rightBraces)
			continue
		}
		strArr = append(strArr, string(addr))
	}

	return strings.Join(strArr, emptyStr), nil
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
