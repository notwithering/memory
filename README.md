# Memory

[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](LICENSE)

**Memory** is a simple package that provides helper functions to interact with the computer's memory from go.

## Example

```go
package main

import (
	"fmt"

	"github.com/notwithering/memory"
)

func main() {
	// data you do not want to be stored into memory for long
	var info string = "my super secret password"

	// zero out the data
	memory.Zero(&info)

	// there is no data to even print
	fmt.Println(info)
}
```
