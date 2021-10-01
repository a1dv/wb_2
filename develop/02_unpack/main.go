package main

import (
	"fmt"
    "unicode"
)

func main() {
    fmt.Println(unpack(""))
}

func unpack(input string) (string, error){
    if !validate(input) {
        return "", fmt.Errorf("incorrect string")
    }
    letters := []rune(input)
    res := ""
    for i, v := range letters {
        if unicode.IsDigit(v) && i != 0  && !unicode.IsDigit(letters[i-1]){
            for j:= 0; j < int(v - '0')-1; j++ {
                res += string(letters[i - 1])
            }
            if int(v - '0') == 0 {
                res = res[:len(res)-1]
            }
        } else if !unicode.IsDigit(v) {
            res += string(v)
        }
    }
    return res, nil
}

func validate(input string) bool{
    for _, v := range(input) {
        if !unicode.IsDigit(v) {
            return true
        }
    }
    if input != "" {
        return false
    } else {
        return true
    }
}
