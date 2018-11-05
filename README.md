# distrodetector

Detects which Linux distro or BSD a system is running.

## Features and limitations

`distrodetec` only aims to be able to detect the following:

* The 100 most popular Linux distros and BSDs, according to [distrowatch](https://distrowatch.com/).
* macOS with Homebrew.

Pull requests for additional systems are welcome!

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

## Help Needed

* More testing is always needed when detecting Linux distros and BSDs.
* Please test the distro detection on your distro/BSD and submit an issue or pull request if it should fail.

## General Info

* License: MIT
* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
