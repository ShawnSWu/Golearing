package main

import "fmt"

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Printf("原始的位址是：%p\n", &s)
}

func modifySlice(i []string) {
	copyNewOne := i[:len(i)-1]            //重新切片出一個
	copyNewOne = append(i, copyNewOne...) // 透過append複製出一個
	copyNewOne[0] = "9"                   // 此時修改的是新複製的Array複製出一個

	fmt.Printf("複製後的的位址是 %p\n", &copyNewOne)
}
