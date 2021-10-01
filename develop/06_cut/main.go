package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "log"
    "strconv"
    "flag"
)

func main() {
    isFields := flag.String("f", "1", "выбрать поля (колонки)")
    isDelimiter := flag.String("d", " ", "использовать другой разделитель")
    isSeparated := flag.Bool("s", false, "только строки с разделителем")
    flag.Parse()

    fileName := os.Args[len(os.Args)-1]
    path, err := os.Getwd()
    check(err)

    file, err := os.Open(path + "/" + fileName)
    check(err)
    defer file.Close()
    var input []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    err = scanner.Err()
    check(err)
    fields := *isFields
    delim := *isDelimiter
    fmt.Println(cut(input, delim, fields, *isSeparated))
}

func cut(str []string, delimiter string, col string,isSeparated bool) [][]string{
    rows := make([][]string, len(str))
    for i, v := range str {
        rows[i] = strings.Split(v, delimiter)
        if isSeparated {
            if len(rows[i]) == 1 {
                rows[i] = []string{""}
            }
        }
    }
    return makeRes(size(col), rows)
}

func check(e error) {
    if e != nil {
        log.Fatal(e.Error())
    }
}

func size(col string) []int{
    sizes := make([]int, 0)
    size := 0
    var err error
    if !(isComma(col)) {
        size, err = strconv.Atoi(col)
        check(err)
        if size < 0 {
            for i := 1; i <= -size; i++ {
                sizes = append(sizes, i)
            }
        } else {
            sizes = append(sizes, size)
        }
    } else {
        for _, v := range strings.Split(col, ",") {
            size, err = strconv.Atoi(v)
            check(err)
            sizes = append(sizes, size)
        }
    }
    return sizes
}

func isComma(col string) bool{
    for _, v := range col {
        if string(v) == "," {
            return true
        }
    }
    return false
}

func makeRes(sizes []int, rows[][]string) [][]string{
    res := make([][]string, len(rows))
    for i := range rows {
        for _, v := range sizes {
            if v - 1 >= 0 && v - 1 < len(rows[i]) {
                res[i] = append(res[i], rows[i][v-1])
            }
        }
    }
    return res
}
