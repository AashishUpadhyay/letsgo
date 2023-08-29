package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println("Item: " + item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Println(size + "-" + strconv.FormatFloat(price, 'E', -1, 64))
		}
	}
}

func (m *menu) add() {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	*m = append(*m, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}

func Print() {
	data.print()
}

func Add() {
	data.add()
}
