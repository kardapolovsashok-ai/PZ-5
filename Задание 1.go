package main

import (
	"fmt"
	"time"
)

func main() {
	resultChan := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		resultChan <- "Результат запроса"
	}()
	select {
	case result := <-resultChan:
		fmt.Println("Успешный ответ:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Ошибка: таймаут запроса")
	}
}
