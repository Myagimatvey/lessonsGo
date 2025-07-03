package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	alreadySeen := make(map[string]bool)
	for in.Scan() {
		txt := in.Text()

		if alreadySeeb

		fmt.Println(txt)
	}
}