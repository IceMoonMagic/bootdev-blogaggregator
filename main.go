package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/icemoonmagic/bootdev-blogaggregator/internal/config"
	"github.com/icemoonmagic/bootdev-blogaggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg := config.Read()
	s := state{}
	s.cfg = &cfg
	s.db = getDatabase(cfg.DbUrl)

	if err := getCommands().run(&s, parseCommand()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getDatabase(dbUrl string) *database.Queries {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	return dbQueries
}

func getCommands() commands {
	cmds := commands{}
	cmds.allcmds = make(map[string]func(*state, command) error)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("users", handlerGetUsers)
	cmds.register("reset", handlerReset)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	return cmds
}

func parseCommand() command {
	args := os.Args
	if err := checkArgsCount(len(args), 2, 256); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := command{}
	cmd.name = args[1]
	cmd.args = args[2:]
	return cmd
}
