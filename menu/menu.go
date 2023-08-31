package menu

import (
	"bufio"
	"errors"
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

func (m *menu) add(itemname string) error {
	for _, item := range data {
		if item.name == itemname {
			return errors.New("item already exists!")
		}
	}
	*m = append(*m, menuItem{name: itemname, prices: make(map[string]float64)})
	return nil
}

func Print() {
	data.print()
}

func Add(itemname string) error {
	return data.add(itemname)
}
