package pattern

import (
    "fmt"
)

type operations struct {
    ops []string
}

type coeffs struct {
    vals []float64
}

type function struct {
    visual string
}

func (f *function) Printing_system(ops operations, coeffs coeffs) {
    f.visual = Construct_function(coeffs.vals, ops.ops)
}

func (f *function) get_string() string {
    return f.visual
}

func Construct_function(coeffs []float64, ops []string) string{
    return fmt.Sprintf("%fx^2 %s %fx %s %f = 0", coeffs[0], ops[0], coeffs[1], ops[1], coeffs[2])
}

type Visitor interface {
    do_for_coeffs(c coeffs)
    do_for_ops(o operations)
    do_for_func(f function)
}

type sizing_visitor struct {
    size int
}

func (s *sizing_visitor) do_for_coeffs(c coeffs) {
    s.size = len(c.vals)
}

func (s *sizing_visitor) do_for_ops(o operations) {
    s.size = len(o.ops)
}

func (s *sizing_visitor) do_for_func(f function) {
    s.size = len(f.visual)
}
