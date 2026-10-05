package main

import "fmt"

func main() {
	seq := 2000
	
	data := []byte("ab")

	ack := seq + len(data)

	fmt.Println("データ:",data)
	fmt.Println("バイト数:",len(data))
	fmt.Println("次に欲しい番号:",ack)
}