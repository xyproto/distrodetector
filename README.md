# distrodetector

[![Build Status](https://travis-ci.org/xyproto/distrodetector.svg?branch=master)](https://travis-ci.org/xyproto/distrodetector) [![GoDoc](https://godoc.org/github.com/xyproto/distrodetector?status.svg)](http://godoc.org/github.com/xyproto/distrodetector) [![License](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/distrodetector/master/LICENSE) [![Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat)](http://goreportcard.com/report/xyproto/distrodetector)

Detects which Linux distro or BSD a system is running.

## Example use

```go
package main

import (
	"fmt"
	"github.com/xyproto/distrodetector"
)

func main() {
	distro := distrodetector.New()
	fmt.Println(distro.Name())
}
```

## Features and limitations

This package only aims to be able to detect the following:

* The 100 most popular Linux distros and BSDs, according to [distrowatch](https://distrowatch.com/).
* macOS with Homebrew.

Pull requests for additional systems are welcome!

## Testing

* More testing is always needed when detecting Linux distros and BSDs.
* Please test the distro detection on your distro/BSD and submit an issue or pull request if it should fail.

## General Info

* License: MIT
* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
