package main

import (
	"bufio"
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
}

var m1 = Monkey{
	Num:       1,
	Items:     []int{54, 65, 75, 74},
	Operation: func(old int) int { return old + 6 },
	Test:      func(val int) bool { return val%19 == 0 },
	TrueDest:  2,
	FalseDest: 0,
}

var m2 = Monkey{
	Num:       2,
	Items:     []int{79, 60, 97},
	Operation: func(old int) int { return old * old },
	Test:      func(val int) bool { return val%13 == 0 },
	TrueDest:  1,
	FalseDest: 3,
}

var m3 = Monkey{
	Num:       3,
	Items:     []int{74},
	Operation: func(old int) int { return old + 3 },
	Test:      func(val int) bool { return val%17 == 0 },
	TrueDest:  0,
	FalseDest: 1,
}

type Monkey struct {
	Num       int
	Items     []int
	Operation OperationFunc
	Test      TestFunc
	TrueDest  int
	FalseDest int
}

func New(num int, items []int, operation OperationFunc, test TestFunc, trueDest int, falseDest int) Monkey {
	m := Monkey{num, items, operation, test, trueDest, falseDest}
	return m
}

func (m Monkey) PrintNumItems() {
	fmt.Printf("Mondy %d has %d items\n", m.Num, len(m.Items))
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
}
