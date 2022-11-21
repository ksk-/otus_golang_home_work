package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	log.SetOutput(os.Stderr)

	timeout := flag.Duration("timeout", 10*time.Second, "connect to server timeout")
	flag.Parse()

	address, err := parseAddress(os.Args[1:])
	if err != nil {
		log.Fatalf("failed to parse arguments: %v", err)
	}

	client := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)
	if err = client.Connect(); err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer cancel()
		runSender(client)
	}()
	go func() {
		defer cancel()
		runReceiver(client)
	}()

	<-ctx.Done()

	if err = client.Close(); err != nil {
		log.Printf("failed to close client: %v", err)
	}
}

func parseAddress(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("should set at 2 positional args")
	}

	host, port := args[len(args)-2], args[len(args)-1]
	if !isInteger(port) {
		return "", errors.New("port should be integer value")
	}

	return net.JoinHostPort(host, port), nil
}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func runSender(client TelnetClient) {
	err := client.Send()
	if err == nil {
		log.Printf("EOF")
		return
	}

	if errors.Is(err, ErrConnectionWasClosedByPeer) {
		log.Printf("Connection was closed by peer")
	} else {
		log.Printf("failed to send to server: %v", err)
	}
}

func runReceiver(client TelnetClient) {
	if err := client.Receive(); err != nil {
		log.Printf("failed to receive from server: %v", err)
	}
}
