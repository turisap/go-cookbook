package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		fmt.Println("__OH no: \n", <-ch)
		os.Exit(0)
	}()

	time.Sleep(1 * time.Second)
}
