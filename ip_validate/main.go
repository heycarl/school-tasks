package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidIP(ip string) bool {
	if ip == "" {
		return false
	}
	okStr := "0987654321."
	for _, sym := range ip {
		if !strings.Contains(okStr, string(sym)){
			return false
		}
	}
	var ipSlice []string = strings.Split(ip, ".")
	if len(ipSlice) != 4 {
		return false
	}
	for _, item := range ipSlice {
		if string(item[0]) == "0" && len(string(item)) > 1{
			return false
		}
		var intIP int
		intIP, _ = strconv.Atoi(item)
		if intIP < 0 || intIP > 255 {
			return false
		}
	}
	return true
}

func main() {
	// fmt.Println(isValidIP(""))
	fmt.Println(isValidIP("123.456.789.0"))
}
