package main


import (
    "fmt"
    "bufio"
    "os"
    "flag"
    "log"
    "sort"
    "reflect"
    "strings"
    "unicode"
    "strconv"
)

type rowedString [][]string

type numString struct {
    numeric []int
    text []string
}

func main() {
    isRowChosen := flag.Int("k", -1, "указание колонки для сортировки")
    isReverse := flag.Bool("r", false, "сортировать в обратном порядке")
    isRepeatingAllowed := flag.Bool("u", false, "не выводить повторяющиеся строки")
    isNumeric := flag.Bool("n", false, "сортировать по числовому значению")
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

    fmt.Println("Input was:")
    print(input)
    if *isNumeric {
        input = numericSort(input)
    }
    if *isRowChosen > -1 {
        input = sortByRow(input, *isRowChosen)
    }
    if *isReverse {
        reverse(input)
    }
    if *isRepeatingAllowed {
        cleanRepeats(input)
    }
    if !*isNumeric && *isRowChosen == -1 && !*isReverse {
        sort.Strings(input)
    }
    fmt.Println("The result is:")
    print(input)
}

func reverse(s [] string) {
    swapF := reflect.Swapper(s)
    for i := 0; i < len(s); i++ {
        if (i >= len(s) - i - 1) {
            break
        }
        swapF(i, len(s)-i - 1)
    }
}

func check(e error) {
    if e != nil {
        log.Fatal(e.Error())
    }
}

func sortByRow(s []string, row int) []string{
    res := make(rowedString, len(s))
    for i, v := range s {
        res[i] = strings.Split(v, " ")
    }
    swapF := reflect.Swapper(s)
    for _, v := range res {
        swapF = reflect.Swapper(v)
        swapF(0, row-1)
    }
    sort.Sort(res)
    for _, v := range res {
        swapF = reflect.Swapper(v)
        swapF(0, row-1)
    }
    return convert(res)
}

func print(s []string) {
    fmt.Println("===================")
	for _, v := range s {
		fmt.Println(v)
	}
    fmt.Println("===================")
}

func convert(r rowedString) []string{
    res := make([]string, len(r))
    for i, v := range r {
        for _, s := range v {
            res[i] += s + " "
        }
    }
    return res
}

func numericSort(s []string) []string{
    i := make([]int, len(s))
    res := numString{i, s}
    for i, v := range s {
        res.numeric[i] = parse(v)
    }
    sort.Sort(res)
    return res.text
}

func parse(s string) int{
    tmp := ""
    res := 0
    if unicode.IsDigit(rune(s[0])) {
        for _, v := range s{
            if !unicode.IsDigit(v) {
                break
            }
            tmp += string(v)
        }
    } else {
        return res
    }
    res, _ = strconv.Atoi(tmp)
    return res
}

func cleanRepeats(s []string) {
    repeat := make(map[string]int)
    for i, v := range s {
        repeat[v]++
        if repeat[v] > 1 {
            s[i] = ""
        }
    }
}

func (r rowedString) Len() int{
    return len(r)
}

func (r rowedString) Less(i, j int) bool {
    return r[i][0] < r[j][0]
}

func (r rowedString) Swap(i , j int) {
    swapF := reflect.Swapper(r)
    swapF(i, j)
}

func (n numString) Len() int {
    return len(n.text)
}

func (n numString) Less(i, j int) bool {
    return n.numeric[i] < n.numeric[j]
}

func (n numString) Swap(i, j int) {
    swapF := reflect.Swapper(n.numeric)
    swapF(i, j)
    swapF = reflect.Swapper(n.text)
    swapF(i, j)
}
