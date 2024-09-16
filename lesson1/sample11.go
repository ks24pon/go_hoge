package main

import (
	"fmt"
)

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main11() {
	var hex Hex = 100
	fmt.Println(hex.String())
}
