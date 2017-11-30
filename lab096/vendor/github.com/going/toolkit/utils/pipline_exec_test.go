package utils

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestPipline(t *testing.T) {
	// Collect directories from the command-line
	var dirs []string
	if len(os.Args) > 1 {
		dirs = os.Args[1:]
	} else {
		dirs = []string{"."}
	}
	// Run the command on each directory
	for _, dir := range dirs {
		// find $DIR -type f # Find all files
		ls := exec.Command("find", dir, "-type", "f")

		// | grep -v '/[._]' # Ignore hidden/temporary files
		visible := exec.Command("cp", "aaa", "bbb")

		// | sort -t. -k2 # Sort by file extension
		sort := exec.Command("sort", "-t.", "-k2")

		// Run the pipeline
		output, stderr, err := Pipeline(ls, visible, sort)
		if err != nil {
			log.Printf("dir %q: %s", dir, err)
		}

		// Print the stdout, if any
		if len(output) > 0 {
			log.Printf("%q:\n%s", dir, output)
		}

		// Print the stderr, if any
		if len(stderr) > 0 {
			log.Printf("%q: (stderr)\n%s", dir, stderr)
		}
	}
}
