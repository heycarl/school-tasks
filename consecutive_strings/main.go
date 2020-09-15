package main

import (
	"fmt"
	"strings"
)

func longestConsec(strarr []string, k int) string {
	output := ""
    n := len(strarr)
    if n == 0 || k > n || k <= 0 {
		return output
	}
    for j := 0; j < n - k + 1; j++ {
        str := strings.Join(strarr[j:j + k], "")
        if len(str) > len(output) {
            output = str
        }
    }
    return output
}

func main() {
	fmt.Println(longestConsec([]string{"itvayloxrp","wkppqsztdkmvcuwvereiupccauycnjutlv","vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2))
}
