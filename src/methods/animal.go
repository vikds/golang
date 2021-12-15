package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func Readline() string {
	fmt.Print("animal func> ")
	if scanner := bufio.NewScanner(os.Stdin); scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ParseLine(line string, animals map[string]Animal, methods map[string]func(_ *Animal)) (*Animal, func(_ *Animal), error) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		return nil, nil, errors.New("Usage: cow|bird|snake eat|move|speak")
	}
	animal, exists := animals[fields[0]]
	if !exists {
		return nil, nil, fmt.Errorf("No such animal: %s", fields[0])
	}
	method, exists := methods[fields[1]]
	if !exists {
		return nil, nil, fmt.Errorf("No such method: %s", fields[1])
	}

	return &animal, method, nil
}

func main() {
	var (
		cow     = Animal{"grass", "walk", "moo"}
		bird    = Animal{"worms", "fly", "peep"}
		snake   = Animal{"mice", "slither", "hsss"}
		animals = map[string]Animal{"cow": cow, "bird": bird, "snake": snake}
		methods = map[string]func(_ *Animal){"eat": (*Animal).Eat, "move": (*Animal).Move, "speak": (*Animal).Speak}
	)
	for true {
		line := Readline()
		if line == "" {
			break
		}
		animal, method, err := ParseLine(line, animals, methods)
		if err != nil {
			fmt.Println(err)
			continue
		}
		method(animal)
	}
}
