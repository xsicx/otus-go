package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connect timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Not enough arguments. Usecase: go-telnet --timeout=3s 1.1.1.1 123")
	}

	address := net.JoinHostPort(args[0], args[1])
	client := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stderr, "...Connected to %s\n", address)
	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := client.Send(); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stderr, "...EOF")
		cancel()
	}()

	go func() {
		if err := client.Receive(); err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
		cancel()
	}()

	<-ctx.Done()
}
