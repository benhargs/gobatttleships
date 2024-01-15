package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type PlayerInput interface {
	Fetch() (string, error)
}

type StubPlayerInput struct {
	inputs  []string
	current int
}

func (i *StubPlayerInput) Fetch() (string, error) {
	if i.current >= len(i.inputs) {
		return "", errors.New("No input")
	}

	text := i.inputs[i.current]
	i.current++
	return text, nil
}

func readAndPrintAll(input PlayerInput) {
	text, err := input.Fetch()

	for err == nil {
		fmt.Println(text)

		text, err = input.Fetch()
		if text == "stop" {

			err = errors.New("stop the code")
		}
	}
}

type KeyboardPlayerInput struct {
	scanner *bufio.Scanner
}

func NewKeyboardPlayerInput() *KeyboardPlayerInput {
	scanner := bufio.NewScanner(os.Stdin)
	return &KeyboardPlayerInput{scanner: scanner}
}

func (k *KeyboardPlayerInput) Fetch() (string, error) {
	hasMore := k.scanner.Scan()
	text := k.scanner.Text()

	if hasMore {
		return text, nil
	}

	return "", errors.New("end of input")
}

func main() {
	stub := &StubPlayerInput{inputs: []string{"Hello", "imagine these are", "lines typed", "23", "00", "", "66"}}

	// Read from the stub player input - for tests
	readAndPrintAll(stub)

	// Read from the real keyboard - for production
	keyboard := NewKeyboardPlayerInput()
	readAndPrintAll(keyboard)
}
