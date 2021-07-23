package main

import (
	"fmt"
)

type runeStack struct {
	stack []rune
}

func newRuneStack() *runeStack {
	s := runeStack{stack: []rune{}}
	return &s
}

func (r *runeStack) isEmpty() bool {
	return len(r.stack) == 0
}

func (r *runeStack) push(arg rune) {
	r.stack = append(r.stack, arg)
}

func (r *runeStack) pop() (rune, bool) {
	value, ok := r.read()
	if ok {
		r.stack = r.stack[:len(r.stack)-1]
		return value, true

	} else {
		return ' ', false
	}

}

func (r *runeStack) read() (rune, bool) {
	if r.isEmpty() {
		return ' ', false
	} else {
		return r.stack[len(r.stack)-1], true
	}

}

func stackStringReverse(arg string) string {
	r := newRuneStack()
	for _, c := range arg {
		r.push(c)
	}

	result := ""
	for !r.isEmpty() {
		register, _ := r.pop()
		result = result + string(register)
	}

	return result
}

func main() {
	s := "abcde"
	fmt.Println(stackStringReverse(s)) // 'edcba'
}
