package main

import "fmt"

/*
 * @description: 使用go切片比较能力，直接暴力比较，算法复杂度为len(haystack) - len(needle)
 * @return:
 * @author: kechans
 * @date:
 */

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	lengthNeedle := len(needle)
	for i := 0; i < len(haystack)-lengthNeedle; i++ {
		if haystack[i:i+lengthNeedle] == needle {
			return i
		}
	}
	return 0
}
func main() {
	fmt.Printf("%d", strStr("sajeee", "aa"))
}
