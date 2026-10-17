package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func run(path string) error {
	old := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(old)
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()
	if err := trace.Start(output); err != nil {
		return err
	}
	defer trace.Stop()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer listener.Close()
	client, err := net.DialTimeout("tcp", listener.Addr().String(), time.Second)
	if err != nil {
		return err
	}
	defer client.Close()
	peer, err := listener.Accept()
	if err != nil {
		return err
	}
	defer peer.Close()
	// Keep a real TCP peer open without writing until the read finishes.
	stop := make(chan struct{})
	done := make(chan struct{})
	go func() {
		defer close(done)
		<-stop
	}()
	defer func() { close(stop); <-done }()
	if err := client.SetReadDeadline(time.Now().Add(200 * time.Millisecond)); err != nil {
		return err
	}
	var buf [1]byte
	_, err = client.Read(buf[:])
	if timeout, ok := err.(net.Error); !ok || !timeout.Timeout() {
		return fmt.Errorf("expected read timeout, got %v", err)
	}
	fmt.Println("TCP read timed out; trace saved")
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: netpoll TRACE_PATH")
		os.Exit(2)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
