package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	Timeflag := flag.Duration("Timeout", 10*time.Second, "Value of timeout in seconds")
	flag.Parse()
	ctx := context.Context(context.Background())
	ResponseChan := make(chan net.Conn)
	defer close(ResponseChan)
	Adresses := os.Args[1:]
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	conn, err := TryingToConn(ctx, Adresses, *Timeflag, ResponseChan)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		<-sigChan
		conn.Close()
		os.Exit(1)
	}()

	for {
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.TrimSpace(line) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
		fmt.Fprintf(conn, line+"\n")

		scn := bufio.NewScanner(conn)
		for scn.Scan() {
			fmt.Print("->: " + scn.Text())
		}

	}

}

func TryingToConn(ctx context.Context, addrs []string, t time.Duration, ch chan net.Conn) (net.Conn, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, t)
	defer cancel()

	go func() {
		conn, _ := net.Dial("tcp", addrs[len(addrs)-2]+":"+addrs[len(addrs)-1])
		ch <- conn
	}()

	select {
	case <-ctxTimeout.Done():
		return nil, errors.New("Context ended up")
	case val := <-ch:
		return val, nil
	}
}
