package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(fileName string) []string {
	var a []string

	filePath := "./seed/room/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a = append(a, scanner.Text())
	}

	return a
}
