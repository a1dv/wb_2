package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
    "strconv"
)


func main() {
    ctx, cancel := context.WithCancel(context.Background())
    input := flag.Duration("timeout", time.Second * 10, "таймаут на подключение к серверу")
    flag.Parse()
    if len(flag.Args()) < 2 {
        flag.Usage()
        os.Exit(1)
    }
    host := flag.Args()[0]
    port := flag.Args()[1]

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT)
	go onCancel(signalChan, cancel)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *input)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go read(conn, cancel)
	go write(conn, cancel)

	<-ctx.Done()
	log.Println("finish telnet client")
}

func read(conn net.Conn, cancelFunc context.CancelFunc) {
	scanner := bufio.NewScanner(conn)
	for {
		if !scanner.Scan() {
			log.Printf("socket was closed")
			cancelFunc()
			break
		}
		fmt.Println(scanner.Text())
	}
}

func write(conn net.Conn, cancelFunc context.CancelFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			log.Printf("no input")
			cancelFunc()
			break
		}
		_, err := conn.Write([]byte(fmt.Sprintln(scanner.Text())))
		check(err)
	}
}

func check(e error) {
    if e != nil {
        log.Fatal(e.Error())
    }
}

func onCancel(sigCh chan os.Signal, cancelFunc context.CancelFunc) {
    log.Println(<-sigCh)
    cancelFunc()
}
func parseFlag(input string) time.Duration {
    res, _ := strconv.ParseFloat(input[:len(input)-1], 64)
    fmt.Println(input[:len(input)-1])
    return time.Duration(res)
}
