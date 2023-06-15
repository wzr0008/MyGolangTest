package Challenge

import (
	"fmt"
	"strings"
)

func StrTest() {
	var str string = "This is an example of a string"
	fmt.Println("Checking Prefix")
	fmt.Println(strings.HasPrefix(str, "Th"))
	fmt.Println("Checking Suffix")
	fmt.Println(strings.HasSuffix(str, "ing"))
	fmt.Println("Checking Contains")
	fmt.Println(strings.Contains(str, "example"))
	str = "Hi, I'm Marc, Hi."
	fmt.Println("Index of some words")
	fmt.Println(strings.Index(str, "Hi"))
	replace := strings.Replace(str, "Hi", "Halohana", -1)
	fmt.Println(replace)
}
