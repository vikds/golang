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

type IAnimal interface {
	Eat()
	Move()
	Speak()
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func Readline() string {
	fmt.Print("> ")
	if scanner := bufio.NewScanner(os.Stdin); scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ParseLine(line string) ([]string, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return nil, errors.New("Usage: ('newanimal'|name|type) or ('query'|name|action)")
	}
	if fields[0] != "newanimal" && fields[0] != "query" {
		return nil, fmt.Errorf("Incorrect command (newanimal/query): %s", fields[0])
	}
	return fields, nil
}

func main() {
	var (
		cow           = Animal{"grass", "walk", "moo"}
		bird          = Animal{"worms", "fly", "peep"}
		snake         = Animal{"mice", "slither", "hsss"}
		animals       = map[string]IAnimal{"cow": cow, "bird": bird, "snake": snake}
		methods       = map[string]func(_ IAnimal){"eat": IAnimal.Eat, "move": IAnimal.Move, "speak": IAnimal.Speak}
		named_animals = make(map[string]IAnimal)
	)

	for true {
		line := Readline()
		if line == "" {
			break
		}
		fields, err := ParseLine(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch fields[0] {
		case "newanimal":
			{
				animal, exists := animals[fields[2]]
				if !exists {
					fmt.Println("No such animal (cow/bird/snake):", fields[2])
					continue
				}
				named_animals[fields[1]] = animal
				fmt.Println("Created it!")
			}
		case "query":
			{
				animal, exists := named_animals[fields[1]]
				if !exists {
					fmt.Println("No such animal with the name:", fields[1])
					continue
				}
				action, exists := methods[fields[2]]
				if !exists {
					fmt.Println("No such animal action (eat/move/speak):", fields[2])
					continue
				}
				action(animal)
			}
		}
	}
}
