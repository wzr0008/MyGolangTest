package Challenge

import (
	"fmt"
	"unsafe"
)

type Node struct {
	data float64
	next *Node
}

func NodeTesting() {
	fmt.Println(unsafe.Sizeof(Node{}))
}
