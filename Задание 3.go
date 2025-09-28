package main

import "fmt"

func main() {
 ch := make(chan int, 1) 

 select {
 case val := <-ch:
  fmt.Println("Получено из канала:", val)
 default:
  fmt.Println("Данных нет")
 }
}
