package main

import (
	m "math"
	r "math/rand"
	"time"
	// could be "math/cmplx"
)

type operation struct {
	label       string
	description string
	command     func(*Stack) error
}

var (
	add = operation{
		"+",
		"(y+x) adds x and y on the stack",
		func(s *Stack) error {
			if s.Size() < 2 {
				return errStackTooSmall
			}

			s.Push(s.Pop() + s.Pop())

			return nil
		},
	}
	minus = operation{
		"-",
		"(y-x) takes x from y on the stack",
		func(s *Stack) error {
			if s.Size() < 2 {
				return errStackTooSmall
			}
			x, y := s.Pop(), s.Pop()
			s.Push(y - x)

			return nil
		},
	}
	multiply = operation{
		"*",
		"(y*x) multiplies x with y on the stack",
		func(s *Stack) error {
			if s.Size() < 2 {
				return errStackTooSmall
			}
			s.Push(s.Pop() * s.Pop())

			return nil
		},
	}
	divide = operation{
		"/",
		"(y/x) divides y by x on the stack",
		func(s *Stack) error {
			if s.Size() < 2 {
				return errStackTooSmall
			}
			x, y := s.Pop(), s.Pop()
			s.Push(y / x)

			return nil
		},
	}
	power = operation{
		"^",
		"(y**x) raises y to the power of x on the stack",
		func(s *Stack) error {
			if s.Size() < 2 {
				return errStackTooSmall
			}
			x, y := s.Pop(), s.Pop()
			s.Push(m.Pow(y, x))

			return nil
		},
	}
	exp = operation{
		"$",
		"(e**x) raises e to the power of x on the stack",
		func(s *Stack) error {
			if s.Size() < 1 {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Exp(x))

			return nil
		},
	}
	log = operation{
		"log",
		"(log x) logs x on the stack",
		func(s *Stack) error {
			if s.Size() < 1 {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Log(x))

			return nil
		},
	}
	sqrt = operation{
		"sqrt",
		"(sqrt x) sqrt x on the stack",
		func(s *Stack) error {
			if s.Size() < 1 {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Sqrt(x))

			return nil
		},
	}
	pi = operation{
		"pi",
		"(pi) pushes π on the stack",
		func(s *Stack) error {
			s.Push(m.Pi)
			return nil
		},
	}
	sin = operation{
		"sin",
		"(sin x) sins x on the stack",
		func(s *Stack) error {
			if s.Size() < 1 {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Sin(x))

			return nil
		},
	}
	cos = operation{
		"cos",
		"(cos x) coses x on the stack",
		func(s *Stack) error {
			if s.Size() < 1 {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Cos(x))

			return nil
		},
	}
	tan = operation{
		"tan",
		"(tan x) tans x on the stack",
		func(s *Stack) error {
			if s.Size() < 1 {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Tan(x))

			return nil
		},
	}
	asin = operation{
		"asin",
		"(asin x) asins x on the stack",
		func(s *Stack) error {
			if s.Size() < 1 {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Asin(x))

			return nil
		},
	}
	acos = operation{
		"acos",
		"(acos x) acoses x on the stack",
		func(s *Stack) error {
			if s.Size() < 1 {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Acos(x))
			return nil
		},
	}
	atan = operation{
		"atan",
		"(atan x) atans x on the stack",
		func(s *Stack) error {
			if s.IsEmpty() {
				return errStackTooSmall
			}
			x := s.Pop()
			s.Push(m.Atan(x))
			return nil
		},
	}
	sum = operation{
		"sum",
		"sums all values on the stack",
		func(s *Stack) error {
			var sum float64

			for !s.IsEmpty() {
				sum += s.Pop()
			}
			s.Push(sum)
			return nil
		},
	}
	mean = operation{
		"mean",
		"takes the mean of all the values on the stack",
		func(s *Stack) error {
			var sum float64
			var count float64

			if s.IsEmpty() {
				return errStackTooSmall
			}

			for !s.IsEmpty() {
				sum += s.Pop()
				count++
			}

			s.Push(sum / count)
			return nil
		},
	}
	clear = operation{
		"clear",
		"clears the stack",
		func(s *Stack) error {
			for !s.IsEmpty() {
				s.Pop()
			}
			return nil
		},
	}
	dup = operation{
		"dup",
		"dup last element of the stack",
		func(s *Stack) error {
			if !s.IsEmpty() {
				s.Dup()
			}
			return nil
		},
	}
	swap = operation{
		"swap",
		"swap two last element of the stack",
		func(s *Stack) error {
			if s.Size() > 1 {
				s.Swap()
			}
			return nil
		},
	}
	drop = operation{
		"drop",
		"drops last element of the stack",
		func(s *Stack) error {
			if s.Size() > 0 {
				s.Pop()
			}
			return nil
		},
	}
	rand = operation{
		"rand",
		"pushes a random number between 0 and 1 onto the stack",
		func(s *Stack) error {
			r := r.New(r.NewSource(time.Now().UnixNano()))
			s.Push(r.Float64())
			return nil
		},
	}
	c = operation{
		"c",
		"pushes the speed of light C (299792458 ms⁻¹) onto the stack",
		func(s *Stack) error {
			s.Push(299792458)
			return nil
		},
	}
	g = operation{
		"G",
		"pushes igravitational constant (6.674×10⁻¹¹ m³kg⁻¹s⁻²)",
		func(s *Stack) error {
			s.Push(0.0000000000667408)
			return nil
		},
	}
	gg = operation{
		"g",
		"pushes earth gravitational force (9.81 ms⁻²) onto the stack",
		func(s *Stack) error {
			s.Push(9.81)
			return nil
		},
	}
	kilo = operation{
		"kilo",
		"pushes kilo unit (1024) onto the stack",
		func(s *Stack) error {
			s.Push(1024)
			return nil
		},
	}
	mega = operation{
		"mega",
		"pushes mega unit (1024²) onto the stack",
		func(s *Stack) error {
			s.Push(1024 * 1024)
			return nil
		},
	}
	giga = operation{
		"giga",
		"pushes giga unit (1024³) onto the stack",
		func(s *Stack) error {
			s.Push(1024 * 1024 * 1024)
			return nil
		},
	}
	tera = operation{
		"tera",
		"pushes tera unit (1024⁴) onto the stack",
		func(s *Stack) error {
			s.Push(1024 * 1024 * 1024 * 1024)
			return nil
		},
	}
)

var operationsKeys = []*operation{
	&add,
	&minus,
	&multiply,
	&divide,
	&power,
	&exp,
	&log,
	&sqrt,
	&pi,
	&sin,
	&cos,
	&tan,
	&asin,
	&acos,
	&atan,
	&sum,
	&mean,
	&clear,
	&rand,
	&c,
	&g,
	&gg,
	&kilo,
	&mega,
	&giga,
	&tera,
}

var operations = map[string]*operation{
	add.label:      &add,
	minus.label:    &minus,
	multiply.label: &multiply,
	divide.label:   &divide,
	power.label:    &power,
	"pow":          &power, // alias
	"**":           &power, // alias
	exp.label:      &exp,
	log.label:      &log,
	sqrt.label:     &sqrt,
	"sq":           &sqrt, // alias
	pi.label:       &pi,
	sin.label:      &sin,
	cos.label:      &cos,
	tan.label:      &tan,
	asin.label:     &asin,
	acos.label:     &acos,
	atan.label:     &atan,
	sum.label:      &sum,
	mean.label:     &mean,
	clear.label:    &clear,
	dup.label:      &dup,
	swap.label:     &swap,
	drop.label:     &drop,
	rand.label:     &rand,
	c.label:        &c,
	g.label:        &g,
	gg.label:       &gg,
	kilo.label:     &kilo,
	mega.label:     &mega,
	giga.label:     &giga,
	tera.label:     &tera,
}
