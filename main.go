package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	panicexample()
}

func loop() {
	i := 1
	for {
		fmt.Println(i)
		i += 1
		if i == 1000 {
			break
		}
	}

	for i < 1500 {
		fmt.Println(i)
		i += 1
	}

	for i := 1500; i < 1600; i += 1 {
		fmt.Println(i)
	}
}

func loopcollection() {
	arr := [3]int{101, 102, 103}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("Done!")

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Caramel Machiato", prices: map[string]float64{"small": 1.65, "medium": 1.95, "large": 2.15}},
		{name: "Penne Pollo", prices: map[string]float64{"half": 5.65, "full": 10.95, "double": 20.15}},
	}

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Println(size + "-" + strconv.FormatFloat(price, 'E', -1, 64))
		}
	}
}

func webapp() {

	fmt.Println("Hello Gophers!")
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:3000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}

func deferfunctions() {
	// defer follows first in last out
	db, _ := sql.Open("drivrName", "connectionStrint")
	defer db.Close()

	rows, _ := db.Query("some query!")
	defer rows.Close()
}

func panicexample() {
	fmt.Printf("%d divide by %d is equal to %d\n", 10, 2, divide(10, 2))
	fmt.Printf("%d divide by %d is equal to %d\n", 10, 0, divide(10, 0))
}

func divide(dividend int, divisor int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	x := dividend / divisor
	return x
}
