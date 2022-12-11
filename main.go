package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		ls()
	} else {
		lswithdir(os.Args[1])
	}
}

func ls() {
	lswithdir(".")
}

func lswithdir(dirname string) {
	dir_entries, err := os.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	output(dir_entries)
}

func output(entries []os.DirEntry) {
	for _, entry := range entries {
		if string(entry.Name())[0] == '.' {
			continue
		}
		fmt.Printf("%s ", entry.Name())
	}
	fmt.Println()
}
