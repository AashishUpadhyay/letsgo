package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"demo/menu"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	loopcollection()
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
loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) print menu")
		fmt.Println("2) add item")
		fmt.Println("3) quit")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.Add()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}

}

func loopArr() {
	arr := [3]int{101, 102, 103}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("Done!")
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
