package main

import "fmt"


func main(){
s :=  []string{"a","b","c"}

fmt.Println(s)

s_appended := append(s, "d","e", "f")

fmt.Println(s_appended)
fmt.Println(len(s), cap(s), len(s_appended), cap(s_appended))

}
