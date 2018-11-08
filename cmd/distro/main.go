package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/distrodetector"
)

const versionString = "distro 1.0"

func main() {
	nFlag := flag.Bool("n", false, "output only the detected distro name")
	aFlag := flag.Bool("a", false, "output a combined string with all available information")
	vFlag := flag.Bool("v", false, "version info")
	flag.Parse()

	if *vFlag {
		fmt.Println(versionString)
		return
	}

	distro := distrodetector.New()
	if *nFlag {
		fmt.Println(distro.Name())
		return
	}
	if *aFlag {
		fmt.Println(distro)
		return
	}
	// Output that should be possible to use as a drop-in replacement for python-distro
	name := distro.Name()
	version := distro.Version()
	if version == "" {
		version = "n/a"
	}
	codename := distro.Codename()
	if codename == "" {
		codename = "n/a"
	}
	fmt.Printf("Name: %s\nVersion: %s\nCodename: %s\n", name, version, codename)
}
