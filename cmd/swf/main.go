package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/funkygao/gafka/ctx"
	"github.com/funkygao/gocli"
	"github.com/funkygao/swf"
)

func init() {
	ctx.LoadFromHome()
}

func main() {
	app := os.Args[0]
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}

		if arg == "--generate-bash-completion" {
			for name := range commands {
				fmt.Println(name)
			}

			return
		}
	}

	c := cli.NewCLI(app, swf.Version+"-"+swf.BuildId)
	c.Args = os.Args[1:]
	c.Commands = commands
	c.HelpFunc = func(m map[string]cli.CommandFactory) string {
		var buf bytes.Buffer
		buf.WriteString(cli.BasicHelpFunc(app)(m))
		return buf.String()
	}

	exitCode, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	} else if c.IsVersion() {
		os.Exit(0)
	}

	os.Exit(exitCode)
}
