package main

import "fmt"
func continua(){
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}


func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
