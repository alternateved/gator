package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/alternateved/gator/internal/config"
	"github.com/alternateved/gator/internal/database"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	st := state{database.New(db), &cfg}
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{name: os.Args[1], args: os.Args[2:]}
	err = cmds.run(&st, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
