package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int

	for _, f := range files {
		if f.Size() == 0 {
			total = len(f.Name()) + 1
		}
	}
	fmt.Printf("Total Required Space: %d Bytes \n", total)

	names := make([]byte, 0, total)
	for _, f := range files {
		if f.Size() == 0 {
			n := f.Name()
			names = append(names, n...)
			names = append(names, '\n')
		}
	}
	err = ioutil.WriteFile("names.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", names)
}
