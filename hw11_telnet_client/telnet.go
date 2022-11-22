package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"syscall"
	"time"
)

var ErrConnectionWasClosedByPeer = errors.New("connection was closed by peer")

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		conn:    nil,
	}
}

type telnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func (t *telnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("create connection: %w", err)
	}

	t.conn = conn
	log.Printf("Connected to %s", t.address)

	return nil
}

func (t *telnetClient) Close() error {
	if t.conn != nil {
		if err := t.conn.Close(); err != nil {
			return fmt.Errorf("connection close: %w", err)
		}
	}
	return nil
}

func (t *telnetClient) Send() error {
	if _, err := io.Copy(t.conn, t.in); err != nil {
		if errors.Is(err, syscall.EPIPE) {
			return ErrConnectionWasClosedByPeer
		}
		return fmt.Errorf("write to socket: %w", err)
	}
	return nil
}

func (t *telnetClient) Receive() error {
	if _, err := io.Copy(t.out, t.conn); err != nil {
		return fmt.Errorf("read from socket: %w", err)
	}
	return nil
}
