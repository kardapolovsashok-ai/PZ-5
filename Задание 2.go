package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func main() {
	done := make(chan bool)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("Работаю...")
			case <-done:
				fmt.Println("Получен сигнал завершения. Останавливаю горутину...")
				return
			}
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	fmt.Println("Получен сигнал завершения от ОС")
	close(done)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Приложение завершено корректно")
}
