package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
    "log"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        var cmd *exec.Cmd
        dir, _ := os.Getwd()
        fmt.Printf("%s$ > ", dir)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSuffix(input, "\n")
        words := strings.Split(input, " ")

        switch words[0] {
        case "cd":
            err := os.Chdir(words[1])
            check(err)
        case "nc":
            netcatClient(fmt.Sprintf("%s:%s", words[1], words[2]))
        case "quit":
            os.Exit(0)
        default:
            cmd = exec.Command("bash", "-c", input)
            cmd.Stderr = os.Stderr
    	    cmd.Stdout = os.Stdout
            cmd.Run()
        }
    }
}

func check(e error) {
    if e != nil {
        log.Fatal(e.Error())
    }
}

func netcatClient(address string) {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        conn, err = net.Dial("udp", address)
        check(err)
    }
    _, err = io.Copy(conn, os.Stdin)
    check(err)
}
