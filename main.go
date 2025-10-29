package main

import _ "github.com/lib/pq"

import (
	"log"
	"os"

	"github.com/alternateved/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	st := state{&cfg}
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	err = cmds.register("login", handlerLogin)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{name: os.Args[1], args: os.Args[2:]}
	err = cmds.run(&st, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
