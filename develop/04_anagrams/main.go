package main

import (
    "fmt"
    "sort"
    "reflect"
    "strings"
)

type runePack []rune

func main(){
    words := []string{"пятак", "тяпка", "пятка", "листок", "слиток", "столик", "Волос", "слово"}
    fmt.Println(anagrams(words))
}

func anagrams(words []string) map[string][]string{
    res := make(map[string][]string)
    for _, v := range words {
        v = strings.ToLower(v)
        res[sortString(v)] = append(res[sortString(v)], v)
    }
    for i, _ := range res {
        res[res[i][0]] = res[i]
        sort.Strings(res[res[i][0]])
        delete(res, i)
    }
    return res
}

func (r runePack) Len() int{
    return len(r)
}

func (r runePack) Swap(i, j int) {
    swapF := reflect.Swapper(r)
    swapF(i, j)
}

func (r runePack) Less(i, j int) bool {
    return r[i] < r[j]
}

func sortString(s string) string {
    r := []rune(s)
    sort.Sort(runePack(r))
    return string(r)
}
