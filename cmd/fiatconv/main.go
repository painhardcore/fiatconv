package main

import (
	"fmt"
	"github.com/painhardcore/fiatconv/pkg/conv"
	"os"
	"strconv"
)

func main() {
	res, err := conv.Currency(checkNgetArgs())
	if err != nil {
		fmt.Printf("Error occured: %s", err)
		os.Exit(1)
	}
	fmt.Println(res)
}

func checkNgetArgs() (float64, string, string) {
	a := os.Args[1:]
	if len(a) != 3 {
		fmt.Println("Wrong number of arguments")
		os.Exit(1)
	}
	f, err := strconv.ParseFloat(a[0], 64)
	if err != nil {
		fmt.Println("First argument should be float number")
		os.Exit(1)
	}
	return f, a[1], a[2]
}
