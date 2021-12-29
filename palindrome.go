package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]
	/*if len(os.Args) > {
		fmt.Println("non hai messo niente")

	}*/
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("errore acqusizione")
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parola := scanner.Text()
		parola = minuscolo(parola)
		fmt.Println(parola, palin(parola))

	}

}
func minuscolo(s string) string {
	return strings.ToLower(s)
}
func palin(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
