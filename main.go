package main

import (
	"bufio"
	"bytes"
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
	generics()
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

type printer interface {
	Print() string
}

type user struct {
	username string
	id       int
}

func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

type menuItemV2 struct {
	name   string
	prices map[string]float64
}

func (mi menuItemV2) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprint(&b, "\t%10s%10.2f\n", size, cost)
	}
	return b.String()
}

func useinterfaces() {
	var p printer
	p = user{username: "Aashish", id: 42}
	fmt.Println(p.Print())

	p = menuItemV2{name: "Caramel Machiato",
		prices: map[string]float64{"small": 1.65, "medium": 1.95, "large": 2.15}}

	fmt.Println(p.Print())

	u, ok := p.(user)
	fmt.Println(u, ok)
	mi, ok := p.(menuItemV2)
	fmt.Println(mi, ok)

	switch v := p.(type) {
	case user:
		fmt.Println("Found a user", v)
	case menuItemV2:
		fmt.Println("Found a menuItemV2", v)
	default:
		fmt.Println("I am not sure!")
	}
}

func generics() {
	var flt = []float64{10.5, 6.9, 129.9}
	var clonedFlt = clone(flt)
	fmt.Println(&flt[0], &clonedFlt[0], clonedFlt)

	testScores := map[string]float64{
		"Harry":    99,
		"Hermoine": 100,
	}

	clonedTestScores := cloneMap(testScores)

	fmt.Println(clonedTestScores)

	int_arr := []int{1, 2, 3}
	flt_arr := []float64{1.2, 2.3, 3.4}
	str_arr := []string{"A", "B", "C"}

	int_arr_sum := add(int_arr)
	flt_arr_sum := add(flt_arr)
	str_arr_sum := add(str_arr)
	fmt.Printf("Sum of %v : %v \n", int_arr, int_arr_sum)
	fmt.Printf("Sum of %v : %v \n", flt_arr, flt_arr_sum)
	fmt.Printf("Sum of %v : %v \n", str_arr, str_arr_sum)
}

func clone[V any](s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

type addable interface {
	int | float64 | string
}

func add[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
