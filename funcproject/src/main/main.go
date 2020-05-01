package main

import (
	"fmt"
	"log"
)

func main() {

	// map操作
	// target := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	// mapController(target, false, 3)
	// mapController(target, true, 3)

	// ポインタ
	// targetStr := "testOutput"
	// nonePointer(targetStr)
	// fmt.Println(targetStr)
	// pointer(&targetStr)
	// fmt.Println(targetStr)

	// defer
	// if false {
	// 	log.Fatal("error!")
	// }
	// defer fmt.Println("defer!")

	// パニック
	defer func() {
		err := recover()
		if err != nil {
			// runtime error: index out of range
			log.Fatal(err)
		}
	}()

	a := []int{1, 2, 3}
	fmt.Println(a[10]) // パニックが発生

}

// ポインタON
func pointer(target *string) {
	*target = "aaa"
	fmt.Println(*target)
}

// ポインタOFF
func nonePointer(target string) {
	target = "bbb"
	fmt.Println(target)
}

// map操作
func mapController(targets map[int]string, isDel bool, indexDel int) {
	if isDel {
		delete(targets, indexDel)
	}
	for i := 0; i <= len(targets); i++ {
		val, isVal := targets[i]
		if isVal {
			fmt.Println(val)
		} else {
			fmt.Println("none")
		}

	}
}

// 関数リテラル
var sum func(target []string) = func(target []string) {
	fmt.Println(target[0:1])
	fmt.Println(target[1:2])
	fmt.Println(target[3:3])
	fmt.Println(target[3:])
	fmt.Println(target[:4])
}
