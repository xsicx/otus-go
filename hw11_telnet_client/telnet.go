package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Send() error
	Receive() error
	Close() error
}

type telnetClient struct {
	conn    net.Conn
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (t *telnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return err
	}

	t.conn = conn
	return nil
}

func (t *telnetClient) Close() error {
	return t.conn.Close()
}

func (t *telnetClient) Send() error {
	_, err := io.Copy(t.conn, t.in)
	return err
}

func (t *telnetClient) Receive() error {
	_, err := io.Copy(t.out, t.conn)
	return err
}
