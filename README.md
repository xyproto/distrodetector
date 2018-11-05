# distrodetector

Detect which Linux distro or BSD a system is running

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

## General Info

* License: MIT
