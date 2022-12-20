package main

import (
	"bufio"
	"fmt"
	"os"
)

type OperationFunc func(old int) int
type TestFunc func(val int) bool

type Monkey struct {  
    Num int
    Items []int
    Operation OperationFunc
    Test TestFunc
    TrueDest int
    FalseDest int
}

func New(num int, items []int, operation OperationFunc, test TestFunc, trueDest int, falseDest int) Monkey {  
    m := Monkey {num, items, operation, test, trueDest, falseDest }
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
