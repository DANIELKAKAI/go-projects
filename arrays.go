package main

import "fmt"


func main(){
	var numbers = [5]int{2,3,4,5};
	
	for x:=0; x<4; x++ {
		fmt.Println(numbers[x])
	}

	for x:= range numbers {
		fmt.Println(numbers[x])
	}

	fmt.Println(len(numbers),cap(numbers))

}