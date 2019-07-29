package main

import "fmt"

type Books struct {
	title string
	author string
	id int
	price float64
}

func main() {
   var book1,book2 Books;
   

   book1.title = "swiney todd";
   book1.author = "dan"
   book1.id = 1
   book1.price = 7.99

   book2.title = "tom sawyer";
   book2.author = "joe"
   book2.id = 2
   book2.price = 5.69


   fmt.Println(book1,book2)
}