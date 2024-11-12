package main

import (
	"gostd/deq/gopooldeq"
	"fmt"
)

func main() {
	q := gopooldeq.NewPoolDequeue(32)
	// str, okp0 := q.PopTail()
	// fmt.Println("pop is ", okp0," and str is ", str)

	ok1 := q.PushHead("x1")
	fmt.Println("ok1 is ", ok1, "\n")

	ok2 := q.PushHead("x2")
	fmt.Println("ok2 is ", ok2, "\n")

	ok3 := q.PushHead("x3")
	fmt.Println("ok3 is ", ok3, "\n")

	ok4 := q.PushHead("x4")
	fmt.Println("ok4 is ", ok4, "\n")

	ok5 := q.PushHead("x5")
	fmt.Println("ok5 is ", ok5, "\n")

	ok6 := q.PushHead("x6")
	fmt.Println("ok6 is ", ok6, "\n")

	ok7 := q.PushHead("x7")
	fmt.Println("ok7 is ", ok7, "\n")

	ok8 := q.PushHead("x8")
	fmt.Println("ok8 is ", ok8, "\n")

	str1, okp1 := q.PopTail()
	fmt.Println("pop is ", okp1, " and str is ", str1, "\n")
	str2, okp2 := q.PopTail()
	fmt.Println("pop is ", okp2, " and str is ", str2, "\n")

	ok11 := q.PushHead("x11")
	fmt.Println("ok11 is ", ok11, "\n")

	ok12 := q.PushHead("x12")
	fmt.Println("ok12 is ", ok12, "\n")
}
