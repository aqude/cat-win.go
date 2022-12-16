package main

import (
	"bufio"
	"fmt"
	"os"
)

func ctr() func() string {
	i := 0
	return func() string {
		i++

		return fmt.Sprintf("%d. ", i)
	}
}
func main() {
	ctr := ctr()
	command := os.Args[1]
	if command == "catw" {
		arg := os.Args[2]
		file, err := os.Open(arg)
		if err != nil {
			panic("write filename")
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		// optionally, resize scanner's capacity for lines over 64K, see next example
		for scanner.Scan() {
			fmt.Println(ctr(), scanner.Text())
		}
		// раз
	} else {
		fmt.Println("write - catw")
	}

}
