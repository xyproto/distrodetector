package main

import (
	"fmt"
	"github.com/xyproto/distrodetector"
)

func main() {
	distro := distrodetector.New()
	fmt.Println(distro.Name())
}
