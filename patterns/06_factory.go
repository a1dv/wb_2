package pattern

type letter interface {
    create()
}

func show(l letter) {
    l.create()
}

type A struct {
    val string
}

type B struct {
    val string
}

type C struct {
    val string
}

func (a A) create() {
    a.val = "A"
}

func (b B) create() {
    b.val = "B"
}

func (c C) create() {
    c.val = "C"
}
