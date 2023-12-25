package main

import (
	"fmt"
	"jf/Ratatouille/staging"
	"runtime"
)

func main() {
	fmt.Printf("Hello Ratatouille\n\n\n")
	operatingSystem := runtime.GOOS
	staging.Stage(operatingSystem)

}
