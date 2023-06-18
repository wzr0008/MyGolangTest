package Concurrency

import (
	"fmt"
	"time"
)

func GoroutineTest5() {
	ticker := time.NewTicker(1 * time.Second)
	stop := time.After(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Ticking")
		case <-stop:
			fmt.Println("Stopping")
			return
		}
	}
}
