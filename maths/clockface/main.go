package main

import (
	"os"
	"time"

	clockFace "github.com/AmiyoKm/testing/maths"
)

func main() {
	t := time.Now()
	clockFace.SVGWriter(os.Stdout, t)
}
