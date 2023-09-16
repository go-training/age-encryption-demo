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
	privateKey, err := age.GenerateX25519Identity()
	if err != nil {
		log.Fatalf("internal error: %v", err)
	}
	publicKey := privateKey.Recipient()

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Fprintf(os.Stderr, "Public key: %s\n", publicKey)
	}

	fmt.Fprintf(os.Stdout, "# created: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(os.Stdout, "# public key: %s\n", publicKey)
	fmt.Fprintf(os.Stdout, "%s\n", privateKey)
}
