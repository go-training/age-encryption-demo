package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"filippo.io/age"
	"golang.org/x/term"
)

func main() {
	k, err := age.GenerateX25519Identity()
	if err != nil {
		log.Fatalf("internal error: %v", err)
	}

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Fprintf(os.Stderr, "Public key: %s\n", k.Recipient())
	}

	fmt.Fprintf(os.Stdout, "# created: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(os.Stdout, "# public key: %s\n", k.Recipient())
	fmt.Fprintf(os.Stdout, "%s\n", k)
}
