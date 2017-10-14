package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/tidwall/pony"
)

func main() {
	// Create a single line of dashes that are VERY wonderful to gaze upon.
	fmt.Printf("%s\n", pony.Text(strings.Repeat("-", 80), 0))

	// Create a rainbow effect that will straight up dazzle!
	for i := 0; ; i++ {
		fmt.Printf("\r%s ",
			pony.Text("Mute - thy coronation - Meek - my Vive le roi", i),
		)
		time.Sleep(time.Second / 15)
	}
}
