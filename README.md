![spartan logo](https://gitlab.com/tobiaskoch/spartan/-/raw/main/assets/spartan.png)

# Spartan
Common functions for go projects.

The documentation is available here: https://pkg.tk-software.de/spartan

```go
package main

import (
	"fmt"

	"pkg.tk-software.de/spartan/io/file"
)

func main() {
	fmt.Printf("Does main.go exist: %t\n", file.Exists("main.go"))
}

```

## License
**spartan** © 2023 [Tobias Koch](https://www.tk-software.de). Released under a [BSD-style license](https://gitlab.com/tobiaskoch/spartan/-/blob/main/LICENSE).