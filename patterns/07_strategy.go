package pattern

import (
    "fmt"
    "strings"
)

type Context struct {
    strat Strategy
}

type Strategy interface {
    execute()
}

type lower_case_print struct {
    val string
}

type upper_case_print struct {
    val string
}

func (c *Context) execute_strat() {
    c.strat.execute()
}

func (c *Context) set_strat(strat Strategy){
    c.strat = strat
}

func (l lower_case_print) execute() {
    fmt.Println(strings.ToLower(l.val))
}

func (u upper_case_print) execute() {
    fmt.Println(strings.ToUpper(u.val))
}
