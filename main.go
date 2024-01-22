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

	var names []byte
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
