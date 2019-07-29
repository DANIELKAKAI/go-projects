package main

import "fmt"

func main() {

   var a,b,sum int;
   a = 44;
   b = 55;
   sum = add(a,b);
   d := subtract(a,b); /* static type declaration */

   fmt.Println(sum,d)
}

func add(x,y int) int{
	return x+y;
}

func subtract(x,y int) int{
	return x-y;
}