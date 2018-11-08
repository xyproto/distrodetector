package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/distrodetector"
)

const versionString = "distro 1.0"

func main() {
	aFlag := flag.Bool("a", false, "output all detected information")
	vFlag := flag.Bool("v", false, "version info")
	flag.Parse()

	if *vFlag {
		fmt.Println(versionString)
		return
	}

	distro := distrodetector.New()
	if *aFlag {
		fmt.Println(distro)
	} else {
		fmt.Println(distro.Name())
	}
}
