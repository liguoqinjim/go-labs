package cli

import (
	"bufio"
	"io"
	"os"
	"os/exec"

	"github.com/murlokswarm/log"
)

// Exec executes cmd with args.
func Exec(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)

	cmdout, err := command.StdoutPipe()
	if err != nil {
		return err
	}

	cmderr, err := command.StderrPipe()
	if err != nil {
		return err
	}

	go printOutput(cmdout, os.Stdout)
	go printOutput(cmderr, os.Stderr)

	if err = command.Start(); err != nil {
		return err
	}
	return command.Wait()
}

func printOutput(r io.Reader, output io.Writer) {
	reader := bufio.NewReader(r)
	b := make([]byte, 256)

	for {
		n, err := reader.Read(b)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Error(err)
			continue
		}
		output.Write(b[:n])
	}
}
