package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/srinivassivaratri/RSSAggregator/internal/config"
	"github.com/srinivassivaratri/RSSAggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries // Holds database query functions to interact with PostgreSQL
	cfg *config.Config    // Stores app configuration like database connection URL
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to db: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	// When you run the program:
	// 1. Reads config from ~/.gatorconfig.json
	// 2. Connects to PostgreSQL database
	// 3. Sets up command handlers
	// 4. Waits for CLI command

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
