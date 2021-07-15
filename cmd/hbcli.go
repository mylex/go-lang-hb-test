package main

import (
	"flag"
	"fmt"
	"github.com/mylex/go-lang-hb-test/server"
	"os"
)

func main() {

	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	servePort := serveCmd.Int("port", 8888, "port")

	if len(os.Args) < 2 {
		fmt.Println("expected Serve subcommands")
		os.Exit(1)
	}

	if os.Args[1] == "serve" {
		serveCmd.Parse(os.Args[2:])
		server.Init(*servePort)
		fmt.Println("Localhost server running on port:", *servePort)
	} else {
		fmt.Println("expected serve subcommands")
		os.Exit(1)
	}
}
