package command

import (
	"fmt"
	"strings"

	"github.com/funkygao/gocli"
)

type Activity struct {
	Ui  cli.Ui
	Cmd string
}

func (this *Activity) Run(args []string) (exitCode int) {
	return
}

func (*Activity) Synopsis() string {
	return "Register/List/Manipulate activity and activity type"
}

func (this *Activity) Help() string {
	help := fmt.Sprintf(`
Usage: %s activity [options]

    Register/List/Manipulate activity and activity type

`, this.Cmd)
	return strings.TrimSpace(help)
}
