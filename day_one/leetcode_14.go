package main

import "fmt"

func main() {
	var strs = []string{"aaa", "bbbb", "ccccc"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	short := findShortestStr(strs)
	for i, v := range short {
		for j := 0; j < len(strs); j++ {
			//如果第j个strs的第i个字符与short的不匹配
			if strs[j][i] != byte(v) {
				return strs[j][:i]
			}
		}
	}
	//循环走完了全部匹配返回short
	return short
}
func findShortestStr(s []string) string {
	if len(s) == 0 {
		return ""
	}
	short := s[0]
	for _, v := range s {
		if len(v) < len(short) {
			short = v
		}
	}
	return short
}
