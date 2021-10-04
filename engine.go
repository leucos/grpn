package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	errStackTooSmall = errors.New("Stack is too small for operation")
	errUnknownOP     = errors.New("Unknown operation")
)

type engine struct {
	stack    *Stack
	previous *Stack
}

func (e *engine) BulkEval(s string) error {
	if strings.TrimSpace(s) == "" {
		return nil
	}

	// Insert a space between [0-9] and other classes
	var re = regexp.MustCompile(`([0-9\.]*)([^0-9\.]*)`)
	s = re.ReplaceAllString(s, `$1 $2 `)
	s = strings.TrimSpace(s)

	if s != "undo" {
		e.previous.Flush()
		*e.previous = *e.stack
	}

	for _, item := range strings.Fields(s) {
		err := e.eval(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *engine) eval(i string) error {
	v, err := strconv.ParseFloat(i, 64)

	// Push the number to stack and return
	if err == nil {
		e.stack.Push(v)
		return nil
	}

	// Otherwise handle functions
	if op, ok := operations[i]; ok {
		err = op.command(e.stack)
		if err != nil {
			return err
		}
	} else if i == "undo" {
		e.Undo()
		return nil
	} else {
		return errUnknownOP
	}

	return nil
}

func (e *engine) Stack() *Stack {
	return e.stack
}

func (e *engine) PreviousStack() *Stack {
	return e.previous
}

func (e *engine) Undo() {
	if e.previous == nil {
		return
	}
	
	e.stack.Flush()
	*e.stack = *e.previous
	e.previous.Flush()
}
