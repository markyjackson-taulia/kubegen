package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/errordeveloper/testcli"

	"github.com/errordeveloper/kubegen/cmd/kubegen/assets/commands"
)

func main() {
	for filename, command := range commands.Commands {
		c := testcli.GoRun("../main.go", append([]string{"../bundle.go", "../module.go", "../self_upgrade.go"}, command...)...)
		c.Run()
		if !c.Success() {
			fmt.Fprintf(os.Stderr, "Command %v was expected to succeed, but failed with error: %s\n%s\n", command, c.Error(), c.StdoutAndStderr())
			os.Exit(1)
		}
		if err := ioutil.WriteFile(filename, []byte(c.Stdout()), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write output to %q – %v", filename, err)
		}
	}
}
