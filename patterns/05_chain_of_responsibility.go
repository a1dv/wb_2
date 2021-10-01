package pattern

import (
    "fmt"
    "math/rand"
)

type chain_el interface {
    decision()
}

type cathedra struct {
    enthropy int
}

type faculty struct {
    enthropy int
}

type profcom struct {
    enthropy int
}

type deanery struct {
    enthropy int
}
type student struct {
    group_number int
}

func (c cathedra) decision(s student) {
    if c.enthropy == s.group_number {
        c.help(s)
    } else {
        fmt.Println("Cathedra can't help you, go to deanery")
        d := deanery{rand.Intn(100)}
        d.decision(s)
    }
}

func (d deanery) decision(s student) {
    if d.enthropy == s.group_number {
        d.help(s)
    } else {
        fmt.Println("Deanery can't help you, go to profcom")
        p := profcom{rand.Intn(100)}
        p.decision(s)
    }
}

func (p profcom) decision(s student) {
    if p.enthropy == s.group_number {
        p.help(s)
    } else {
        fmt.Println("Profcom can't help you, go to faculty")
        f := faculty{rand.Intn(100)}
        f.decision(s)
    }
}

func (f faculty) decision(s student) {
    if f.enthropy == s.group_number {
        f.help(s)
    } else {
        fmt.Println("No one here can help you\nGo to cathedra again")
    }
}

func (c cathedra) help(s student) {
    fmt.Println("The problem solved")
}

func (d deanery) help(s student) {
    fmt.Println("The problem solved")
}

func (f faculty) help(s student) {
    fmt.Println("The problem solved")
}

func (p profcom) help(s student) {
    fmt.Println("The problem solved")
}
