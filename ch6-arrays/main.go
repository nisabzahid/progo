package main

import "fmt"



func main (){

array := [3]string{"a", "b", "c"}

fmt.Println(array)


fmt.Println("hello nisab")

array1 := [3]string{"x", "y", "z"}

for index, value:= range array1{
fmt.Println("index: ", index, " value: ", value)

}



}
