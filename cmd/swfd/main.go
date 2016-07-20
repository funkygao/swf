package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/funkygao/swf"
	"github.com/funkygao/swf/cmd/swfd/server"
	//_ "github.com/go-sql-driver/mysql"
)

func init() {
	server.ParseFlags()

	if server.Options.ShowVersion {
		fmt.Fprintf(os.Stderr, "%s-%s\n", swf.Version, swf.BuildId)
		os.Exit(0)
	}

	if swf.BuildId == "" {
		fmt.Fprintf(os.Stderr, "empty BuildId, please rebuild with build.sh\n")
	}

	debug.SetGCPercent(800)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	server.ValidateFlags()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	fmt.Fprintln(os.Stderr, strings.TrimSpace(logo))

	server.New().ServeForever()
}
