package main

import (
    "fmt"
    "os"
    "io"
    "net/http"
    "log"
)

func main() {
    input := os.Args[1:]
    fmt.Println("wget", input)
    if input[1] == "" {
        path, err := os.Getwd()
        check(err)
        input[1] = path
    }
    downloadSite("http://" + input[0], input[1])
}


func downloadSite(url string, path string) {
    resp, err := http.Get("http://example.com/")
    check(err)
    defer resp.Body.Close()
    fileWriter(path, resp)
}

func fileWriter(path string, information *http.Response) {
    file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
    check(err)
    defer file.Close()
	_, err = io.Copy(file, information.Body)
    check(err)
}


func check(e error) {
    if e != nil {
        log.Fatal(e.Error())
    }
}
