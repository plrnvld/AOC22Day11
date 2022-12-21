package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type OperationFunc func(old int) int
type TestFunc func(val int) bool

var m0 = Monkey{
	Num:       0,
	Items:     []int{79, 98},
	Operation: func(old int) int { return old * 19 },
	Test:      func(val int) bool { return val%23 == 0 },
	TrueDest:  2,
	FalseDest: 3,
    Inspect: 0,
}

var m1 = Monkey{
	Num:       1,
	Items:     []int{54, 65, 75, 74},
	Operation: func(old int) int { return old + 6 },
	Test:      func(val int) bool { return val%19 == 0 },
	TrueDest:  2,
	FalseDest: 0,
    Inspect: 0,
}

var m2 = Monkey{
	Num:       2,
	Items:     []int{79, 60, 97},
	Operation: func(old int) int { return old * old },
	Test:      func(val int) bool { return val%13 == 0 },
	TrueDest:  1,
	FalseDest: 3,
    Inspect: 0,
}

var m3 = Monkey{
	Num:       3,
	Items:     []int{74},
	Operation: func(old int) int { return old + 3 },
	Test:      func(val int) bool { return val%17 == 0 },
	TrueDest:  0,
	FalseDest: 1,
    Inspect: 0,
}

var monkeys = []*Monkey{&m0, &m1, &m2, &m3}

type Monkey struct {
	Num       int
	Items     []int
	Operation OperationFunc
	Test      TestFunc
	TrueDest  int
	FalseDest int
    Inspect int
}

func Send(val int, to Monkey) {
	to.Items = append(to.Items, val)
}

func WorryLevel(from *Monkey, val int) int {
	return from.Operation(val) / 3
}

func TestIt(from *Monkey, worryLevel int) bool {
	return from.Test(worryLevel)
}

func Destination(from *Monkey, worryLevel int) *Monkey {
	isTrue := from.Test(worryLevel)

	var monkeyNum int
	if isTrue {
		monkeyNum = from.TrueDest
	} else {
		monkeyNum = from.FalseDest
	}

	monkey, _ := NumToMonkey(monkeyNum)
	return monkey
}

func NumToMonkey(num int) (*Monkey, error) {
	if num == 0 {
		return &m0, nil
	} else if num == 1 {
		return &m1, nil
	} else if num == 2 {
		return &m2, nil
	} else if num == 3 {
		return &m3, nil
	} else {
		return &Monkey{}, errors.New(fmt.Sprintf("No monkey found with number %d", num))
	}
}

func New(num int, items []int, operation OperationFunc, test TestFunc, trueDest int, falseDest int, inspect int) Monkey {
	m := Monkey{num, items, operation, test, trueDest, falseDest, inspect}
	return m
}

func (m Monkey) PrintNumItems() {
	fmt.Printf("Monkey %d has %d items\n", m.Num, len(m.Items))
}

func (m *Monkey) AddItem(item int) {
	m.Items = append(m.Items, item)
}

func runOneRound() {
	for _, m := range monkeys {
		for _, item := range m.Items {
			level := WorryLevel(m, item)
			dest := Destination(m, level)
			fmt.Printf("> Sending item %d with level %d from m%d to m%d\n", item, level, m.Num, dest.Num)
			dest.AddItem(level)
            m.Inspect = m.Inspect + 1

			fmt.Printf(">>> Monkey %d now has %v\n\n", dest.Num, dest.Items)
		}

		m.Items = m.Items[:0]
	}
}

func printItems() {
	for _, m := range monkeys {
		fmt.Printf("Monkey %d: %v\n", m.Num, m.Items)
	}

	fmt.Printf("*****\n\n")
}

func printInspect() {
	for _, m := range monkeys {
		fmt.Printf("Monkey %d inspected %d times\n", m.Num, m.Inspect)
	}
}

func main() {
	readFile, err := os.Open("Example.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()

	printItems()

	for i := 0; i < 20; i++ {
        fmt.Printf("\n=== Round %d ===\n\n", i+1)
		runOneRound()

		printItems()
	}

    printInspect()
}
