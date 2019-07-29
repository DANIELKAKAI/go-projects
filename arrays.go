package main

import "fmt"


func main(){
	var numbers = [4]int{2,3,4,5};
	
	for x:=0; x<4; x++ {
		fmt.Println(numbers[x])
	}

}