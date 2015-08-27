package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE:", os.Args[0], "filename")
		os.Exit(1)
	}

	filter := NewBloomFilter(8)

	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		filter.AddElement(scanner.Text())
	}

	fmt.Println(filter.GetEmpty())

	fmt.Println(filter.ContainsElement("holy handgrenade"))
	fmt.Println(filter.ContainsElement("horsemeat"))
	fmt.Println(filter.ContainsElement("sandwich"))

	fmt.Println(filter.ContainsElement("JENNIFER"))
	fmt.Println(filter.CountElements())

	filter.RemoveElement("JENNIFER")
	fmt.Println(filter.ContainsElement("JENNIFER"))
	fmt.Println(filter.CountElements())
	fmt.Println(filter.FalsePositiveRate())
}
