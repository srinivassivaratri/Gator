package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/srinivassivaratri/RSSAggregator/internal/config"
	"github.com/srinivassivaratri/RSSAggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	// Get command line args (excluding program name)
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		return
	}

	// Create command from args
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Run the command
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}
}
