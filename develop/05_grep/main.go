package main

import (
    "fmt"
    "regexp"
    "strings"
    "flag"
)


func main() {
    isAfter := flag.Int("A", -1, "печатать +N строк после совпадения")
    isBefore := flag.Int("B", -1, "печатать +N строк до совпадения")
    isContext := flag.Int("C", -1, "печатать ±N строк вокруг совпадения")
    isCount := flag.Bool("c", false, "количество строк")
    isIgnore := flag.Bool("i", false, "игнорировать регистр")
    isInvert := flag.Bool("v", false, "вместо совпадения, исключать")
    isFixed := flag.Bool("f", false, "точное совпадение со строкой")
    isLineNum := flag.Bool("n", false, "печатать номер строки")
    flag.Parse()

    input := "c"
    res := "a\nb\nc\nd\ne"

    if *isAfter > -1 {
        fmt.Println("The result with flag After:")
        fmt.Println(after(res, input, *isAfter))
    }
    if *isBefore > -1 {
        fmt.Println("The result with flag Before:")
        fmt.Println(before(res, input, *isBefore))
    }
    if *isContext > -1 {
        fmt.Println("The result with flag Context:")
        fmt.Println(context(res, input, *isContext))
    }
    if *isCount {
        fmt.Println("The result with flag Count:")
        fmt.Println(count(res))
    }
    if *isIgnore {
        fmt.Println("The result with flag Ignore:")
        fmt.Println(ignore(res, input))
    }
    if *isInvert {
        fmt.Println("The result with flag Invert:")
        fmt.Println(invert(res, input))
    }
    if *isFixed {
        fmt.Println("The result with flag Fixed:")
        fmt.Println(fixed(res, input))
    }
    if *isLineNum {
        fmt.Println("The result with flag LineNum:")
        fmt.Println(lineNum(res, input))
    }

}

func after(text, pattern string, n int) string{
    re := regexp.MustCompile(pattern)
    res := ""
    strs := strings.Split(text, "\n")
    for i, v := range strs {
        if re.MatchString(v) {
            res += v + "\n"
            for j, k := range strs {
				if j > i && j <= i+n {
					if j == len(strs)-1 || j == i+n {
						res += k
					} else {
						res += k + "\n"
					}
				}
            }
        }
    }
    return res
}

func before(text, pattern string, n int) string {
    re := regexp.MustCompile(pattern)
    res := ""
    strs := strings.Split(text, "\n")
    for i, v := range strs {
        if re.MatchString(v) {
            for j := i - n; j < i; j++ {
                if j < 0 {
                    break
                }
                res += strs[j] + "\n"
            }
            res += v + "\n"
        }
    }
    return res[:len(res)-1]
}

func context(text, pattern string, n int) string {
    re := regexp.MustCompile(pattern)
    res := ""
    strs := strings.Split(text, "\n")
    for i, v := range strs {
        if re.MatchString(v) {
            for j := i - n; j < i; j++ {
                if j < 0 {
                    break
                }
                res += strs[j] + "\n"
            }
            res += v + "\n"
            for j := i + 1; j < i+n+1; j++ {
                if j >= len(strs) {
                    break
                } else if j != len(strs) - 1 {
                    res += strs[j] + "\n"
                } else {
                    res += strs[j]
                }
            }
        }
    }
    return res
}

func count(pattern string) int {
    return len(strings.Split(pattern, "\n"))
}

func ignore(text, pattern string) string {
    text = strings.ToLower(text)
    pattern = strings.ToLower(pattern)
    re := regexp.MustCompile(pattern)
    res := ""
    strs := strings.Split(text, "\n")
    for _, v := range strs {
        if re.MatchString(v) {
            res += v + "\n"
        }
    }
    res = res[:len(res)-1]
    return res
}

func invert(text, pattern string) string {
    re := regexp.MustCompile(pattern)
    res := ""
    strs := strings.Split(text, "\n")
    for _, v := range strs {
        if !re.MatchString(v) {
            res += v + "\n"
        }
    }
    res = res[:len(res)-1]
    return res
}

func fixed(text, pattern string) string {
    res := ""
    strs := strings.Split(text, "\n")
    for _, v := range strs {
        if v == pattern {
            res += v + "\n"
        }
    }
    res = res[:len(res)-1]
    return res
}


func lineNum(text, pattern string) string {
    re := regexp.MustCompile(pattern)
    res := ""
    strs := strings.Split(text, "\n")
    for i, v := range strs {
        if re.MatchString(v) {
            res += v + " " + fmt.Sprint(i) + "\n"
        }
    }
    res = res[:len(res)-1]
    return res
}
