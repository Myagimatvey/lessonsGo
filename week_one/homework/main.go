package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	var path string
	for id, item := range args {
		if id == 1 {
			path = item
		} else if id == 2 && item == "-f" {
			fmt.Println("c файлами")
		}
	}


	openFile(path, 0)
}

func openFile(fileName string, widthFiles int) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	resultFuncOpen, errFuncOpen := file.ReadDir(0)
	if errFuncOpen != nil {
		return
	}
	// prevSpace := []int{}
	for id, item := range resultFuncOpen {
		if item.IsDir() {
			printName(item, widthFiles, id + 1 == len(resultFuncOpen))
			openFile(filepath.Join(fileName, item.Name()), widthFiles+1)
		} else {
			printName(item, widthFiles, id + 1 == len(resultFuncOpen))
		}
	}
}

func printName(name os.DirEntry, widthFile int, last bool) {
	for range widthFile {
		fmt.Printf(" ")
	}
	if last {
		fmt.Printf(" └───")
	} else {
		fmt.Printf("├───")
	}
	fmt.Printf("%s\n", name.Name())
}
