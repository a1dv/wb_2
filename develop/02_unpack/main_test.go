package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
)

func Test_unpack(t *testing.T) {
    testTable := []struct{
        input string
        output string
    }{
        {"a4b2","aaaabb"},
        {"22221",""},
        {"i1j0", "i"},
        {"qwerty", "qwerty"},
        {"",""},
    }

    for _, expected := range testTable {
        res, _ := unpack(expected.input)
        t.Logf("Calling unpack(%s), result is %s\n", expected.input, res)
        assert.Equal(t, expected.output, res, fmt.Sprintf("test : %s, func result : %s, correct: %s", expected.input, res, expected.output))
    }
}

func Test_validate(t * testing.T) {
    testTable := []struct {
        input string
        output bool
    } {
        {"45", false},
        {"241421421a21412", true},
        {"", true},
        {"a4a4", true},
    }

    for _, expected := range testTable {
        res := validate(expected.input)
        t.Logf("Calling validate(%s), result is %t\n", expected.input, res)
        assert.Equal(t, expected.output, res, fmt.Sprintf("test : %s, func result : %t, correct: %t", expected.input, res, expected.output))
    }
}
