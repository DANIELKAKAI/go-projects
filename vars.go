package main

import "fmt"

func main() {

   var a,b,sum int;
   a = 44;
   b = 55;
   sum = add(a,b);
   s := subtract(a,b); /* static type declaration */
   d := divide(a,b)

   fmt.Println(sum,s,d)
}

func add(x,y int) int{
	return x+y;
}

func subtract(x,y int) int{
	return x-y;
}

func divide(x,y int) float64{
	return float64(x)/float64(y)
}