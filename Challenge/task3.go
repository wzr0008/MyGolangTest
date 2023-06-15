package Challenge

import (
	"fmt"
	"time"
)

func TimeTest() {
	t := time.Now()
	time.Sleep(1 * time.Second)
	since := time.Since(t)
	fmt.Println(since)
}
