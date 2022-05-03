package Slice

import "fmt"

func main() {
	var s = []string{"1", "2", "3"}
	s = modifySlice2(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i = append(i, "4")
}

func modifySlice2(i []string) []string {
	return append(i, "4")
}
