package main

import (
	"flag"
	"fmt"
	"gosort/utils"
	"strings"
)

var length, maximum int
var name string
var avaible = []string{"bubble", "bucket", "counting", "heap", "insertion", "merge", "quick", "radix", "selection", "shell"}

func init() {
	flag.IntVar(&length, "l", 20, "the length of the array to be sorted.")
	flag.IntVar(&maximum, "m", 100, "the maximum of the array's number.")
	flag.StringVar(&name, "n", "bubble", "the name of sort method used to sort the numbers. 10 names are avaible:\n"+strings.Join(avaible, ", "))
	flag.Parse()
}

func main() {
	fmt.Printf("Generating the random array...")
	arr := utils.RandArr(length, maximum)
	fmt.Printf("Success!\n")

	fmt.Printf("The random array is:\n")
	fmt.Printf("%v\n", arr)

	method := utils.SelMethod(name)
	method(arr)
	fmt.Printf("Success!\n")

	fmt.Printf("The sorted array is:\n")
	fmt.Printf("%v\n", arr)
}
