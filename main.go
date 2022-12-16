package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
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
	lines_ctr := 0
	command := os.Args[1]
	if command == "catw" {
		arg := os.Args[2]
		file, err := os.Open(arg)
		if err != nil {
			panic("write filename")
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		// Вызываем метод Scan() до тех пор, пока он не вернет false

		for scanner.Scan() {
			// разделитель стрку на слова
			str := strings.Fields(scanner.Text())
			for _, word := range str {
				if word == "def" || word == "if" || word == "return" || word == "else" || word == "elif" || word == "for" || word == "while" || word == "try" || word == "except" || word == "finally" || word == "with" || word == "class" {
					fmt.Println(ctr(), color.InRed(word))

				} else {
					fmt.Println(ctr(), word)
				}
			}

			// fmt.Println(ctr(), scanner.Text())
			lines_ctr++
		}
		fmt.Println("lines:", lines_ctr)

	} else {
		fmt.Println("write - catw")
	}

}
