package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"pronghorn/pkg/cli"
)

func main() {

	command := cli.NewCommand()
	if err := command.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
